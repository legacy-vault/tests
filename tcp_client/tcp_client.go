// tcp_client.go

/*

	Simple TCP / HTTP Client Test.

	Version: 0.1.
	Date of Creation: 2017-05-30.
	Author: McArcher.

	This is a simple testing Client for TCP / HTTP Protocol Connection.

	...

*/

//------------------------------------------------------------------------------

package main

import (
	"log"
	"net"
)

//------------------------------------------------------------------------------

func main() {

	const PORT_DELIMITER = ":"
	const BUFFER_SIZE = 1024
	const NL = "\r\n"

	var Request, ResponseBuffer []byte
	var Request_str string
	var Connection net.Conn
	var err error
	var ConnectionType = "tcp"
	var ConnectionHost = "ya.ru"
	var ConnectionPort = "80"
	var ConnectionAddress string
	var RequestLength, ResponseLength int

	// Connect
	ConnectionAddress = ConnectionHost + PORT_DELIMITER + ConnectionPort
	Connection, err = net.Dial(ConnectionType, ConnectionAddress)
	if err != nil {
		log.Println("Error. Can not connect with", ConnectionAddress) //
		return
	}

	// Prepare Request
	Request_str = "GET / HTTP/1.1" + NL +
		"Host: ya.ru" + NL +
		"Accept: text/plain" + NL +
		"Accept-Charset: utf-8" + NL +
		NL

	Request = []byte(Request_str)
	log.Println(Request_str) //

	// Send Request
	RequestLength, err = Connection.Write(Request)
	if err != nil {

		log.Println("Error. Can not send Request.") //
		Connection.Close()
		return
	}
	if RequestLength == 0 {

		log.Println("Error. Request is empty.") //
		Connection.Close()
		return
	}

	// Get Response
	ResponseBuffer = make([]byte, BUFFER_SIZE)
	ResponseLength, err = Connection.Read(ResponseBuffer)
	if err != nil {

		log.Println("Error. Can not read Response.") //
		Connection.Close()
		return
	}
	if ResponseLength == 0 {

		log.Println("Error. Empty Response.") //
		Connection.Close()
		return
	}

	// Show Response
	log.Println(string(ResponseBuffer)) //

	// Close Connection
	Connection.Close()
}
