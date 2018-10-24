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

// oss.go.

// Operation System Signals Handler.

// Date: 2018-10-19.

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Messages.
const MsgFormatSignalReceived = "'%v' Signal is received."

// Signal Names.
const SignalSIGTERM = "SIGTERM"
const SignalSIGINT = "SIGINT"

// Initializes OS Signals Handling.
func ossInit(app *Application) error {

	// Prepare Channels.
	app.OsSignalsTermChan = make(chan os.Signal, OsSignalsTermChanSize)
	app.AppQuitChannel = make(chan bool, AppQuitChanSize)

	// Bind O.S. Signals to Channel.
	signal.Notify(
		app.OsSignalsTermChan,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	// Start Signals Monitoring.
	go handlerOfTermination(app.OsSignalsTermChan, app.AppQuitChannel)

	return nil
}

// Handler of terminating Signals: 'SIGTERM' and 'SIGINT'.
// Due to the Nature of the Stream Reading Algorithm,
// there may be a Situation when the Quit Signal has been sent to
// the Reader, but the Reader is blocked (halt) waiting for Data
// from the input Stream. To prevent such Halts, the User may send
// an 'End of Stream' Signal to the Terminal (Ctrl+D).
func handlerOfTermination(
	signalsChannel chan os.Signal,
	quitChannel chan bool,
) {

	var msgSigInt string
	var msgSigTerm string
	var osSignal os.Signal

	// Prepare Messages.
	msgSigInt = fmt.Sprintf(MsgFormatSignalReceived, SignalSIGINT)
	msgSigTerm = fmt.Sprintf(MsgFormatSignalReceived, SignalSIGTERM)

	// Listen for O.S. Signals.
	for osSignal = range signalsChannel {

		// Got a Signal.
		switch osSignal {

		case syscall.SIGTERM:

			// Start Application Termination.
			quitChannel <- AppQuitChanSignalQuit

			// Log.
			log.Println(msgSigTerm)

		case syscall.SIGINT:

			// Start Application Termination.
			quitChannel <- AppQuitChanSignalQuit

			// Log.
			log.Println(msgSigInt)
		}
	}
}
