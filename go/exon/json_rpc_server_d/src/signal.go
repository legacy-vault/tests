// signal.go

// Version: 0.4.
// Date: 2017-07-08.
// Author: McArcher.

//------------------------------------------------------------------------------
/*

	/!\  WARNING  /!\

	These Signals work in 'UNIX' and 'GNU/Linux' Operating Systems.

*/
//------------------------------------------------------------------------------

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

//------------------------------------------------------------------------------
// Variables
//------------------------------------------------------------------------------

var OSSignalChan chan os.Signal

//------------------------------------------------------------------------------

const OSSignalChanSize = 16

//------------------------------------------------------------------------------

func OSSignalHandler() {

	// Handles Signals from Operating System (O.S.).

	var sig os.Signal
	var SignalIsGood bool

	OSSignalChan = make(chan os.Signal, OSSignalChanSize)

	// Redirect Signals to the Channel
	signal.Notify(OSSignalChan, os.Interrupt)    // SIGINT is ANSI
	signal.Notify(OSSignalChan, syscall.SIGQUIT) // SIGQUIT is POSIX
	signal.Notify(OSSignalChan, syscall.SIGTERM) // SIGTERM is ANSI
	signal.Notify(OSSignalChan, syscall.SIGTSTP) // SIGTSTP is POSIX

	// Start Monitoring Process
	for {

		sig = <-OSSignalChan

		SignalIsGood = (sig == os.Interrupt) || (sig == syscall.SIGQUIT) ||
			(sig == syscall.SIGTERM) || (sig == syscall.SIGTSTP)

		if SignalIsGood {

			log.Println("Received Shutdown Signal. Starting Shutdown...") //
			EmergencyShutdownIsNeeded = true

		} else {

			panic("Strange Signal Received!") //
		}

	}
}

//------------------------------------------------------------------------------
