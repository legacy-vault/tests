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

// worker.go.

// Main Worker.

// Date: 2018-10-19.

package main

import (
	"log"
	"sync"
	"time"
)

// Messages.
const MsgRawDataChannelClosed = "Raw Data Channel Read Failure. " +
	"Channel is closed."
const MsgWorkerStop = "Worker has been stopped."

// Message Formats.
const MsgFormatWorkerStatistics = "Worker has received %v Lines.\r\n"

const SleepTickMsWorker = 250

// Main Worker.
// The Worker always reads and processes all Tasks from the Queue,
// even if the input Stream Reader has been stopped.
// Such Behaviour is implemented to ensure that all Tasks received
// are going to be finished, this is a safe Method to process Queues.
func worker(
	rawDataChannel chan string,
	resultsChannel chan ResultEntry,
	busyWorkersCount *sync.WaitGroup,
	workersQuitChannel chan bool,
	verboseMode bool,
) {

	var err error
	var loop bool
	var rawText string
	var received bool
	var receivedLinesCount uint64
	var sleepTick time.Duration
	var stopSignal bool
	var validURL string

	defer busyWorkersCount.Done()

	// Preparations.
	sleepTick = time.Millisecond * SleepTickMsWorker

	// Receive raw Data.
	loop = true
	for loop {

		select {

		case rawText, received = <-rawDataChannel:
			if received == false {

				// Channel is not available.

				// Log.
				log.Println(MsgRawDataChannelClosed)
				loop = false

			} else {

				// New incoming raw Data.
				receivedLinesCount++

				// Check it.
				validURL, err = syntaxCheck(rawText)
				if err == nil {
					// Valid URL must be processed.
					urlProcess(
						validURL,
						resultsChannel,
						verboseMode,
					)
				}
				// URLs which are not valid are simply ignored.
			}

		default:
			// No Tasks.
			// Such Situation can happen in two Cases:
			//	1.	We are using manual (Keyboard) Input and the User has not
			// 		yet put anything into the Stream.
			//	2.	The Stream Reader has been stopped and the Queue is empty.

			// Try to get a Stop Signal.
			select {

			case stopSignal = <-workersQuitChannel:

				if stopSignal == WorkersQuitChanSignalQuit {
					// The Queue is empty and we must quit.
					loop = false
				}

			default:

				// No Tasks are available & no Stop Signal.
				// Sleep for a while...
				time.Sleep(sleepTick)
			}
		}
	}

	// Verbose Log.
	if verboseMode {
		log.Printf(MsgFormatWorkerStatistics, receivedLinesCount)
		log.Println(MsgWorkerStop)
	}
}
