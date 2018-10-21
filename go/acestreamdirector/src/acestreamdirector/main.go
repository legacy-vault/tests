// acestreamdirector/main.go

package main

//------------------------------------------------------------------------------

import "acestreamdirector/config"
import "acestreamdirector/client"
import "log"
import "net"
import "os"

//------------------------------------------------------------------------------

//	Ace Stream Media API :
//		http://wiki.acestream.org/wiki/index.php/Engine_API
//
//	Notes:
//		All this does not work, the Forum of 'Ace Stream Media' is silent and
//		Support is not provided. This Program is just a Test of a non-working
//		API.

//------------------------------------------------------------------------------

var Connection *net.TCPConn

//------------------------------------------------------------------------------

// Main Function.
func main() {

	var err error

	config.Init()

	Connection, err = client.Connect(&config.ConnAddress, config.Verbose)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = client.DisConnect(Connection, config.Verbose)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

//------------------------------------------------------------------------------
