// client.go

/*

	HTTPS Client Test...

	Version: 0.1.
	Date of Creation: 2018-01-00.
	Author: McArcher.

	...

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
	"strings"
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

const FOLDER_SEPARATOR = "/"
const FOLDER_DATA = "data"
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

var out_file string
var out_folder string
var out_path string
var out_path_exists bool

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
	header.Add("Accept", "*/*")
	header.Add("Accept-Encoding", "identity")
	header.Add("Connection", "*/*")
	header.Add("User-Agent", "Keep-Alive")

	// Output File.
	out_folder = FOLDER_DATA + FOLDER_SEPARATOR + client_connection_host
	out_file = strings.Replace(client_connection_path,
		FOLDER_SEPARATOR, "_", -1)
	out_path = out_folder + FOLDER_SEPARATOR + out_file
	out_path_exists = file_exists(out_path)
	if out_path_exists {

		log.Println("Output File already exists. Quitting.") //
		return
	}

	// Send Request & get Reply.
	reply, err = client.Do(request)
	check_error(err)
	defer reply.Body.Close()
	reply_body, err = ioutil.ReadAll(reply.Body)

	// Save Reply to File.
	err = os.MkdirAll(out_folder, os.ModePerm)
	check_error(err)
	err = ioutil.WriteFile(out_path, reply_body, 0644)
	check_error(err)
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

func check_error(err error) {

	if err != nil {

		log.Println(err) //

		os.Exit(1)
	}
}

//------------------------------------------------------------------------------
