// function.go

// Version: 0.4.
// Date: 2017-07-08.
// Author: McArcher.

package main

//------------------------------------------------------------------------------

func EternalLoop() {

	// This Function loops forever and does not load CPU at the same Time.
	// It is used when we want to wait for Something at the End of our Program.
	// For Example, we must wait for a Go-Routine to do its Job.

	var EternalLoopChan chan bool

	EternalLoopChan = make(chan bool)

	<-EternalLoopChan
}

//------------------------------------------------------------------------------
