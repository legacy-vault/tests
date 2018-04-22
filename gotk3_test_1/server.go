// server.go

// Server Functions.

//=============================================================================|

package main

//=============================================================================|

import "hash/crc64"
import "log"
import "crypto/md5"
import "net"
import "strconv"
import "sync"
import "github.com/gotk3/gotk3/gtk"
import "time"

//=============================================================================|

type ServerConfiguration struct {
	Host string
	Port string
}

//=============================================================================|

// Signals.
const SERVER_MANAGER_SIGNAL_STOP byte = 1

// Manager Configuration.
const SERVER_LISTEN_TIMEOUT = 10

const SERVER_REPLY_ERROR byte = 0
const SERVER_REPLY_OK byte = 255

//=============================================================================|

var serverAddress *net.TCPAddr
var serverConfig ServerConfiguration
var serverConnection *net.TCPListener
var serverConnectionsActive sync.WaitGroup
var serverManager sync.WaitGroup

// Flags.
var serverIsWorking bool

// Channels.
var serverManagerChan chan byte

//=============================================================================|

// Reads Server Port from UI into Variable.
func config_server_port_read() error {

	// Since we read this Value rarely, we do not store the Buffer globally.
	var buffer *gtk.EntryBuffer
	var err error

	// Get Buffer.
	buffer, err = inputServerPort.GetBuffer()
	if err != nil {
		return err
	}

	// Get Port.
	serverConfig.Port, err = buffer.GetText()
	if err != nil {
		return err
	}

	return nil
}

//=============================================================================|

// Server Manager.
func server_manager() {

	var deadline time.Time
	var err error
	var err2 net.Error
	var ok bool
	var loop bool
	var newConn net.Conn
	var signal byte

	// Wait for Signal to stop.
	loop = true
	for loop {

		select {

		case signal = <-serverManagerChan:

			// Signal is received.

			// Verbose Report.
			if verbose {
				log.Println(MSG_5)
			}

			if signal == SERVER_MANAGER_SIGNAL_STOP {

				// Stop listening to new incoming Connections.
				loop = false
			}

		default:

			// No Signal is received. Serve Clients.

			// Get new incoming Connection.
			deadline = time.Now().Add(time.Second * SERVER_LISTEN_TIMEOUT)
			serverConnection.SetDeadline(deadline)
			newConn, err = serverConnection.Accept()
			if err != nil {

				// Error occurred.

				err2, ok = err.(net.Error)
				if ok {
					if err2.Timeout() {
						// Error is due to Timeout.
					} else {
						log.Println(err)
					}

				} else {
					log.Println(err)
				}

			} else {

				// No Error.

				// Handle new Connection.
				serverConnectionsActive.Add(1)
				go server_request_handle(newConn)
			}
		}
	}

	ui_console_msg_add(MSG_6)

	serverManager.Done()
}

//=============================================================================|

// Checks Data of an incoming Request from Client.
func server_request_data_check(data []byte) (bool, []byte) {

	var crcData []byte
	var crcSum uint64
	var crcSumExpected uint64
	var crcTable *crc64.Table
	var i int
	var j int
	var md5Data []byte
	var md5SumExpected [md5.Size]byte
	var netMsgPostfix uint16
	var netMsgPrefix uint16
	var netMsgSize int
	var netMsgText []byte
	var netMsgTextSize uint16
	var text []byte

	// 1. Network Message Size.
	netMsgSize = len(data)
	if (netMsgSize < MESSAGE_SIZE_MIN) || (netMsgSize > MESSAGE_SIZE_MAX) {
		return false, text
	}

	// 2. Network Message Prefix.
	netMsgPrefix = (uint16(data[0]) << 8) + uint16(data[1])
	if netMsgPrefix != MESSAGE_PREFIX {
		return false, text
	}

	// 3. Network Message Text Size.
	netMsgTextSize = (uint16(data[2]) << 8) + uint16(data[3])
	if (uint16(netMsgSize) - MESSAGE_AUX_FIELDS_SIZE) != netMsgTextSize {
		return false, text
	}

	// Get Text from Network Message.
	netMsgText = data[4 : 4+netMsgTextSize]

	// 4. CRC-64.
	crcTable = crc64.MakeTable(crc64.ECMA)
	crcSumExpected = crc64.Checksum(netMsgText, crcTable)

	j = 4 + int(netMsgTextSize)
	crcData = data[j : j+crc64.Size]

	crcSum = 0
	crcSum = crcSum + uint64(crcData[0])<<56
	crcSum = crcSum + uint64(crcData[1])<<48
	crcSum = crcSum + uint64(crcData[2])<<40
	crcSum = crcSum + uint64(crcData[3])<<32
	crcSum = crcSum + uint64(crcData[4])<<24
	crcSum = crcSum + uint64(crcData[5])<<16
	crcSum = crcSum + uint64(crcData[6])<<8
	crcSum = crcSum + uint64(crcData[7])<<0

	if crcSum != crcSumExpected {
		return false, text
	}

	// 5. MD5.
	j = j + 8
	md5Data = data[j : j+md5.Size]
	md5SumExpected = md5.Sum(netMsgText)
	for i, _ = range md5SumExpected {
		if md5Data[i] != md5SumExpected[i] {
			return false, text
		}
	}

	// 6. Network Message Postfix.
	j = j + 16
	netMsgPostfix = (uint16(data[j]) << 8) + uint16(data[j+1])
	if netMsgPostfix != MESSAGE_POSTFIX {
		return false, text
	}

	// Prepare Text.
	text = netMsgText

	return true, text
}

//=============================================================================|

// Handles an incoming Request from Client.
func server_request_handle(conn net.Conn) {

	var buffer []byte
	var clientAddress net.Addr
	var data []byte
	var err error
	var msg string
	var ok bool
	var requestLen int
	var text []byte

	// Prepare Buffer.
	buffer = make([]byte, MESSAGE_SIZE_MAX)

	// Request -> Buffer.
	requestLen, err = conn.Read(buffer)
	if err != nil {
		log.Println(err)
		conn.Close()
		serverConnectionsActive.Done()
		return
	}
	data = buffer[0:requestLen]

	// Check Request Data.
	ok, text = server_request_data_check(data)
	if ok == false {

		// Send bad Response.
		_, err = conn.Write([]byte{SERVER_REPLY_ERROR})
		if err != nil {
			log.Println(err)
			conn.Close()
			serverConnectionsActive.Done()
			return
		}

		// Dis-Connect.
		err = conn.Close()
		if err != nil {
			log.Println(err)
			serverConnectionsActive.Done()
			return
		}
	}

	// Send good Response.
	_, err = conn.Write([]byte{SERVER_REPLY_OK})
	if err != nil {
		log.Println(err)
		conn.Close()
		serverConnectionsActive.Done()
		return
	}

	// Dis-Connect.
	err = conn.Close()
	if err != nil {
		log.Println(err)
		serverConnectionsActive.Done()
		return
	}

	// No Errors.

	// Show Message from Client.
	clientAddress = conn.RemoteAddr()
	msg = "Message from [" + clientAddress.String() + "]:" + NL + string(text)
	ui_console_msg_add(msg)

	serverConnectionsActive.Done()
}

//=============================================================================|

// Starts the Server.
func server_start() error {

	var err error
	var msg string
	var port int
	var tmpUint64 uint64

	// Convert Port (string -> int).
	tmpUint64, err = strconv.ParseUint(serverConfig.Port, 10, 32)
	if err != nil {
		return err
	}
	port = (int)(tmpUint64)

	// Configure Server Address.
	serverAddress.IP = net.IPv4zero
	serverAddress.Port = port
	serverAddress.Zone = ""

	// Start the Server.
	serverConnection, err = net.ListenTCP(NETWORK_PROTOCOL, serverAddress)
	if err != nil {
		return err
	}

	// Show LEDs.
	indicatorServerOn.SetFromPixbuf(pixbufLEDOnActive)
	indicatorServerOff.SetFromPixbuf(pixbufLEDOffPassive)

	// Report to Console.
	msg = "Listening at " +
		serverAddress.IP.String() + ":" +
		strconv.FormatInt(int64(serverAddress.Port), 10) + "."

	ui_console_msg_add(msg)

	// Start Server Manager.
	serverManager.Add(1)
	go server_manager()

	// Update Server Status.
	serverIsWorking = true

	return nil
}

//=============================================================================|

// Stops the Server.
func server_stop() error {

	var err error
	var msg string

	// Change Indicators.
	indicatorServerOn.SetFromPixbuf(pixbufLEDOnTransit)
	indicatorServerOff.SetFromPixbuf(pixbufLEDOffTransit)

	// Send Signal to Manager.
	serverManagerChan <- SERVER_MANAGER_SIGNAL_STOP

	// Wait for Server Manager to stop listening Network.
	msg = "Waiting for Server Manager to stop..."
	ui_console_msg_add(msg)
	serverManager.Wait()

	// Wait for all active Connections to become finished.
	msg = "Waiting for active Connections to finish..."
	ui_console_msg_add(msg)
	serverConnectionsActive.Wait()

	// Stop the Server.
	msg = "Stopping the Server..."
	ui_console_msg_add(msg)
	err = serverConnection.Close()
	if err != nil {
		return err
	}

	// Show 'Done' LEDs.
	indicatorServerOn.SetFromPixbuf(pixbufLEDOnPassive)
	indicatorServerOff.SetFromPixbuf(pixbufLEDOffActive)

	// Report to Console.
	msg = "Stopped listening at " +
		serverAddress.IP.String() + ":" +
		strconv.FormatInt(int64(serverAddress.Port), 10) + "."

	ui_console_msg_add(msg)

	// Update Server Status.
	serverIsWorking = false

	return nil
}

//=============================================================================|
