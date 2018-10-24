//============================================================================//
//
// Copyright © 2018 by McArcher.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
//============================================================================//
//
// Web Site:		'https://github.com/legacy-vault'.
// Author:			McArcher.
// Creation Date:	2018-10-19.
// Web Site Address is an Address in the global Computer Internet Network.
//
//============================================================================//

// main.go.

// Standard Input Stream Reader and Processor :: Main File.

// Date: 2018-10-19.

// The Stream Reader tries to read the entire input Stream.
// If an Event preventing the reading Process occurs,
// the Reader stops receiving Data and waits for all the workers to finish
// their work. While the Reader is halt and workers are finishing existing
// Tasks, the Application does not respond to normal Termination Signals.
// Normal Termination Signals (SIGTERM, SIGINT) are used to stop Data receiving
// from 'stdin'. Such Behaviour is implemented to ensure that all Tasks
// received are going to be finished, this is a safe Method to process Queues.
// To stop the Application entirely, one may use something like SIGSTOP or
// SIGKILL, but it is not recommended to do so.

package main

import (
	"io"
	"log"
	"os"
	"sync"
)

// Exit Codes.
const ExitCodeNormal = 0
const ExitCodeError = 255

// Messages.
const MsgSignalBad = "Bad Signal is received."
const MsgEOS = "End of Stream is reached."

// Error Message Prefixes.
const ErrPrefixInitFailure = "Initialization Failure. "
const ErrPrefixReadFailure = "Read Failure. "

// Program's Entry Point.
func main() {

	var app Application
	var busyCollector sync.WaitGroup
	var busyWorkersCount sync.WaitGroup
	var err error
	var i int
	var msg string

	// Initialize the Application.
	err = initialize(&app)
	if err != nil {
		msg = ErrPrefixInitFailure + err.Error()
		log.Println(msg)
		os.Exit(ExitCodeError)
	}

	// Start Results Collector.
	busyCollector.Add(1)
	go resultsCollector(
		app.ResultsChannel,
		&app.Results,
		&busyCollector,
		app.ResultsCollectorQuitChan,
		app.VerboseMode,
	)

	// Start Workers.
	for i = 0; i < WorkersCount; i++ {
		busyWorkersCount.Add(1)
		go worker(
			app.RawDataChannel,
			app.ResultsChannel,
			&busyWorkersCount,
			app.WorkersQuitChannel,
			app.VerboseMode,
		)
	}

	// Receive Data (in a Loop) from Standard Input Stream.
	err = stdinReader(&app)
	if err != nil {

		// Check Error Type.
		if err == io.EOF {
			// Verbose Log.
			if app.VerboseMode {
				log.Println(MsgEOS)
			}
		} else {
			// Log.
			msg = ErrPrefixReadFailure + err.Error()
			log.Println(msg)
		}
	}

	// Wait for all active Workers to finish their Job.
	busyWorkersCount.Wait()

	// Send Stop Signal to the Results Collector.
	app.ResultsCollectorQuitChan <- ResultsCollectorQuitChanSignalQuit

	// Wait for Results Collector to finish its Job.
	busyCollector.Wait()

	// Show Results.
	resultsShow(app.Results)

	// Normal Exit.
	os.Exit(ExitCodeNormal)
}
