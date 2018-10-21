// tcp_server_tester.go.

// Stress Load Tester.

package main

import "fmt"
import "log"
import "net"

//import "strconv"
import "sync"

var Request []byte
var wg sync.WaitGroup
var ConnectionAddress string

func main() {

	const PORT_DELIMITER = ":"
	const BUFFER_SIZE = 1024
	const NL = "\r\n"

	var Request_ba []byte
	var Request_str string
	var ConnectionHost = "localhost"
	var ConnectionPort = "2000"
	var i int64

	fmt.Println("TCP Server Tester.")

	// Prepare Request
	Request_str = ""
	for i = 0; i < 1024; i++ {
		Request_str = Request_str + "This is a Test. " // 16 Bytes.
	}
	Request_ba = []byte(Request_str)
	Request = []byte{0, 127, 255} // This is Bad!
	Request = append(Request, Request_ba...)
	//fmt.Println(Request_str) //

	ConnectionAddress = ConnectionHost + PORT_DELIMITER + ConnectionPort

	// Start Stress Test.
	fmt.Print("[START] ")

	// 32 "Threads".
	for i = 0; i < 32; i++ {
		wg.Add(1)
		go makeSpam()
	}

	wg.Wait()
	fmt.Println("[STOP]")
}

func makeSpam() {

	var i int64
	var iMax int64
	var Connection net.Conn
	var RequestLength int
	var err error

	iMax = 100

	for i = 0; i < iMax; i++ {

		// Connect.
		Connection, err = net.Dial("tcp", ConnectionAddress)
		if err != nil {
			log.Println("Error. Can not connect with", ConnectionAddress) //
			return
		}

		// Send something.
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

		// Quit.
		Connection.Close()
	}

	fmt.Print(".")

	wg.Done()
}
