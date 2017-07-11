// rpc_json_client_1.go

// Version: 0.4.
// Date: 2017-07-08.
// Author: McArcher.

package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strconv"
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

type AddResult bool

type GetResult struct {
	UUID    uint64
	LOGIN   string
	REGDATE int64 // UNIX Timestamp
}

type ModifyResult bool

//------------------------------------------------------------------------------
// Methods
//------------------------------------------------------------------------------

func main() {

	// Floods the Server with 'User.Add' Requests.

	var err error
	var srv_conntype, srv_host, srv_port, srv_addr string
	//var srv_path string
	var client *rpc.Client
	var args Arguments
	var result_add AddResult
	var serviceName, methodName_Add, funcName string
	var connection net.Conn
	var i, i_first, i_last, i_step, i_period uint64
	var time_start, time_end, time_1, time_2 time.Time
	var time_duration, time_3 time.Duration

	srv_conntype = "tcp"
	srv_host = "0.0.0.0"
	srv_port = "3000"
	srv_addr = srv_host + ":" + srv_port
	//srv_path = "/"

	// Connect to RPC Server
	connection, err = net.Dial(srv_conntype, srv_addr)
	if err != nil {
		log.Fatalf("Error. Can not connect to %s. %s\r\n", srv_addr, err) //dbg
	}
	defer connection.Close()

	// Client
	client = jsonrpc.NewClient(connection)

	// Prepare Call
	serviceName = "User"
	methodName_Add = "Add"
	funcName = serviceName + "." + methodName_Add
	args.REGDATE = 0

	i_first = 1
	i_last = 1000 * 1000 // 1 M
	i_step = 1
	i_period = 1000 * 10

	i = i_first
	time_1 = time.Now() // Time of the Start of the whole Test
	time_start = time_1 // Time of the Start of a Series of Requests

	for {

		if i > i_last {
			break
		}

		if i%i_period == 0 {

			time_end = time.Now()
			time_duration = time_end.Sub(time_start)
			time_start = time_end

			fmt.Println("i=", i, "duration=", time_duration.Seconds())
		}

		// Configure Call Parameters
		args.UUID = i
		args.LOGIN = "Login #" + strconv.FormatUint(i, 10)

		// Call remote Procedure
		err = client.Call(funcName, args, &result_add)

		// Show Info

		//fmt.Println("[", i, "] Arguments:", args) //dbg
		if err != nil {
			fmt.Printf("(!) Error. %s\r\n", err.Error()) //dbg
		}
		//fmt.Print("Results:", result_add, "\r\n\r\n") //dbg

		// Next i
		i = i + i_step
	}

	time_2 = time.Now() // Time of the End of the whole Test
	time_3 = time_2.Sub(time_1)
	fmt.Println("Whole Test Execution Time:", time_3.Seconds())

}
