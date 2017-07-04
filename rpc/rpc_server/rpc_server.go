// rpc_server.go

package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
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

func (t *Arithmetic) Multiply(args *Arguments, res *Result) error {

	*res = Result(args.A * args.B)

	return nil
}

//------------------------------------------------------------------------------

func main() {

	var srv *rpc.Server
	var err error
	var arith *Arithmetic
	var listener net.Listener
	var srv_conntype, srv_host, srv_port, srv_addr, srv_path string
	var srv_debugPath string

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
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatalf("Serve Error. %s\r\n", err) //dbg
	}
}

//------------------------------------------------------------------------------
