// server.go

// Version: 0.1.
// Date: 2017-07-06.
// Author: McArcher.

package main

import (
	"fmt"
	"time"
)

//------------------------------------------------------------------------------

var debug_add_runs int = 0
var debug_db_saved_users int = 0
var debug_db_saves_count int = 0

//------------------------------------------------------------------------------

func debug_db_print_addedusersmap() {

	// Prints the AddedUsers Map.
	// This Function is used for testing Purposes.

	var cur_uuid uint64

	fmt.Print("AddedUsers Map [UUID] :") //dbg

	AddedUsersMapAccess.RLock()
	for cur_uuid, _ = range AddedUsersMap {

		fmt.Printf(" [%d]", cur_uuid) //dbg
	}
	AddedUsersMapAccess.RUnlock()

	fmt.Println("") //dbg
}

//------------------------------------------------------------------------------

func debug_db_print_modifiedusersmap() {

	// Prints the ModifiedUsers Map.
	// This Function is used for testing Purposes.

	var cur_uuid uint64

	fmt.Print("ModifiedUsers Map [UUID] :") //dbg

	ModifiedUsersMapAccess.RLock()
	for cur_uuid, _ = range ModifiedUsersMap {

		fmt.Printf(" [%d]", cur_uuid) //dbg
	}
	ModifiedUsersMapAccess.RUnlock()

	fmt.Println("") //dbg
}

//------------------------------------------------------------------------------

func debug_db_print_recordsmap() {

	// Prints the Records Map.
	// This Function is used for testing Purposes.

	var cur_uuid uint64
	var cur_record_index uint64

	fmt.Println("Records Map.")                          //dbg
	fmt.Println("-------------------------------------") //dbg
	fmt.Println(" [UUID] [Record's Index in File]")      //dbg
	fmt.Println("-------------------------------------") //dbg

	for cur_uuid, cur_record_index = range RecordsMap {

		fmt.Printf("[%d][%d].\r\n", cur_uuid, cur_record_index) //dbg
	}

	fmt.Println("-------------------------------------") //dbg
}

//------------------------------------------------------------------------------

func debugger() {

	// Run-Time simple Debugger

	var debug_interval time.Duration

	debug_interval = time.Second * 1

	fmt.Println("Debugger started.") //

	for {

		time.Sleep(debug_interval)

		fmt.Println("======================================")        //
		fmt.Println("debug_db_saved_users = ", debug_db_saved_users) //
		fmt.Println("len(RecordsMap) = ", len(RecordsMap))           //
		fmt.Println("debug_add_runs = ", debug_add_runs)             //
		fmt.Println("len(Users) = ", len(Users))                     //
		fmt.Println("debug_db_saves_count = ", debug_db_saves_count) //
		fmt.Println("len(AddedUsersMap) = ", len(AddedUsersMap))     //
		fmt.Println("======================================")        //
	}
}

//------------------------------------------------------------------------------
