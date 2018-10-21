// proxy_checker_1.go

/*

	Simple HTTP Proxy Checker.

	Version: 0.1.
	Date of Creation: 2018-01-09.
	Author: McArcher.

	Reads the input File with Servers' Addresses and tries to connect to them.
	Successful Connection Attempts are reported to the external output File.
	To get Command Line Parameters Usage Hint, Run the Program with '-h'
	Argument.

	Addresses are specified as 	Hostnames and Ports, separated by Semicolon ':'
	Symbol. Example: '1.2.3.4:80' or 'proxyserver.net:8080'. Each Line of the
	input File contains only a single Address. Lines of the input File are
	separated by the Combination of CR+LF Symbols (as in Text Files of Microsoft
	Windows).

	In each Connection Attempt the Program tries to use HTTP 'CONNECT' Method
	at 	first, and tries HTTP 'GET' Method afterwards. 'GET' Method tries to
	fetch a remote 	File to compare it with the known existing Sample. If the
	remote File's Contents match the known Sample, then the Connection Attempt
	is considered as successful.

	Usage Example:
		./proxy_checker_1 -if aaa.txt -of bbb.txt -tf ccc.gif

*/

//------------------------------------------------------------------------------

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

//------------------------------------------------------------------------------

// Program's Exit Codes.
const EXIT_CODE_FILE_INPUT_NOT_FOUND int = 1
const EXIT_CODE_FILE_INPUT_SYNTAX_BAD int = 2
const EXIT_CODE_FILE_OUTPUT_NOT_FOUND int = 3
const EXIT_CODE_FILE_TESTER_NOT_FOUND int = 4

// Various Parameters.
const BUFFER_SIZE = 1024 * 1024   // 1 MiB Buffer.
const CONNECTION_TIMEOUT_SEC = 30 // Connection Timeout in Seconds.
const NL = "\r\n"                 // New Line.
const PORT_DELIMITER = ":"        // Hostname-Port Delimiter.
const STATUS_SUCCESS_POSTFIX = "200 OK"
const TIME_FORMAT string = "2006-01-02 15:04:05"

// Command Line Arguments' (C.L.A.) Parameters.
const CLA_PARAM_FILE_INPUT = "if"
const CLA_PARAM_FILE_OUTPUT = "of"
const CLA_PARAM_FILE_TESTER = "tf"

const CLA_FILE_INPUT_DEFAULT = "input.txt"
const CLA_FILE_OUTPUT_DEFAULT = "output.txt"
const CLA_FILE_TESTER_DEFAULT = "sample.file"

// Connection Parameters.
const ConnectMethod string = "CONNECT"
const RequestAccept string = "text/plain"
const RequestAcceptVN string = "Accept"
const RequestAcceptCharset string = "utf-8"
const RequestAcceptCharsetVN string = "Accept-Charset"
const RequestConnectionClose string = "Close"
const RequestConnectionCloseVN string = "Connection"
const RequestHostVN string = "Host"
const RequestHTTPVersion string = "1.1"
const RequestHTTPVersionString string = "HTTP/" + RequestHTTPVersion
const RequestMethod string = "GET"
const RequestParameterValueDelimiter string = ":"
const RequestSpace string = " "
const RequestStatusSuccessPattern string = RequestHTTPVersionString +
	RequestSpace + STATUS_SUCCESS_POSTFIX

// Error Messages.
const MSG_ERROR_1 = "Can not find input file : "
const MSG_ERROR_2 = "Bad syntax in input file at line : "
const MSG_ERROR_3 = "Can not find output file : "
const MSG_ERROR_4 = "Can not find tested file : "
const MSG_ERROR_END = "."

// Tester-File Address.
const test_address_1 = "http://ping.eu/img/icons/22_bandwidth.gif"
const test_host_1 = "ping.eu"

//------------------------------------------------------------------------------

// Text Separators.
var NLB []byte = []byte(NL)
var HTTP_HeadBodySeparator []byte = []byte(NL + NL)
var HTTP_HeadBodySeparatorLen int = len(HTTP_HeadBodySeparator)

// Command Line Arguments' Pointers.
var cla_file_input_ptr *string
var cla_file_output_ptr *string
var cla_file_tester_ptr *string

// Files' Parameters.
var FileInputPath string
var FileOutputPath string
var FileTesterPath string

var FileInputExists bool
var FileOutputExists bool
var FileTesterExists bool

var file_input *os.File
var FileTesterContent []byte

// Connection Parameters.
var ConnectionType = "tcp"
var ConnectionTimeout time.Duration

// Timing Parameters.
var delay_1 time.Duration
var tick time.Duration

// Work Data.
var input_data []string
var addresses_count int64
var address_num int64
var SaveChan chan string
var Tasks uint
var SaveIsInProgress bool
var GoodConnectServers []string
var GoodGetServers []string

//------------------------------------------------------------------------------

// Command Line Arguments Initialization.

func cla_init() {

	cla_file_input_ptr = flag.String(CLA_PARAM_FILE_INPUT,
		CLA_FILE_INPUT_DEFAULT, "Input File")
	cla_file_output_ptr = flag.String(CLA_PARAM_FILE_OUTPUT,
		CLA_FILE_OUTPUT_DEFAULT, "Output File")
	cla_file_tester_ptr = flag.String(CLA_PARAM_FILE_TESTER,
		CLA_FILE_TESTER_DEFAULT, "Tested File")

	flag.Parse()

	FileInputPath = *cla_file_input_ptr
	FileOutputPath = *cla_file_output_ptr
	FileTesterPath = *cla_file_tester_ptr
}

//------------------------------------------------------------------------------

// Checks Existance of File.

func file_exists(FilePath string) bool {

	// Returns 'false' if File does not exist or is not accessible.
	// Returns 'true' if File exists and is accessible.

	var err error

	fmt.Println("Checking existance of file :", FilePath) //

	_, err = os.Stat(FilePath)

	if os.IsNotExist(err) {

		return false
	}

	if err == nil {

		return true

	} else {

		return false
	}
}

//------------------------------------------------------------------------------

// Checks Existance of Input File.

func file_input_check() {

	FileInputExists = file_exists(FileInputPath)

	if !FileInputExists {

		fmt.Println(MSG_ERROR_1 + FileInputPath + MSG_ERROR_END)
		os.Exit(EXIT_CODE_FILE_INPUT_NOT_FOUND)
	}
}

//------------------------------------------------------------------------------

// Checks Existance of Output File.

func file_output_check() {

	FileOutputExists = file_exists(FileOutputPath)

	if !FileOutputExists {

		fmt.Println(MSG_ERROR_3 + FileOutputPath + MSG_ERROR_END)
		os.Exit(EXIT_CODE_FILE_OUTPUT_NOT_FOUND)
	}
}

//------------------------------------------------------------------------------

// Checks Existance of Tested File.

func file_tester_check() {

	FileTesterExists = file_exists(FileTesterPath)

	if !FileTesterExists {

		fmt.Println(MSG_ERROR_4 + FileTesterPath + MSG_ERROR_END)
		os.Exit(EXIT_CODE_FILE_TESTER_NOT_FOUND)
	}
}

//------------------------------------------------------------------------------

// Loads the Tested File.

func file_tester_load() {

	var err error

	FileTesterContent, err = ioutil.ReadFile(FileTesterPath)
	if err != nil {

		log.Fatal(err)
	}
}

//------------------------------------------------------------------------------

// Reads input File Line by Line and stores them into Array.

func file_input_read() {

	var err error
	var scanner *bufio.Scanner

	file_input, err = os.Open(FileInputPath)
	if err != nil {

		log.Fatal(err)
	}
	defer file_input.Close()

	scanner = bufio.NewScanner(file_input)
	for scanner.Scan() {

		input_data = append(input_data, scanner.Text())
		addresses_count++
	}

	err = scanner.Err()
	if err != nil {

		// Not EOF Error.
		log.Fatal(err)
	}
}

//------------------------------------------------------------------------------

// Processes an Array of Addresses.

func process_data() {

	var address string
	var host string
	var host_len int
	var msg_error string
	var n int
	var port string
	var port_len int

	for n, address = range input_data {

		Tasks++

		arr := strings.Split(address, PORT_DELIMITER)

		if len(arr) < 2 {

			msg_error = MSG_ERROR_2 + strconv.Itoa(n+1) + MSG_ERROR_END
			fmt.Println(msg_error)
			os.Exit(EXIT_CODE_FILE_INPUT_SYNTAX_BAD)
		}

		host = arr[0]
		port = arr[1]

		// Check Length of Data.
		host_len = len(host)
		port_len = len(port)
		if (host_len == 0) || (port_len == 0) {

			msg_error = MSG_ERROR_2 + strconv.Itoa(n+1) + MSG_ERROR_END
			fmt.Println(msg_error)
			os.Exit(EXIT_CODE_FILE_INPUT_SYNTAX_BAD)
		}

		address_num++
		go query_server(address_num, host, port)

		time.Sleep(delay_1)
	}

	// Wait for Results
	for (Tasks > 0) || (SaveIsInProgress) {

		time.Sleep(tick)
	}
}

//------------------------------------------------------------------------------

// Queries the Server.

func query_server(num int64, host, port string) {

	var connect_str string // CONNECT Request
	var get_str string     // GET Request
	var request []byte

	fmt.Printf("Processing Address [%d of %d]\r\n", num, addresses_count)

	// 'CONNECT' Method.
	connect_str = ConnectMethod + RequestSpace +
		host + PORT_DELIMITER + port + RequestSpace +
		RequestHTTPVersionString +
		NL + NL
	request = []byte(connect_str)
	fmt.Println("CONNECT", host, ":", port)
	//fmt.Printf("\r\n%v", connect_str)
	host_query_1(num, host, port, ConnectionType, request, ConnectionTimeout)

	// 'GET' Method.
	get_str = RequestMethod + RequestSpace +
		test_address_1 + RequestSpace +
		RequestHTTPVersionString +
		NL +
		// Host:
		RequestHostVN + RequestParameterValueDelimiter + RequestSpace +
		test_host_1 +
		NL + NL
	request = []byte(get_str)
	fmt.Println("GET", host, ":", port)
	//fmt.Printf("\r\n%v", get_str) ///
	host_query_2(num, host, port, ConnectionType, request, ConnectionTimeout)

	Tasks--
}

//------------------------------------------------------------------------------

// Tries to use HTTP 'CONNECT' Method to connect to Server.

func host_query_1(
	num int64,
	host, port, connType string,
	req []byte,
	timeout time.Duration) {

	var ConnectionAddress string
	var Connection net.Conn
	var err error
	var RequestLength, ResponseLength int
	var Response string
	var ResponseBuffer []byte
	var isOK bool
	var saveText string
	var now time.Time

	// Connect
	ConnectionAddress = host + PORT_DELIMITER + port
	Connection, err = net.DialTimeout(connType, ConnectionAddress, timeout)
	if err != nil {

		log.Println(num, "Failed to connect with ", ConnectionAddress) //
		log.Println(err)
		return
	}

	// Send Request
	RequestLength, err = Connection.Write(req)
	if err != nil {

		Connection.Close()
		log.Println(num, "Failed to send request to ", ConnectionAddress) //
		log.Println(err)
		return
	}
	if RequestLength == 0 {

		Connection.Close()
		log.Println(num, "Empty request to ", ConnectionAddress) //
		return
	}

	// Get Response
	ResponseBuffer = make([]byte, BUFFER_SIZE)
	ResponseLength, err = Connection.Read(ResponseBuffer)
	if err != nil {

		Connection.Close()
		log.Println(num, "Failed to read response from ", ConnectionAddress) //
		log.Println(err)
		return
	}
	if ResponseLength == 0 {

		Connection.Close()
		log.Println(num, "Empty response to ", ConnectionAddress) //
		return
	}

	Response = string(ResponseBuffer)
	//fmt.Printf("Response:\r\n%v\r\n", Response)

	// Check Response Status
	isOK = false
	isOK = strings.HasPrefix(Response, RequestStatusSuccessPattern)

	if isOK {

		GoodConnectServers = append(GoodConnectServers, ConnectionAddress)

		now = time.Now()
		saveText = now.Format(TIME_FORMAT) +
			" GOOD CONNECT " + ConnectionAddress
		//SaveChan <- saveText // Do not save CONNECT requests.
		fmt.Println(num, saveText)

	} else {

		fmt.Println(num, "BAD CONNECT", ConnectionAddress)
	}

	// Close Connection
	Connection.Close()
}

//------------------------------------------------------------------------------

// Tries to use HTTP 'GET' Method to connect to Server.

func host_query_2(
	num int64,
	host, port, connType string,
	req []byte,
	timeout time.Duration) {

	var ConnectionAddress string
	var Connection net.Conn
	var err error
	var RequestLength, ResponseLength int
	var Response string
	var ResponseBuffer []byte
	var isOK bool
	var saveText string

	var ResponseHead []byte
	var ResponseBody []byte
	var ResponseContent []byte
	var ResponseIndexSep int

	var ResponseContentLength int64
	var ResponseContentLength_str string
	var now time.Time

	// Connect
	ConnectionAddress = host + PORT_DELIMITER + port
	Connection, err = net.DialTimeout(connType, ConnectionAddress, timeout)
	if err != nil {

		log.Println(num, "Failed to connect with ", ConnectionAddress) //
		log.Println(err)
		return
	}

	// Send Request
	RequestLength, err = Connection.Write(req)
	if err != nil {

		Connection.Close()
		log.Println(num, "Failed to send request to ", ConnectionAddress) //
		log.Println(err)
		return
	}
	if RequestLength == 0 {

		Connection.Close()
		log.Println(num, "Empty request to ", ConnectionAddress) //
		return
	}

	// Get Response
	ResponseBuffer = make([]byte, BUFFER_SIZE)
	ResponseLength, err = Connection.Read(ResponseBuffer)
	if err != nil {

		Connection.Close()
		log.Println(num, "Failed to read response from ", ConnectionAddress) //
		log.Println(err)
		return
	}
	if ResponseLength == 0 {

		Connection.Close()
		log.Println(num, "Empty response to ", ConnectionAddress) //
		return
	}

	Response = string(ResponseBuffer)
	//fmt.Printf("Response:\r\n%v\r\n", Response)///

	// Check Response Status
	isOK = false
	isOK = strings.HasPrefix(Response, RequestStatusSuccessPattern)

	if !isOK {
		Connection.Close()
		log.Println(num, "Bad HTTP Status from", ConnectionAddress) //
		return
	}

	// Check Content
	ResponseIndexSep = bytes.Index(ResponseBuffer, HTTP_HeadBodySeparator)
	ResponseHead = ResponseBuffer[:ResponseIndexSep]
	ResponseBody = ResponseBuffer[ResponseIndexSep+HTTP_HeadBodySeparatorLen:]

	// Find Content-Length.
	ResponseContentLength_str = getHTTPRequestHeaderParameter(
		"Content-Length", ResponseHead)
	ResponseContentLength, err = strconv.ParseInt(ResponseContentLength_str, 10, 0)
	if err != nil {
		Connection.Close()
		log.Println(num, "Failed to get Content-Length from", ConnectionAddress) //
		log.Println(err)
		return
	}

	// Check Size of Content-Length.
	if ResponseContentLength > int64(len(ResponseBody)) {

		Connection.Close()
		log.Println(num, "Bad Content-Length from", ConnectionAddress) //
		return
	}

	ResponseContent = ResponseBody[:ResponseContentLength]
	//log.Printf("CONTENT:\r\n%s", ResponseContent) ///
	isOK = bytes.Equal(ResponseContent, FileTesterContent)

	if isOK {

		GoodGetServers = append(GoodGetServers, ConnectionAddress)

		now = time.Now()
		saveText = now.Format(TIME_FORMAT) +
			" GOOD GET " + ConnectionAddress
		SaveChan <- saveText
		fmt.Println(num, saveText)

	} else {

		fmt.Println(num, "BAD GET", ConnectionAddress)
	}

	// Close Connection
	Connection.Close()
}

//------------------------------------------------------------------------------

// Searches Request Header for the specified Parameter.

func getHTTPRequestHeaderParameter(parameter string, header []byte) string {

	var p1 int // Position where Parameter's Name starts.
	var p2 int // Position where Parameter's Value starts.
	var tmp1 []byte
	var value []byte
	var value_len int

	p1 = bytes.Index(header, []byte(parameter))
	p2 = p1 + len(parameter) + 2 // Name + ':' + ' '.
	tmp1 = header[p2:]
	value_len = bytes.Index(tmp1, NLB)

	if value_len == -1 {

		// No CRLF after the Value => it is the last Line in Header.
		value = header[p2:]

		log.Println("ACHTUNG!")       ///
		log.Println("SHORT HEADER!")  ///
		log.Println("value =", value) ///
		/*
			log.Println("ACHTUNG!")               ///
			log.Println("p2 =", p2)               ///
			log.Println("value_len =", value_len) ///
			log.Printf("header = %s\r\n", header) ///
		*/

	} else {

		value = header[p2 : p2+value_len]
	}

	return string(value)
}

//------------------------------------------------------------------------------

// Manager which saves successfull Attempts to an external File.

func file_output_manager() {

	var file *os.File
	var err error
	var data string

	file, err = os.OpenFile(FileOutputPath, os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for {

		data = <-SaveChan

		SaveIsInProgress = true

		_, err = file.WriteString(data + NL)
		if err != nil {
			panic(err)
		}

		SaveIsInProgress = false
	}
}

//------------------------------------------------------------------------------

// Program's Entry Point.

func main() {

	// Various Initializations.
	ConnectionTimeout = time.Second * CONNECTION_TIMEOUT_SEC
	tick = time.Second * 1
	delay_1 = time.Microsecond * 10
	SaveChan = make(chan string)
	Tasks = 0
	address_num = 0

	cla_init()
	file_input_check()
	file_output_check()
	file_tester_check()
	file_tester_load()

	// Save-Manager.
	go file_output_manager()

	// Work.
	file_input_read()
	process_data()

	// Show Results.
	fmt.Println("\r\nGood CONNECT Servers:")
	for _, v := range GoodConnectServers {

		fmt.Println(v)
	}
	fmt.Println("\r\nGood GET Servers:")
	for _, v := range GoodGetServers {

		fmt.Println(v)
	}
}

//------------------------------------------------------------------------------
