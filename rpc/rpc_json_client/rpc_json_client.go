// rpc_json_client.go

package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//------------------------------------------------------------------------------
// Types
//------------------------------------------------------------------------------

type Arithmetic int // Used as RPC Service called 'Arithmetic'

type Arguments struct {
	A int
	B int
}

type Result int

//------------------------------------------------------------------------------
// Methods
//------------------------------------------------------------------------------

func main() {

	var err error
	var srv_conntype, srv_host, srv_port, srv_addr string
	//var srv_path string
	var client *rpc.Client
	var args Arguments
	var result Result
	var serviceName, methodName, funcName string
	var connection net.Conn

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
	serviceName = "Arithmetic"
	methodName = "Multiply"
	funcName = serviceName + "." + methodName
	args.A = 7
	args.B = 8

	// Call remote Procedure
	err = client.Call(funcName, args, &result)
	if err != nil {
		log.Fatalf("Error. %s\r\n", err) //dbg
	}

	// Show Results
	fmt.Printf("[%d; %d] -> [%d].\r\n", args.A, args.B, result) //dbg
}
