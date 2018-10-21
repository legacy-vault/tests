// service_user.go

// Version: 0.1.
// Date: 2017-07-06.
// Author: McArcher.

package main

//------------------------------------------------------------------------------
// Types
//------------------------------------------------------------------------------

type User int // Used as RPC Service called 'User'

type Arguments struct {
	UUID    uint64
	LOGIN   string
	REGDATE int64 // UNIX Timestamp
}

type AddResult bool

type GetResult struct {
	UUID    uint64
	LOGIN   string
	REGDATE int64 // UNIX Timestamp
}

type ModifyResult bool

//------------------------------------------------------------------------------
// Variables
//------------------------------------------------------------------------------

// Service
var service_user *User

//------------------------------------------------------------------------------
// Methods
//------------------------------------------------------------------------------

func (u *User) Add(args *Arguments, result *AddResult) error {

	// RPC Action: 'User.Add'.

	var uuid uint64
	var regDate int64
	var login string
	var user UserModel

	var ok bool
	var err error
	var ret_err error

	// Read Arguments
	uuid = args.UUID
	regDate = args.REGDATE
	login = args.LOGIN

	// Preparations & preliminary Checks
	// ...

	// Create Object of UserModel
	user.uuid = uuid
	user.regDate = regDate
	user.login = login

	// Call Model's Method
	ok, err = Users.AddAsRequest(user)
	if !ok {

		// Error
		*result = false
		ret_err = err

		return ret_err
	}

	// No Errors
	*result = true

	return nil
}

//------------------------------------------------------------------------------

func (u *User) Get(args *Arguments, result *GetResult) error {

	// 'User.Get' RPC Action.

	var uuid uint64
	var regDate int64
	var login string
	var user UserModel
	var reply UserModel

	var ok bool
	var err error
	var ret_err error

	// Read Arguments
	uuid = args.UUID
	regDate = args.REGDATE
	login = args.LOGIN

	// Preparations & preliminary Checks
	// ...

	// Create Object of UserModel
	user.uuid = uuid
	user.regDate = regDate
	user.login = login

	// Call Model's Method
	ok, err = Users.Get(user, &reply)
	if !ok {

		// Error
		result.UUID = 0
		result.REGDATE = 0
		result.LOGIN = ""
		ret_err = err

		return ret_err
	}

	// No Errors
	result.UUID = reply.uuid
	result.REGDATE = reply.regDate
	result.LOGIN = reply.login

	return nil
}

//------------------------------------------------------------------------------

func (u *User) Modify(args *Arguments, result *ModifyResult) error {

	// 'User.Modify' RPC Action.

	var uuid uint64
	var regDate int64
	var login string
	var user UserModel

	var ok bool
	var err error
	var ret_err error

	// Read Arguments
	uuid = args.UUID
	regDate = args.REGDATE
	login = args.LOGIN

	// Preparations & preliminary Checks
	// ...

	// Create Object of UserModel
	user.uuid = uuid
	user.regDate = regDate
	user.login = login

	// Call Model's Method
	ok, err = Users.ModifyAsRequest(user)
	if !ok {

		// Error
		*result = false
		ret_err = err

		return ret_err
	}

	// No Errors
	*result = true

	return nil
}

//------------------------------------------------------------------------------
