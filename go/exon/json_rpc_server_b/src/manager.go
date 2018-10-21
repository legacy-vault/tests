// manager.go

// Version: 0.1.
// Date: 2017-07-06.
// Author: McArcher.

package main

import (
	"errors"
	"log"
	"time"
)

//------------------------------------------------------------------------------
// Variables
//------------------------------------------------------------------------------

var manager_sleep_interval = 10 // Seconds
// N.B.
// Large Interval makes Lag longer but rarer.
// Small Interval makes Lag shorter but more often.

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

func manager() {

	// Periodically saves the List to the DataBase.

	var sleepInterval time.Duration

	var ok bool
	var err error
	var error_msg string
	var ret_err error

	sleepInterval = time.Second * time.Duration(manager_sleep_interval)

	for {

		time.Sleep(sleepInterval)

		ok, err = Users.StoreDelta()
		if !ok {

			error_msg = "Manager can not store the List. " + err.Error()
			ret_err = errors.New(error_msg)
			log.Println(error_msg) //

			// No return in Go-Routine => panic
			panic(ret_err.Error())
		}
	}
}

//------------------------------------------------------------------------------
