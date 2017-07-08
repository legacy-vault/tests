// server.go

// Version: 0.4.
// Date: 2017-07-08.
// Author: McArcher.

package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"sync"
	"time"
)

//------------------------------------------------------------------------------
// Variables
//------------------------------------------------------------------------------

// Server
var srv *rpc.Server
var codec rpc.ServerCodec
var listener *net.TCPListener
var connection net.Conn
var srv_conntype, srv_host, srv_port, srv_addr, srv_path string
var srv_netAddr *net.TCPAddr
var srv_debugPath string
var server_ping_interval = 1      // Seconds
var server_listen_interval = 5    // Seconds
var server_rpc_wait_interval = 30 // Seconds

var EmergencyShutdownIsNeeded bool   // Raised by OSSignalHandler after a Signal from O.S.
var EmergencySaveIsNeeded bool       // Raised by Server after Server's Shutdown
var RequestsWaitGroup sync.WaitGroup // Number of Requests to the Server being processed now

/*
	To make a graceful Shutdown we need to:

	1.	Stop accepting new Connections.

	2.	For new Requests in existing Connections return an Error & close
		these Connection.

	3.	Save all new Results from Memory to File.

	4.	Guarantee that Results of current Requests are properly delivered to
		the Clients.

	---

	The 4-th Step we can not do while we do not control the RPC Server's
	Engine, which is a Part of Golang. At least, we can not do it in this
	Implementation of the RPC Server. We can either create our own RPC
	Server (or modify an existing one) or invent some other Things.

*/

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

func main() {

	var pingInterval time.Duration
	var listenInterval time.Duration
	var rpcWaitInterval time.Duration
	var net_err *net.OpError

	var ok bool
	var err error

	// Initialize
	Users.Init()
	db_init()
	server_init(&pingInterval, &listenInterval, &rpcWaitInterval)

	// Load Data from DB
	ok, err = Users.Load()
	if !ok {

		panic(err.Error())
	}
	log.Printf("The DataBase has been loaded.\r\n") //

	// User Manager
	go UserManager()
	go ManagerPinger()

	// O.S. Signal Handler
	go OSSignalHandler()

	// Debugger
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

		if EmergencyShutdownIsNeeded {

			// Stop accepting new Connections
			break
		}

		listener.SetDeadline(time.Now().Add(listenInterval))
		connection, err = listener.Accept()

		if err != nil {

			// Deadline?
			net_err, ok = err.(*net.OpError)
			if ok && net_err.Timeout() {

				continue
			}

			// Other Error
			log.Println(net_err.Error()) //
			continue
		}

		log.Printf("Accepted Connection from [%s].\r\n", connection.RemoteAddr().String()) //

		codec = jsonrpc.NewServerCodec(connection)

		go srv.ServeCodec(codec)
	}

	// Emergency Shutdown
	listener.Close()
	log.Println("Stopped accepting new Connections.") //

	// While the RequestsWaitGroup is changed not by the RPC Server itself,
	// but by the RPC Actions (Functions), this Method of getting the Number
	// of running Requests is not precise. The Reason is simple: when the RPC
	// Action calls 'RequestsWaitGroup.Done()' it needs some additional Time
	// to return Results to its Caller (internal Engine of the RPC Server)
	// and the Caller itself must return Results to the Server's Client.
	// While this Program uses built-in Golang's RPC Server, we can not
	// control how the Server works. To make Everything well, we should
	// manipulate the RequestsWaitGroup in the Code of RPC Server. So, as a
	// temporary (and not very reliable) Solution there is a Way to correct
	// such Behaviour by waitng some additional Time after the RPC Action has
	// done Manipulations with RequestsWaitGroup. Of course, it is a bad
	// Practice, but it is better than Nothing. I hope that in future Versions
	// of Golang Google will add 'Stop' and 'Shutdown' Methods to their RPC
	// Server.

	// Wait for existing Requests to "finish"
	RequestsWaitGroup.Wait()
	log.Println("Waiting", server_rpc_wait_interval, "Seconds for the RPC Actions to finish.") //
	time.Sleep(rpcWaitInterval)

	// Start Emergency Save
	EmergencySaveIsNeeded = true

	// Wait for Emergency Save to complete
	for EmergencySaveHasFinished == false {

		time.Sleep(pingInterval)
	}

	// Quit to O.S.
	os.Exit(0)
}

//------------------------------------------------------------------------------

func server_configure() {

	// Sets Server's Parameters.

	var IP net.IP

	srv_conntype = "tcp"
	srv_host = "0.0.0.0"
	srv_port = "3000"
	srv_addr = srv_host + ":" + srv_port

	IP = net.ParseIP(srv_host)
	if IP == nil {
		log.Println("Bad IP Address.") //
	}

	srv_netAddr = new(net.TCPAddr)
	srv_netAddr.IP = IP
	srv_netAddr.Port = 3000
	srv_path = "/"
	srv_debugPath = "/debug"
}

//------------------------------------------------------------------------------

func server_init(pingInterval, listenInterval, rpcWaitInterval *time.Duration) {

	// Initializes the Server.

	EmergencyShutdownIsNeeded = false
	EmergencySaveIsNeeded = false

	*pingInterval = time.Second * time.Duration(server_ping_interval)
	*listenInterval = time.Second * time.Duration(server_listen_interval)
	*rpcWaitInterval = time.Second * time.Duration(server_rpc_wait_interval)
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
	//listener, err = net.Listen(srv_conntype, srv_addr)
	listener, err = net.ListenTCP(srv_conntype, srv_netAddr)
	if err != nil {

		error_msg = fmt.Sprintf("Error. Can not listen on %s. %s", srv_addr, err.Error())
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}

	return true, nil
}

//------------------------------------------------------------------------------
