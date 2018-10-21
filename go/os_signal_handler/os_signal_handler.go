// os_signal_handler.go

// Version: 0.1.
// Date: 2017-07-07.
// Author: McArcher.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

//------------------------------------------------------------------------------
// Variables
//------------------------------------------------------------------------------

var OSSignalChan chan os.Signal

//------------------------------------------------------------------------------

const OSSignalChanSize = 1

//------------------------------------------------------------------------------
// Signals
//------------------------------------------------------------------------------
/*

	Useful Sources of Information:

	https://godoc.org/os/signal
	https://golang.org/pkg/os/signal/
	https://en.wikipedia.org/wiki/Unix_signal


	Types of Signals (popular):

	SIGINT		Interrupt from Terminal (Control+C, ^C).
	SIGKILL		Is not captured normally. Is used by O.S. to stop Process Execution (destroy Process).
	SIGQUIT		Quit-Signal from Terminal (Control-Backslash, ^\).
	SIGSTOP		Is not captured normally. Is used by O.S. to pause Process Execution.
	SIGTERM		Termination of Process.
	SIGTSTP		Terminal Stop (Control+Z, ^Z).

*/
//------------------------------------------------------------------------------

func main() {

	fmt.Println("Hello.") //

	// Start Signal Handler
	go OSSignalHandler()

	// Eternal Loop of Nothing
	EternalLoop()

}

//------------------------------------------------------------------------------

func EternalLoop() {

	// This Function loops forever and does not load CPU at the same Time.

	var EternalLoopChan chan bool

	EternalLoopChan = make(chan bool)

	<-EternalLoopChan
}

//------------------------------------------------------------------------------

func OSSignalHandler() {

	// Handles Signals from Operating System (O.S.).

	var sig os.Signal

	OSSignalChan = make(chan os.Signal, OSSignalChanSize)

	// Redirect Signals to the Channel
	signal.Notify(OSSignalChan, os.Interrupt)    // SIGINT is ANSI
	signal.Notify(OSSignalChan, syscall.SIGKILL) // SIGKILL is POSIX, cannot be captured
	signal.Notify(OSSignalChan, syscall.SIGQUIT) // SIGQUIT is POSIX
	signal.Notify(OSSignalChan, syscall.SIGSTOP) // SIGSTOP is POSIX, cannot be captured
	signal.Notify(OSSignalChan, syscall.SIGTERM) // SIGTERM is ANSI
	signal.Notify(OSSignalChan, syscall.SIGTSTP) // SIGTSTP is POSIX

	// Start Monitoring Process
	for {

		sig = <-OSSignalChan

		if sig == os.Interrupt {

			fmt.Println("Received SIGINT.") //
			os.Exit(0)

		}

		if sig == syscall.SIGTERM {

			fmt.Println("Received SIGTERM.") //
			os.Exit(0)

		}

		if sig == syscall.SIGQUIT {

			fmt.Println("Received SIGQUIT.") //
			os.Exit(0)

		}

		if sig == syscall.SIGKILL {

			fmt.Println("Received SIGKILL.") //
			os.Exit(0)

		}

		if sig == syscall.SIGSTOP {

			fmt.Println("Received SIGSTOP.") //
			os.Exit(0)

		}

		if sig == syscall.SIGTSTP {

			fmt.Println("Received SIGTSTP.") //
			os.Exit(0)

		}

		fmt.Println("UnKnown Signal Received!") //
		os.Exit(0)
	}
}

//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
