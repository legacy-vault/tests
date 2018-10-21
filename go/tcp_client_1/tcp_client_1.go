// tcp_client_1.go

/*

	Simple TCP Client Test.

	Version: 0.1.
	Date of Creation: 2017-06-00.
	Author: McArcher.

	...

*/

//------------------------------------------------------------------------------

package main

import (
	"crypto/rand"
	"fmt"
	//"log"
	"math/big"
	"net"
	"os"
	"strings"
	"time"
)

//------------------------------------------------------------------------------

const PORT_DELIMITER = ":"
const BUFFER_SIZE = 1024
const NL = "\r\n"

//------------------------------------------------------------------------------

var Tasks uint
var GoodServers, BadServers []string
var ResultsFile string
var SaveChan chan string

//------------------------------------------------------------------------------

func main() {

	var Request []byte
	var Request_str string
	var ConnectionType = "tcp"
	var ConnectionHost string
	var ConnectionPort string
	var ConnectionPorts []string
	var rnd_a, rnd_b, rnd_c, rnd_d, rnd_max *big.Int
	var ConnectionTimeout, tick time.Duration
	var delay_1 time.Duration
	var ResponseMarker string
	//var iter_num, iter_i uint64

	Tasks = 0
	//iter_num = 1000 * 1000 // 1 M
	tick = time.Second * 1
	ConnectionTimeout = time.Second * 5
	delay_1 = time.Microsecond * 10
	ResultsFile = "results.html"
	SaveChan = make(chan string)

	go saveManager()

	ResponseMarker = "HTTP/1.1 302 Found"
	Request_str = "GET http://ya.ru/ HTTP/1.1" + NL +
		"Host: ya.ru" + NL +
		"Accept: text/plain" + NL +
		"Accept-Charset: utf-8" + NL +
		NL
	Request = []byte(Request_str)
	rnd_max = big.NewInt(256)
	ConnectionPorts = make([]string, 3)
	ConnectionPorts[0] = "80"
	ConnectionPorts[1] = "3128"
	ConnectionPorts[2] = "8080"

	//fmt.Println(Request_str) //

	//for iter_i = 1; iter_i <= iter_num; iter_i++ {
	for {

		// Generate Random IPA
		// Rnd IPA
		rnd_a, _ = rand.Int(rand.Reader, rnd_max)
		rnd_b, _ = rand.Int(rand.Reader, rnd_max)
		rnd_c, _ = rand.Int(rand.Reader, rnd_max)
		rnd_d, _ = rand.Int(rand.Reader, rnd_max)

		ConnectionHost = rnd_a.String() + "." + rnd_b.String() + "." +
			rnd_c.String() + "." + rnd_d.String()

		/*
			if iter_i == iter_num {
				ConnectionHost = "74.208.150.246" //
			}
		*/
		//fmt.Println("[", iter_i, "]") //

		// Iterate Ports
		for _, v := range ConnectionPorts {

			ConnectionPort = v
			Tasks++

			go examineHost(ConnectionHost, ConnectionPort, ConnectionType,
				Request, ConnectionTimeout, ResponseMarker)

			//fmt.Print(" ")
			time.Sleep(delay_1)
		}

	}

	// Wait for Results
	for Tasks > 0 {

		time.Sleep(tick)
	}

	// Show Results
	/*
		fmt.Println("\r\n\r\nBad Servers:")
		for _, v := range BadServers {

			fmt.Println(v)
		}
	*/
	fmt.Println("\r\nGood Servers:")
	for _, v := range GoodServers {

		fmt.Println(v)
	}
}

//------------------------------------------------------------------------------

func examineHost(host, port, connType string,
	req []byte, timeout time.Duration, marker string) {

	var ConnectionAddress string
	var Connection net.Conn
	var err error
	var RequestLength, ResponseLength int
	var ResponseBuffer []byte
	var isProxy bool
	var saveText string

	// Connect
	ConnectionAddress = host + PORT_DELIMITER + port
	Connection, err = net.DialTimeout(connType, ConnectionAddress, timeout)
	if err != nil {
		//fmt.Println("FAIL", ConnectionAddress) //
		//fmt.Print(".") //
		BadServers = append(BadServers, ConnectionAddress)
		Tasks--
		return
	}

	// Send Request
	RequestLength, err = Connection.Write(req)
	if err != nil {

		//log.Println("Error. Can not send Request.") //
		//fmt.Print(".") //
		Connection.Close()
		Tasks--
		return
	}
	if RequestLength == 0 {

		//log.Println("Error. Request is empty.") //
		//fmt.Print(".") //
		Connection.Close()
		Tasks--
		return
	}

	// Get Response
	ResponseBuffer = make([]byte, BUFFER_SIZE)
	ResponseLength, err = Connection.Read(ResponseBuffer)
	if err != nil {

		//log.Println("Error. Can not read Response.") //
		//fmt.Print(".") //
		Connection.Close()
		Tasks--
		return
	}
	if ResponseLength == 0 {

		//log.Println("Error. Empty Response.") //
		//fmt.Print("0") //
		Connection.Close()
		Tasks--
		return
	}

	// Search Marker
	isProxy = strings.Contains(string(ResponseBuffer), marker)
	if isProxy {
		GoodServers = append(GoodServers, ConnectionAddress)
		fmt.Print("[", ConnectionAddress, "]") //
		saveText = "<a target='_blank' href='http://" + ConnectionAddress +
			"'>" + ConnectionAddress + "</a><br>\r\n"
		SaveChan <- saveText
		//fmt.Print("|")
	} else {
		//fmt.Println("BAD ", ConnectionAddress) //
		//fmt.Print(".") //
		BadServers = append(BadServers, ConnectionAddress)
	}

	// Close Connection
	Connection.Close()
	Tasks--
}

//------------------------------------------------------------------------------

func saveManager() {

	var file *os.File
	var err error
	var data string

	file, err = os.OpenFile(ResultsFile, os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for {

		data = <-SaveChan

		_, err = file.WriteString(data)
		if err != nil {
			panic(err)
		}
	}
}

//------------------------------------------------------------------------------
