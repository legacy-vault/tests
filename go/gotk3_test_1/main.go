// main.go

// A Test of GTK Usage in Go Language (Golang).

//=============================================================================|

package main

//=============================================================================|

import "hash/crc64"
import "errors"
import "encoding/hex"
import "io/ioutil"
import "log"
import "crypto/md5"
import "net"
import "strings"
import "strconv"
import "time"

//=============================================================================|

// Signals.
const APP_MANAGER_SIGNAL_STOP = true

// Manager Configuration.
const APP_MANAGER_WAIT_SIGNAL_INTERVAL = 3

// Messages.

// 1. Normal Messages.
const MSG_1 = "Program has started."
const MSG_2 = "Program is shutting down..."
const MSG_3 = "Signal to UI Manager."
const MSG_4 = "Signal to Application Manager."
const MSG_5 = "Signal to Server Manager."
const MSG_6 = "Server Manager has stopped."
const MSG_7 = "Application Manager has stopped."
const MSG_8 = "UI Manager has stopped."

// 2. Error Messages.
const MSG_ERR_1 = "Assertion Error."
const MSG_ERR_2 = "Scroll Error."
const MSG_ERR_3 = "Abnormal Program Shutdown! " +
	"User Interface Loop has not been stopped!"
const MSG_ERR_4 = "Client Message Length Error."
const MSG_ERR_5 = "Network Message Creation Error."
const MSG_ERR_6 = "Network Message Send Error."
const MSG_ERR_7 = "Check Sum Error."

// Special Characters.
const NL = "\r\n"

//=============================================================================|

var clientMessage string
var consoleText string
var verbose bool

// Channels.
var appManagerChan chan bool

//=============================================================================|

// Application (main Program) Manager.
func app_manager() {

	var loop bool
	var sleepPeriod time.Duration
	var signal bool

	sleepPeriod = time.Second * APP_MANAGER_WAIT_SIGNAL_INTERVAL

	// Wait for Signal to stop.
	loop = true
	for loop {

		select {

		case signal = <-appManagerChan:

			// Signal is received.

			// Verbose Report.
			if verbose {
				log.Println(MSG_4)
			}

			if signal == APP_MANAGER_SIGNAL_STOP {
				loop = false
			}

		default:

			// No Signal is received.
			time.Sleep(sleepPeriod)
		}
	}

	// Abnormal Shutdown?
	if uiHasStopped == false {
		log.Println(MSG_ERR_3)
	}

	log.Println(MSG_7)
}

//=============================================================================|

// Stops the Program by starting the Chain of Stop-Signals.
func app_quit() {

	// Report to Console.
	ui_console_msg_add(MSG_2)

	// Send a Signal.
	uiManagerChan <- UI_MANAGER_SIGNAL_STOP
}

//=============================================================================|

// Program's Entry point.
func main() {

	var err error

	// Self-Test.
	err = self_test()
	if err != nil {
		log.Println(err)
		return
	}

	// Initialize internal Data.

	// 1. Various Variables.
	consoleText = ""
	serverAddress = &net.TCPAddr{}
	serverIsWorking = false
	verbose = true

	// Initialize User Interface.
	err = ui_init()
	if err != nil {
		log.Println(err)
		return
	}

	// Create Channels.
	appManagerChan = make(chan bool, 1)
	serverManagerChan = make(chan byte, 1)
	uiManagerChan = make(chan bool, 1)

	// Start Managers.
	go ui_manager()
	app_manager()
}

//=============================================================================|

// Self-Test.
func self_test() error {

	var crcSum uint64
	var crcSumStr string
	var crcTable *crc64.Table
	var err error
	var fileData []byte
	var md5Sum [md5.Size]byte
	var md5SumBA []byte
	var md5SumStr string

	// Preparations.
	crcTable = crc64.MakeTable(crc64.ECMA)

	// Read UI File.
	fileData, err = ioutil.ReadFile(UI_FILE_PATH)
	if err != nil {
		return err
	}

	// Check Integrity.

	// 1. UI File's CRC-64.
	crcSum = crc64.Checksum(fileData, crcTable)
	crcSumStr = strconv.FormatUint(crcSum, 16)
	crcSumStr = strings.ToUpper(crcSumStr)

	if crcSumStr != UI_FILE_CHECKSUM_CRC64_HEX {
		err = errors.New(MSG_ERR_7)
		return err
	}

	// 1. UI File's MD5.
	md5Sum = md5.Sum(fileData)
	md5SumBA = md5Sum[:]
	md5SumStr = hex.EncodeToString(md5SumBA)
	md5SumStr = strings.ToUpper(md5SumStr)

	if md5SumStr != UI_FILE_CHECKSUM_MD5_HEX {
		err = errors.New(MSG_ERR_7)
		return err
	}

	return nil
}

//=============================================================================|
