// server_tester.go

/*

	Simple Test for Simple Queue Server.

	Version: 0.1.
	Date of Creation: 2017-06-01.
	Author: McArcher.

	This is a simple testing Client for TCP / HTTP Protocol Connection.

	This Program has been written to test the 'Simple Queue Server'.

*/

//------------------------------------------------------------------------------

package main

import (
	"fmt"
	"net"
	"time"
)

//------------------------------------------------------------------------------

const BUFFER_SIZE = 1024

var ConnectionType string

//------------------------------------------------------------------------------

func main() {

	const PORT_DELIMITER = ":"
	const NL = "\r\n"

	var Request []byte
	var Connection net.Conn
	var ConnectionHost = "127.0.0.1"
	var ConnectionPort = "5555"
	var ConnectionAddress string
	var ValueString string
	var ValueBA []byte
	var SleepTime int64

	// Connection Parameters
	ConnectionType = "tcp"
	ConnectionAddress = ConnectionHost + PORT_DELIMITER + ConnectionPort

	// Connect
	connect(&Connection, ConnectionType, ConnectionAddress)

	// Size of Queue
	Request = []byte("s")
	send_request(Connection, &Request)
	get_response(Connection)

	/*
		// Simulate bad 'S' Request
		Request = []byte("s123")
		send_request(Connection, &Request)
		get_response(Connection)

		// Simulate bad 'G' Request
		Request = []byte("gXYZ")
		send_request(Connection, &Request)
		get_response(Connection)

		// Simulate bad 'P' Request
		Request = []byte("p#%^*@#")
		send_request(Connection, &Request)
		get_response(Connection)
	*/

	// Put an Item
	ValueString = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	putString(&ValueString, Connection)

	// Simulate a slow Network or a slow Machine
	SleepTime = 5
	fmt.Println("Now I am idle for", SleepTime, "seconds...") //
	time.Sleep(time.Second * time.Duration(SleepTime))

	// Put an Item
	ValueBA = make([]byte, 3)
	ValueBA[0] = 0x00
	ValueBA[1] = 0x01
	ValueBA[2] = 0x02
	putByteArray(&ValueBA, 3, Connection)

	// Size of Queue
	Request = []byte("s")
	send_request(Connection, &Request)
	get_response(Connection)

	// Simulate a slow Network or a slow Machine
	SleepTime = 5
	fmt.Println("Now I am idle for", SleepTime, "seconds...") //
	time.Sleep(time.Second * time.Duration(SleepTime))

	// Size of Queue
	Request = []byte("s")
	send_request(Connection, &Request)
	get_response(Connection)

	// Get Item
	Request = []byte("g")
	// # 1
	send_request(Connection, &Request)
	get_response(Connection)
	// # 2
	send_request(Connection, &Request)
	get_response(Connection)
	// # 3
	send_request(Connection, &Request)
	get_response(Connection)

	// Simulate aborted Connection
	disconnect(&Connection)

	// Disconnect by Server
	Request = []byte("q")
	send_request(Connection, &Request)
	get_response(Connection)
}

//------------------------------------------------------------------------------

func connect(Connection *net.Conn, ConnectionType, ConnectionAddress string) {

	// Connects to the Server.

	var err error

	*Connection, err = net.Dial(ConnectionType, ConnectionAddress)
	if err != nil {
		fmt.Println("Error. Can not connect with", ConnectionAddress) //
		return
	}
}

//------------------------------------------------------------------------------

func disconnect(Connection *net.Conn) {

	// Dis-connects from the Server.

	(*Connection).Close()
}

//------------------------------------------------------------------------------

func send_request(Connection net.Conn, Request *[]byte) {

	// Sends a Request.
	// If Error occurs, then closes the Connection.

	var RequestLength int
	var err error

	//============================================
	fmt.Println("\r\nRequest:")   // Show string
	fmt.Println(string(*Request)) // Show string
	fmt.Println(*Request)         // Show raw Bytes
	//============================================

	RequestLength, err = Connection.Write(*Request)
	if err != nil {

		fmt.Println("Error. Can not send Request.") //
		disconnect(&Connection)
		return
	}
	if RequestLength == 0 {

		fmt.Println("Error. Request is empty.") //
		disconnect(&Connection)
		return
	}
}

//------------------------------------------------------------------------------

func get_response(Connection net.Conn) {

	// Gets a Response, shows it, then closes Connection.

	var ResponseBuffer []byte
	var ResponseLength int
	var err error
	var ValueLen int

	// Get Response
	ResponseBuffer = make([]byte, BUFFER_SIZE)
	ResponseLength, err = Connection.Read(ResponseBuffer)
	if err != nil {

		fmt.Println("Error. Can not read Response.") //
		disconnect(&Connection)
		return
	}
	if ResponseLength == 0 {

		fmt.Println("Error. Empty Response.") //
		disconnect(&Connection)
		return
	}

	// Show Response
	//=======================================================================
	//fmt.Println(string(ResponseBuffer)) //
	fmt.Println(">>>> Response Type is (", string(ResponseBuffer[0]), ").") //
	if ResponseBuffer[0] == 's' {

		fmt.Println(ResponseBuffer[:6]) //

	} else if ResponseBuffer[0] == 'i' {

		ValueLen = int(ResponseBuffer[2])*256 + int(ResponseBuffer[3])
		fmt.Println("ValueLen:", ValueLen)       //
		fmt.Println(ResponseBuffer[:5+ValueLen]) //
	} else if ResponseBuffer[0] == 'e' {

		fmt.Println("Warning. The Queue is empty.") //

	} else if ResponseBuffer[0] == 'f' {

		fmt.Println("Warning. The Queue is full.") //

	} else if ResponseBuffer[0] == 'r' {

		fmt.Println("Achtung! Something has gone wrong way :-/") //

	} else if ResponseBuffer[0] == 'k' {

		fmt.Println("LG. Life is Good :-)") //
	}
	//=======================================================================
}

//------------------------------------------------------------------------------

func putString(Value *string, Connection net.Conn) {

	// Puts an Item given as String, shows Response.

	var ValueBA []byte
	var ValueLen, i int
	var Request []byte
	var ValueLen16 uint16

	// Prepare Request
	ValueBA = []byte(*Value)
	ValueLen = len(*Value)
	Request = make([]byte, ValueLen+5)
	Request[0] = 'p'
	Request[1] = ' '
	ValueLen16 = uint16(ValueLen)
	Request[2] = byte(ValueLen16 >> 8)
	Request[3] = byte(ValueLen16 % 256)
	Request[4] = ' '
	for i = 0; i < ValueLen; i++ {
		Request[5+i] = ValueBA[i]
	}

	send_request(Connection, &Request)
	get_response(Connection)
}

//------------------------------------------------------------------------------

func putByteArray(Value *[]byte, Length uint16, Connection net.Conn) {

	// Puts an Item given as String, shows Response.

	var i uint16
	var Request []byte

	// Prepare Request
	Request = make([]byte, Length+5)
	Request[0] = 'p'
	Request[1] = ' '
	Request[2] = byte(Length >> 8)
	Request[3] = byte(Length % 256)
	Request[4] = ' '
	for i = 0; i < Length; i++ {
		Request[5+i] = (*Value)[i]
	}

	send_request(Connection, &Request)
	get_response(Connection)
}

//------------------------------------------------------------------------------
