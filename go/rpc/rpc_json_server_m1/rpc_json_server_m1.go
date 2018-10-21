// rpc_json_server_m1.go

package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	//"strconv"
)

//------------------------------------------------------------------------------
// Types
//------------------------------------------------------------------------------

type Arithmetic bool // Used as RPC Service called 'Arithmetic'

type Arguments struct {
	a     int
	B     int
	NUM   uint64
	UID32 uint32
	UID64 uint64
	TEXT  string
}

type Result struct {
	X float64
	Z string
}

//------------------------------------------------------------------------------
// Methods
//------------------------------------------------------------------------------

func (t *Arithmetic) Multiply(args *Arguments, res *Result) error {

	log.Println("args:", args) ////////////////////////

	//res.X = float64(args.A) * float64(args.B)
	//res.Z = "Hello, " + args.login + " ! Your UID is " + strconv.Itoa(int(args.uid)) + " !"

	return nil
}

//------------------------------------------------------------------------------

func main() {

	var srv *rpc.Server
	var err error
	var arith *Arithmetic
	var listener net.Listener
	var codec rpc.ServerCodec
	var srv_conntype, srv_host, srv_port, srv_addr, srv_path string
	var srv_debugPath string
	var connection net.Conn

	srv_conntype = "tcp"
	srv_host = "0.0.0.0"
	srv_port = "3000"
	srv_addr = srv_host + ":" + srv_port
	srv_path = "/"
	srv_debugPath = "/debug"

	// Create Server, register Service
	srv = rpc.NewServer()
	arith = new(Arithmetic)
	err = srv.Register(arith)
	if err != nil {
		log.Fatalf("Error. Service Format is not correct. %s\r\n", err) //dbg
	}

	// Handle, listen
	srv.HandleHTTP(srv_path, srv_debugPath)
	listener, err = net.Listen(srv_conntype, srv_addr)
	if err != nil {
		log.Fatalf("Error. Can not listen on %s. %s\r\n", srv_addr, err) //dbg
	}
	log.Printf("Started RPC Handler at %s.\r\n", srv_addr) //dbg

	// Serve
	for {

		connection, err = listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		codec = jsonrpc.NewServerCodec(connection)

		go srv.ServeCodec(codec)
	}
}

//------------------------------------------------------------------------------
