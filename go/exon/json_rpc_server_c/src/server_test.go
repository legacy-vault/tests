// server_test.go

// Version: 0.3.
// Date: 2017-07-07.
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

	// This Test just simulates (internally) Requests & shows some Results.
	// Was used in early Versions.

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
		debug_db_print_recordsmap()       // dbg
		debug_db_print_addedusersmap()    // dbg
		debug_db_print_modifiedusersmap() // dbg
		fmt.Print("\r\n\r\n")             // dbg

		if i > 11 {
			break
		}

		// Delay
		time.Sleep(sleep_timer)
	}
}

//------------------------------------------------------------------------------

func test_server_2(t *testing.T) {

	fmt.Println("\r\nTest #28. Test the Server.") //

	/*
		/!\ This Test is NOT automatical /!\
	*/

	// This Test simulates (internally) a lot of Requests.
	// Was used in early Versions.

	main()

	var i, i_max, i_step int
	var j int
	var user UserModel
	var sleep_timer time.Duration

	sleep_timer = time.Second * 4

	i = 1
	i_max = 10
	i_step = 1

	for {

		if i > i_max {
			break
		}

		// Add User
		user.login = "Login A-" + strconv.Itoa(i)
		user.regDate = time.Now().Unix()
		user.uuid = uint64(i)
		Users.AddAsRequest(user)

		j = i % 5

		if j == 0 { // Every 5-th Reuest

			// Modifies the User's login
			user.login = "Login B-" + strconv.Itoa(i)
			Users.ModifyAsRequest(user)
		}

		// Print
		Users.debug_Print()               // dbg
		debug_PrintLoginsMap()            // dbg
		debug_db_print_recordsmap()       // dbg
		debug_db_print_addedusersmap()    // dbg
		debug_db_print_modifiedusersmap() // dbg
		fmt.Print("\r\n\r\n")             // dbg

		// Delay
		time.Sleep(sleep_timer)

		// Next i
		i = i + i_step
	}
}

//------------------------------------------------------------------------------
