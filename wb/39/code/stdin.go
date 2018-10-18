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

// stdin.go.

// Standard Input Stream Reading.

// Date: 2018-10-19.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Messages.
const MsgReadingStopped = "Reader has stopped receiving Data from 'stdin'."
const MsgConsoleInputIvitationSymbol = ">"
const MsgReaderStart = "Reader has started."
const MsgReaderStop = "Reader has stopped."

// Message Formats.
const MsgFormatReaderStatistics = "Reader has received %v Lines.\r\n"

const ASCIINewLine = '\n'

// Standard Input Stream Reader.
// The Reader reads Data from 'stdin' Stream.
// The Reader stops reading and quits if any of the following Events happens:
//	*	The Reader finds a Stream's End;
//	*	Read Failure occurs;
//	*	An O.S. sends a Signal for Termination.
func stdinReader(app *Application) error {

	var delimiter byte
	var err error
	var i int
	var receivedLinesCount uint64
	var quitSignal bool
	var stdinReader *bufio.Reader
	var textLine string

	// Verbose Log.
	if app.VerboseMode {
		fmt.Println(MsgReaderStart)
	}

	// Prepare Reader.
	stdinReader = bufio.NewReader(os.Stdin)
	delimiter = ASCIINewLine

	// Read Loop.
	for app.ReadingFromStdinIsEnabled {

		// Keyboard Input Mode Invitation.
		if app.KeyboardInputMode {
			fmt.Print(MsgConsoleInputIvitationSymbol)
		}

		// Get next text Line.
		textLine, err = stdinReader.ReadString(delimiter)
		if err != nil {

			// An Error occurred.

			// Forbid Standard Input Stream Reading.
			app.ReadingFromStdinIsEnabled = false

		} else {

			// Received a text Line.
			// Send it to a next Stage.
			receivedLinesCount++
			app.RawDataChannel <- textLine
		}

		// Check Quit Channel for Signals.
		select {

		case quitSignal = <-app.AppQuitChannel:

			// Check Signal.
			if quitSignal == AppQuitChanSignalQuit {

				// Forbid Standard Input Stream Reading.
				app.ReadingFromStdinIsEnabled = false

				// Verbose Log.
				if app.VerboseMode {
					log.Println(MsgReadingStopped)
				}

			} else {

				// Verbose Log.
				if app.VerboseMode {
					log.Println(MsgSignalBad)
				}
			}

		default:
			// Do Nothing.
		}
	}

	// Send Stop Signals to all Workers.
	// Each Workers requires a single Signal to quit.
	for i = 0; i < WorkersCount; i++ {
		app.WorkersQuitChannel <- WorkersQuitChanSignalQuit
	}

	// Verbose Log.
	if app.VerboseMode {
		log.Printf(MsgFormatReaderStatistics, receivedLinesCount)
		log.Println(MsgReaderStop)
	}

	return err
}
