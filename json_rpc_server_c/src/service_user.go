// service_user.go

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
// Types
//------------------------------------------------------------------------------

type User int // Used as RPC Service called 'User'

type Arguments struct {
	UUID    uint64
	LOGIN   string
	REGDATE int64 // UNIX Timestamp
}

type GetResult struct {
	UUID    uint64
	LOGIN   string
	REGDATE int64 // UNIX Timestamp
	OK      bool
}

//------------------------------------------------------------------------------
// Variables
//------------------------------------------------------------------------------

// Service
var service_user *User

//------------------------------------------------------------------------------
// Methods
//------------------------------------------------------------------------------

func (u *User) Add(args *Arguments, result *bool) error {

	// RPC Action: 'User.Add'.

	var uuid uint64
	var regDate int64
	var login string
	var user UserModel

	var uuidIsEmpty bool
	var loginIsEmpty bool
	var loginIsTooLong bool
	var regDateIsEmpty bool

	var task, feedback UserTask
	var rcvChan chan UserTask

	var ret_err error
	var error_msg string

	// Read Arguments
	uuid = args.UUID
	regDate = args.REGDATE
	login = args.LOGIN

	// Preparations
	// ...

	// 1. Check for empty UUID.
	uuidIsEmpty = (uuid == 0)
	if uuidIsEmpty {

		error_msg = "UUID is empty."
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		*result = false
		return ret_err
	}

	// 2. Check for empty login.
	loginIsEmpty = (login == "")
	if loginIsEmpty {

		error_msg = "Login is empty."
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		*result = false
		return ret_err
	}

	// 3. Check Length of login
	loginIsTooLong = !isLoginLengthGood(&login)
	if loginIsTooLong {

		error_msg = "Login is too long."
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		*result = false
		return ret_err
	}

	// 4. Check empty regDate
	regDateIsEmpty = (regDate == 0)
	if regDateIsEmpty {

		regDate = time.Now().Unix()
	}

	// Create Object of UserModel
	user.uuid = uuid
	user.regDate = regDate
	user.login = login

	// Fill Task
	rcvChan = make(chan UserTask)
	task.sender = rcvChan
	task.user = user
	task.result = false
	task.err = nil

	// Send Task
	AddedUsersChan <- task

	// Get Feedback
	feedback = <-rcvChan

	// Return
	*result = feedback.result
	return feedback.err
}

//------------------------------------------------------------------------------

func (u *User) Get(args *Arguments, result *GetResult) error {

	// 'User.Get' RPC Action.

	var uuid uint64
	var regDate int64
	var login string
	var user UserModel

	var uuidIsEmpty bool
	var loginIsEmpty bool
	var loginIsTooLong bool

	var task, feedback UserTask
	var rcvChan chan UserTask

	var ret_err error
	var error_msg string

	// Read Arguments
	uuid = args.UUID
	regDate = args.REGDATE
	login = args.LOGIN

	// Preparations
	uuidIsEmpty = (uuid == 0)
	loginIsEmpty = (login == "")

	// 1. Both UUID & login are empty
	if uuidIsEmpty && loginIsEmpty {

		error_msg = "Both UUID & login are empty."
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		result.OK = false
		return ret_err
	}

	// 2. Check Length of login
	loginIsTooLong = !isLoginLengthGood(&login)
	if loginIsTooLong {

		error_msg = "Login is too long."
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		result.OK = false
		return ret_err
	}

	// Create Object of UserModel
	user.uuid = uuid
	user.regDate = regDate
	user.login = login

	// Fill Task
	rcvChan = make(chan UserTask)
	task.sender = rcvChan
	task.user = user
	task.result = false
	task.err = nil

	// Send Task
	GotUsersChan <- task

	// Get Feedback
	feedback = <-rcvChan

	// Return
	result.LOGIN = feedback.user.login
	result.REGDATE = feedback.user.regDate
	result.UUID = feedback.user.uuid
	return feedback.err
}

//------------------------------------------------------------------------------

func (u *User) Modify(args *Arguments, result *bool) error {

	// 'User.Modify' RPC Action.

	var uuid uint64
	var regDate int64
	var login string
	var user UserModel

	var uuidIsEmpty bool
	var loginIsEmpty bool
	var loginIsTooLong bool
	var regDateIsEmpty bool

	var task, feedback UserTask
	var rcvChan chan UserTask

	var ret_err error
	var error_msg string

	// Read Arguments
	uuid = args.UUID
	regDate = args.REGDATE
	login = args.LOGIN

	// Preparations
	// ...

	// 1. Check for empty UUID.
	uuidIsEmpty = (uuid == 0)
	if uuidIsEmpty {

		error_msg = "UUID is empty."
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		*result = false
		return ret_err
	}

	// 2. Check for empty login.
	loginIsEmpty = (login == "")
	if loginIsEmpty {

		error_msg = "Login is empty."
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		*result = false
		return ret_err
	}

	// 3. Check Length of login
	loginIsTooLong = !isLoginLengthGood(&login)
	if loginIsTooLong {

		error_msg = "Login is too long."
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		*result = false
		return ret_err
	}

	// 4. Check empty regDate
	regDateIsEmpty = (regDate == 0)
	if regDateIsEmpty {

		regDate = time.Now().Unix()
	}

	// Create Object of UserModel
	user.uuid = uuid
	user.regDate = regDate
	user.login = login

	// Fill Task
	rcvChan = make(chan UserTask)
	task.sender = rcvChan
	task.user = user
	task.result = false
	task.err = nil

	// Send Task
	ModifiedUsersChan <- task

	// Get Feedback
	feedback = <-rcvChan

	// Return
	*result = feedback.result
	return feedback.err
}

//------------------------------------------------------------------------------
