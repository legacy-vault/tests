// server_test.go

// Version: 0.1.
// Date: 2017-07-06.
// Author: McArcher.

package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

//------------------------------------------------------------------------------
// Variables
//------------------------------------------------------------------------------

// ...

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

func test_server_1(t *testing.T) {

	fmt.Println("\r\nTest #27. Test the Server.") //

	/*
		/!\ This Test is NOT automatical /!\
	*/

	main()

	var i int
	var j int
	var user UserModel
	var sleep_timer time.Duration

	sleep_timer = time.Second * 4

	i = 1
	j = 1

	for {

		if i%5 == 0 {

			// Modify User
			user.login = "Login B-" + strconv.Itoa(j)
			user.regDate = time.Now().Unix()
			user.uuid = uint64(j*1000 + j)
			Users.ModifyAsRequest(user)
			//
			i = i + 1
			j = j + 1

		} else {

			// Add User
			user.login = "Login A-" + strconv.Itoa(i)
			user.regDate = time.Now().Unix()
			user.uuid = uint64(i*1000 + i)
			Users.AddAsRequest(user)
			//
			i = i + 1
		}

		// Print
		Users.debug_Print()               // dbg
		db_debug_print_recordsmap()       // dbg
		db_debug_print_addedusersmap()    // dbg
		db_debug_print_modifiedusersmap() // dbg
		fmt.Print("\r\n\r\n")             // dbg

		if i > 11 {
			break
		}

		// Delay
		time.Sleep(sleep_timer)
	}
}

//------------------------------------------------------------------------------
