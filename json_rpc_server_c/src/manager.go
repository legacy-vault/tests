// manager.go

// Version: 0.3.
// Date: 2017-07-07.
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

var manager_save_interval = 10 // Seconds
var manager_ping_interval = 1  // Seconds
// N.B.
// Large Interval makes Lag longer but rarer.
// Small Interval makes Lag shorter but more often.

// Channels
var ManagerAliveChan chan bool // Prevents Store Freeze

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

func UserManager() {

	// Manages User Tasks.

	var task UserTask
	var user UserModel
	var gotUser UserModel
	var lastSaveTime time.Time
	var saveInterval time.Duration
	var ok bool
	var err error

	saveInterval = time.Second * time.Duration(manager_save_interval)
	lastSaveTime = time.Now()

	for {

		select {

		case task = <-AddedUsersChan:

			// Add User
			user.login = task.user.login
			user.regDate = task.user.regDate
			user.uuid = task.user.uuid
			ok, err = Users.Add(user, false)

			// Send Results
			task.result = ok
			task.err = err
			task.sender <- task

			// Store if needed
			UserManagerStoreData(&lastSaveTime, saveInterval)

		case task = <-ModifiedUsersChan:

			// Modify User
			user.login = task.user.login
			user.regDate = task.user.regDate
			user.uuid = task.user.uuid
			ok, err = Users.Modify(user)

			// Send Results
			task.result = ok
			task.err = err
			task.sender <- task

			// Store if needed
			UserManagerStoreData(&lastSaveTime, saveInterval)

		case task = <-GotUsersChan:

			// Get User
			user.login = task.user.login
			user.regDate = task.user.regDate
			user.uuid = task.user.uuid
			ok, err = Users.Get(user, &gotUser)

			// Send Results
			task.result = ok
			task.err = err
			task.user = gotUser
			task.sender <- task

			// Store if needed
			UserManagerStoreData(&lastSaveTime, saveInterval)

		case <-ManagerAliveChan:

			// Store if needed
			UserManagerStoreData(&lastSaveTime, saveInterval)

		}

	}
}

//------------------------------------------------------------------------------

func UserManagerStoreData(lastSaveTime *time.Time, saveInterval time.Duration) {

	// Stores the Data into File if needed.

	var now time.Time
	var timeSinceLastSave time.Duration
	var ret_err error
	var error_msg string
	var ok bool
	var err error

	now = time.Now()
	timeSinceLastSave = now.Sub(*lastSaveTime)
	if timeSinceLastSave > saveInterval {

		// Save to File
		ok, err = Users.StoreDelta()
		if !ok {

			error_msg = "Manager can not store the List. " + err.Error()
			ret_err = errors.New(error_msg)
			log.Println(error_msg) //

			// No return in Go-Routine => panic
			panic(ret_err.Error())
		}

		// Update lastSaveTime
		*lastSaveTime = time.Now()
	}

}

//------------------------------------------------------------------------------
func ManagerPinger() {

	// Keeps UserManager alive.

	// N.B.
	// When no Requests are sent to UserManager, it halts waiting for next
	// Task from Channels. All the Records which needed to be saved to File
	// also halt in the Queue waiting for the Manager to receive any Task.
	// This Pinger periodically sends 'I am alive' empty Tasks to prevent
	// halting. Usage of Golang's 'default' case in 'select' cannot help, as
	// it loads the CPU Core to 100%.

	var pingInterval time.Duration

	pingInterval = time.Second * time.Duration(manager_ping_interval)

	for {

		ManagerAliveChan <- true

		time.Sleep(pingInterval)
	}

}

//------------------------------------------------------------------------------
