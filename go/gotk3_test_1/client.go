// client.go

// Client Functions.

//=============================================================================|

package main

//=============================================================================|

import "bytes"
import "hash/crc64"
import "errors"
import "io/ioutil"
import "crypto/md5"
import "net"
import "github.com/gotk3/gotk3/gtk"

//=============================================================================|

type ClientConfiguration struct {
	Host string
	Port string
}

//=============================================================================|

var clientConfig ClientConfiguration

//=============================================================================|

// Prepares a Byte Array of Network Message.
func client_message_prepare() ([]byte, error) {

	var crcSum uint64
	var crcTable *crc64.Table
	var err error
	var i int
	var j int
	var md5Sum [md5.Size]byte
	var msgBA []byte
	var msgSizeBytes int
	var netMsg []byte
	var netMsgPostfix uint16
	var netMsgPrefix uint16
	var netMsgSize int

	// Message -> Byte Array.
	msgBA = []byte(clientMessage)

	// Check Message.
	msgSizeBytes = len(msgBA)
	if msgSizeBytes > MESSAGE_TEXT_SIZE_MAX {
		err = errors.New(MSG_ERR_4)
		return nil, err
	}

	// Create CRC-64 Check Sum.
	crcTable = crc64.MakeTable(crc64.ECMA)
	crcSum = crc64.Checksum(msgBA, crcTable)

	// Create MD5 Check Sum.
	md5Sum = md5.Sum(msgBA)

	// Network Message Prefix.
	netMsgPrefix = MESSAGE_PREFIX

	// Network Message Postfix.
	netMsgPostfix = MESSAGE_POSTFIX

	// Form the Network Message.
	netMsgSize =
		2 + // Prefix.
			2 + // Message (Text) Size.
			msgSizeBytes + // Message (Text).
			crc64.Size + // CRC-64.
			md5.Size + // MD5.
			2 // Postfix.
	netMsg = make([]byte, netMsgSize)

	// 1. Fill Prefix.
	netMsg[0] = byte(netMsgPrefix >> 8)
	netMsg[1] = byte(netMsgPrefix)

	// 2. Fill Message (Text) Size.
	netMsg[2] = byte(msgSizeBytes >> 8)
	netMsg[3] = byte(msgSizeBytes)

	// 3. Fill Message (Text).
	for i, _ = range msgBA {
		j = 4 + i
		netMsg[j] = msgBA[i]
	}
	j++

	// 4. Fill CRC-64 Check Sum (8 Bytes)
	netMsg[j] = byte(crcSum >> 56)
	j++
	netMsg[j] = byte(crcSum >> 48)
	j++
	netMsg[j] = byte(crcSum >> 40)
	j++
	netMsg[j] = byte(crcSum >> 32)
	j++
	netMsg[j] = byte(crcSum >> 24)
	j++
	netMsg[j] = byte(crcSum >> 16)
	j++
	netMsg[j] = byte(crcSum >> 8)
	j++
	netMsg[j] = byte(crcSum >> 0)
	j++

	// 5. Fill MD5 Check Sum (16 Bytes)
	for i, _ = range md5Sum {
		netMsg[j] = md5Sum[i]
		j++
	}

	// 6. Fill Postfix.
	netMsg[j] = byte(netMsgPostfix >> 8)
	j++
	netMsg[j] = byte(netMsgPostfix)
	j++

	// Check Number of Bytes Written.
	if j != netMsgSize {
		err = errors.New(MSG_ERR_5)
		return nil, err
	}

	return netMsg, nil
}

//=============================================================================|

// Sends the Message from Client.
func client_message_send() error {

	var bytesCount int
	var clientAddress string
	var connection net.Conn
	var err error
	var msg string
	var networkMessage []byte
	var responseBA []byte

	// Read Client Message.
	err = config_client_message_read()
	if err != nil {
		return err
	}

	// Read Client Host.
	err = config_client_host_read()
	if err != nil {
		return err
	}

	// Read Client Port.
	err = config_client_port_read()
	if err != nil {
		return err
	}

	// Prepare Network Message.
	networkMessage, err = client_message_prepare()
	if err != nil {
		return err
	}

	// Configure Client.
	clientAddress = clientConfig.Host + ":" + clientConfig.Port

	// Connect the Client to Remote Server.
	connection, err = net.Dial(NETWORK_PROTOCOL, clientAddress)
	if err != nil {
		return err
	}

	// Report to Console.
	msg = "Connected to " + clientAddress + "."
	ui_console_msg_add(msg)

	// Send the Message.
	bytesCount, err = connection.Write(networkMessage)
	if err != nil {
		return err
	}
	if bytesCount != len(networkMessage) {

		err = errors.New(MSG_ERR_6)
		connection.Close()
		return err
	}

	// Get Response.
	responseBA, err = ioutil.ReadAll(connection)

	if bytes.Equal(responseBA, []byte{SERVER_REPLY_OK}) {

		// Report to Console.
		msg = "Message has been successfully sent."
		ui_console_msg_add(msg)

	} else {

		// Report to Console.
		msg = "Message delivery has failed."
		ui_console_msg_add(msg)
	}

	// Dis-Connect.
	err = connection.Close()
	if err != nil {
		return err
	}

	// Report to Console.
	msg = "Disconnected from " + clientAddress + "."
	ui_console_msg_add(msg)

	return nil
}

//=============================================================================|

// Reads Client Host from UI into Variable.
func config_client_host_read() error {

	var err error

	// Get Text from Buffer.
	clientConfig.Host, err = inputClientHostBuffer.GetText()
	if err != nil {
		return err
	}

	return nil
}

//=============================================================================|

// Reads Client Message from UI into Variable.
func config_client_message_read() error {

	var err error
	var ihc bool
	var posEnd *gtk.TextIter
	var posStart *gtk.TextIter

	// Prepare Arguments.
	ihc = true
	posStart = inputClientMessageBuffer.GetStartIter()
	posEnd = inputClientMessageBuffer.GetEndIter()

	// Get Text from Buffer.
	clientMessage, err = inputClientMessageBuffer.GetText(
		posStart, posEnd, ihc)

	if err != nil {
		return err
	}

	return nil
}

//=============================================================================|

// Reads Client Port from UI into Variable.
func config_client_port_read() error {

	var err error

	// Get Text from Buffer.
	clientConfig.Port, err = inputClientPortBuffer.GetText()
	if err != nil {
		return err
	}

	return nil
}

//=============================================================================|
