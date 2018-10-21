// server.go

// Version: 0.1.
// Date: 2017-07-06.
// Author: McArcher.

package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//------------------------------------------------------------------------------
// Variables
//------------------------------------------------------------------------------

// Server
var srv *rpc.Server
var codec rpc.ServerCodec
var listener net.Listener
var connection net.Conn
var srv_conntype, srv_host, srv_port, srv_addr, srv_path string
var srv_debugPath string

var debug_enabled bool = true

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

func main() {

	var ok bool
	var err error

	// Initialize Models & DataBase
	Users.Init()
	db_init()

	// Load Data from DB
	ok, err = Users.Load()
	if !ok {

		panic(err.Error())
	}
	log.Printf("The DataBase has been loaded.\r\n") //

	// Save Manager
	go manager()

	// Debug
	if debug_enabled {
		go debugger()
	}

	// Start Server & Service
	server_configure()
	ok, err = server_start()
	if !ok {

		panic(err.Error())
	}
	log.Printf("Started RPC Handler at %s.\r\n", srv_addr) //

	// Serve
	for {

		connection, err = listener.Accept()
		if err != nil {

			panic(err.Error())
		}

		codec = jsonrpc.NewServerCodec(connection)

		go srv.ServeCodec(codec)
	}
}

//------------------------------------------------------------------------------

func server_configure() {

	// Sets Server's Parameters.

	srv_conntype = "tcp"
	srv_host = "0.0.0.0"
	srv_port = "3000"
	srv_addr = srv_host + ":" + srv_port
	srv_path = "/"
	srv_debugPath = "/debug"
}

//------------------------------------------------------------------------------

func server_start() (bool, error) {

	// Starts the Server & JSON-RPC Service.

	var err error
	var error_msg string
	var ret_err error

	// Create Server, register Service
	srv = rpc.NewServer()
	service_user = new(User)
	err = srv.Register(service_user)
	if err != nil {

		error_msg = fmt.Sprintf("Can not register Service. %s", err.Error())
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}

	// Handle, listen
	srv.HandleHTTP(srv_path, srv_debugPath)
	listener, err = net.Listen(srv_conntype, srv_addr)
	if err != nil {

		error_msg = fmt.Sprintf("Error. Can not listen on %s. %s", srv_addr, err.Error())
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}

	return true, nil
}

//------------------------------------------------------------------------------
