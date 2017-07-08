// manager.go

// Version: 0.4.
// Date: 2017-07-08.
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
// N.B.
// Large Interval makes Lag longer but rarer.
// Small Interval makes Lag shorter but more often.
// Do not set it less than 1 Second.

var manager_ping_interval = 1 // Seconds

var SaveIsInProgress bool        // Saving Process is now running
var EmergencySaveHasStarted bool // Is used to guarantee that Emergency
// Save has started after the 'EmergencySaveIsNeeded' Flag was risen.
var EmergencySaveHasFinished bool
var EmergencySaveMsgA string // Message written to Log in Case of Emergency Save
var EmergencySaveMsgB string

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

	EmergencySaveMsgA = "Emergency Save started..."
	EmergencySaveMsgB = "Emergency Save completed."
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
	var SaveIsNeeded, RegularSaveIsNeeded bool

	var ret_err error
	var error_msg string
	var ok bool
	var err error

	now = time.Now()
	timeSinceLastSave = now.Sub(*lastSaveTime)

	RegularSaveIsNeeded = (timeSinceLastSave > saveInterval)
	SaveIsNeeded = (RegularSaveIsNeeded || EmergencySaveIsNeeded) && (!EmergencySaveHasFinished)

	if SaveIsNeeded {

		// Emergency?
		if EmergencySaveIsNeeded {

			EmergencySaveHasStarted = true
			EmergencySaveHasFinished = false
			log.Println(EmergencySaveMsgA) // 'Emergency Save started'
		}

		// Save to File
		SaveIsInProgress = true
		ok, err = Users.StoreDelta()
		SaveIsInProgress = false
		if !ok {

			error_msg = "Manager can not store the List. " + err.Error()
			ret_err = errors.New(error_msg)
			log.Println(error_msg) //

			// No return in Go-Routine => panic
			panic(ret_err.Error())
		}

		// Update lastSaveTime
		*lastSaveTime = time.Now()

		// Emergency?
		if EmergencySaveHasStarted {

			log.Println(EmergencySaveMsgB) // 'Emergency Save completed'
			EmergencySaveHasFinished = true
			// Now we are sure that we have saved (or, at least, have tried
			// to save) new Data to File.
		}
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
