// client.go

/*

	Simple HTTPS Client Test.

	Version: 0.1.
	Date of Creation: 2018-01-26.
	Author: McArcher.

	This is a simple testing Client for HTTPS Protocol Connection.
	It gets the Resource (Page) at the given Address (Host and Port).
	Validation of Certificates can be disabled.
	Start with '-h' Argument to see Help.

*/

//------------------------------------------------------------------------------

package main

import (
	"crypto/tls"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

//------------------------------------------------------------------------------

const BUFFER_SIZE = 1024

const CLIENT_HOST_DEFAULT = "localhost"
const CLIENT_PATH_DEFAULT = "/"
const CLIENT_PORT_DEFAULT = "443"
const CLIENT_PROTOCOL = "https://"
const CLIENT_PROTOCOL_METHOD = "GET"
const CLIENT_PROTOCOL_VERSION_FULL = "HTTP/1.1"
const CLIENT_USER_AGENT = "Mozilla/5.0" +
	" (X11; Ubuntu; Linux x86_64; rv:58.0) Gecko/20100101 Firefox/58.0"

const NL = "\r\n"

const OUTPUT_FILE_PATH_DEFAULT = "out.txt"
const TLS_VERIFICATION_DEFAULT = "on"
const TLS_VERIFICATION_ON = "on"
const TLS_VERIFICATION_OFF = "off"

//------------------------------------------------------------------------------

var cla_host *string
var cla_file_out *string
var cla_path *string
var cla_port *string
var cla_tls_verify *string

var client_connection_host string
var client_connection_path string
var client_connection_port string
var client_connection_type string
var client_connection_address string

var request_uri string

var tls_config *tls.Config
var tls_connection *tls.Conn
var tls_verification_str string
var tls_verification bool

var file_out string
var file_out_exists bool

//------------------------------------------------------------------------------

func main() {

	var client *http.Client
	var err error
	var header http.Header
	var reply *http.Response
	var reply_body []byte
	var request *http.Request
	var tls_config *tls.Config
	var transport *http.Transport

	// Command Line Arguments.
	read_cla()

	// Various Settings.
	log.SetFlags(log.Lshortfile)

	// Connection Type.
	client_connection_type = "tcp"

	// Address of the requested Resource.
	client_connection_address =
		client_connection_host + ":" + client_connection_port

	request_uri =
		CLIENT_PROTOCOL + client_connection_address + client_connection_path

	// TLS Configuration.
	tls_config = &tls.Config{

		ServerName:         client_connection_host,
		InsecureSkipVerify: false,
	}
	if tls_verification == false {

		(*tls_config).InsecureSkipVerify = true
	}

	// HTTP Transport.
	transport = &http.Transport{

		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		TLSClientConfig:    tls_config,
	}

	// HTTP Client.
	client = &http.Client{Transport: transport}

	// Prepare Request.
	request, err = http.NewRequest(CLIENT_PROTOCOL_METHOD, request_uri, nil)
	request.Proto = CLIENT_PROTOCOL_VERSION_FULL // Works?
	header = request.Header
	header.Add("User-Agent", CLIENT_USER_AGENT)

	// Send Request & get Reply.
	reply, err = client.Do(request)
	if err != nil {

		log.Println(err) //
		return
	}
	defer reply.Body.Close()
	reply_body, err = ioutil.ReadAll(reply.Body)

	// Save Reply to File.
	file_out_exists = file_exists(file_out)
	if file_out_exists {

		log.Println("Output File already exists. Skipping the Output.") //

	} else {

		err = ioutil.WriteFile(file_out, reply_body, 0644)
	}
}

//------------------------------------------------------------------------------

// Read Command Line Arguments (Keys, Flags, Switches).

func read_cla() {

	// Set Rules.
	cla_host = flag.String("host", CLIENT_HOST_DEFAULT,
		"Host Name.")
	cla_port = flag.String("port", CLIENT_PORT_DEFAULT,
		"Port Number.")
	cla_path = flag.String("path", CLIENT_PATH_DEFAULT,
		"Resource Path.")
	cla_file_out = flag.String("of", OUTPUT_FILE_PATH_DEFAULT,
		"Output File.")
	cla_tls_verify = flag.String("tlsv", TLS_VERIFICATION_DEFAULT,
		"TLS Verification.")

	// Read C.L.A.
	flag.Parse()

	// Connection Parameters.
	client_connection_host = *cla_host
	client_connection_port = *cla_port
	client_connection_path = *cla_path

	// TLS.
	tls_verification_str = *cla_tls_verify
	tls_verification = true
	if tls_verification_str == TLS_VERIFICATION_ON {

		tls_verification = true

	} else if tls_verification_str == TLS_VERIFICATION_OFF {

		tls_verification = false
	}

	// Output File.
	file_out = *cla_file_out
}

//------------------------------------------------------------------------------

// File exists ?

func file_exists(filepath string) bool {

	var err error

	_, err = os.Stat(filepath)

	if err != nil {

		if os.IsNotExist(err) {

			return false
		}
	}

	return true
}

//------------------------------------------------------------------------------
