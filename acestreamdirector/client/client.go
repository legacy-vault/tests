// acestreamdirector/client.go

package client

//------------------------------------------------------------------------------

import "acestreamdirector/config"
import "acestreamdirector/message"
import "bufio"
import "errors"
import "encoding/hex"
import "log"
import "net"
import "crypto/sha1"
import "strings"
import "strconv"

//------------------------------------------------------------------------------

// 	Ace Stream Media API :
//		http://wiki.acestream.org/wiki/index.php/Engine_API

//------------------------------------------------------------------------------

var ConnBufferIn *bufio.Reader

//------------------------------------------------------------------------------

// Initializes Buffer for incoming Messages.
func BufferInInit(conn *net.TCPConn) {

	ConnBufferIn = bufio.NewReader(conn)
}

//------------------------------------------------------------------------------

// Connects to Server.
func Connect(connAddress *net.TCPAddr, verbose bool) (*net.TCPConn, error) {

	var conn *net.TCPConn
	var err error

	// Connect.
	conn, err = net.DialTCP(config.CONNECTION_TYPE, nil, connAddress)
	if err != nil {
		return nil, err
	}

	// Create Buffer.
	BufferInInit(conn)

	// Say Hello to Server.
	err = HelloRequestSend(conn, verbose)
	if err != nil {
		return nil, err
	}

	// Read Server's Reply.
	err = HelloResponseGet(ConnBufferIn, verbose)
	if err != nil {
		return nil, err
	}

	// Say Ready.
	err = ReadySend(conn, verbose)
	if err != nil {
		return nil, err
	}

	/////////////////////////////////////////////////////
	var msg *string                                  //
	msg, err = MessageReceive(ConnBufferIn, verbose) //
	if err != nil {                                  //
		return nil, err //
	} //
	msg = msg //
	////////////////////////////////////////////////////

	return conn, nil
}

//------------------------------------------------------------------------------

// Dis-Connects from Server.
func DisConnect(conn *net.TCPConn, verbose bool) error {

	var err error

	err = conn.Close()
	if err != nil {
		return err
	}

	return nil
}

//------------------------------------------------------------------------------

// Sends Hello-Message to Server.
func HelloRequestSend(conn *net.TCPConn, verbose bool) error {

	var err error
	var requestStr string

	requestStr = message.TYPE_1 + " " +
		message.PARAM_VERSION + "=" + config.API_VERSION_MAJOR + message.END

	err = MessageSend(conn, &requestStr, verbose)

	if err != nil {
		return err
	}

	return nil
}

//------------------------------------------------------------------------------

// Get Server's Response to Hello-Message.
func HelloResponseGet(buffer *bufio.Reader, verbose bool) error {

	var err error
	var inMsg *string // Incoming Message.
	var key string
	var msg *message.Message
	var param message.Parameter
	var portHTTPStr string
	var portHTTP int64
	var versionFull string
	var versionMajor string

	inMsg, err = MessageReceive(buffer, verbose)
	if err != nil {
		return err
	}

	msg, err = message.Parse(inMsg)
	if err != nil {
		return err
	}

	if msg.Type != message.TYPE_2 {
		err = errors.New("Wrong message type received.")
		return err
	}

	// Get some Parameters.
	for _, param = range msg.Parameters {

		if param.Key == message.PARAM_KEY {
			key = param.Value
			continue
		}

		if param.Key == message.PARAM_HTTP_PORT {
			portHTTPStr = param.Value
			continue
		}

		if param.Key == message.PARAM_VERSION {
			versionFull = param.Value
			continue
		}
	}

	// Get major Version & check it.
	versionMajor = message.VersionMajor(versionFull)
	if versionMajor != config.API_VERSION_MAJOR {
		err = errors.New("Unsupported version detected.")
		return err
	}

	// Save Configuration Parameters received from Server.
	portHTTP, err = strconv.ParseInt(portHTTPStr, 10, config.INT_SIZE)
	if err != nil {
		return err
	}
	config.ServerPortHTTP = int(portHTTP)
	config.ServerKey = key

	return nil
}

//------------------------------------------------------------------------------

// Receives a Message from Server.
func MessageReceive(buffer *bufio.Reader, verbose bool) (*string, error) {

	var err error
	var msgLen int
	var msgStr string
	var plb byte

	// API says that each Message must end with CR+LF Symbols Pair.

	// Get Message.
	msgStr, err = buffer.ReadString(message.LAST_BYTE)
	if err != nil {
		return nil, err
	}

	// Check Length.
	msgLen = len(msgStr)
	if msgLen < message.LENGTH_MIN {
		err = errors.New("The length of the message is too small.")
		return nil, err
	}

	// Check pre-last Byte.
	plb = msgStr[msgLen-2]
	if plb != message.PRELAST_BYTE {
		err = errors.New("Bad character at the end of the message.")
		return nil, err
	}

	if verbose {
		log.Println(msgStr)
	}

	return &msgStr, nil
}

//------------------------------------------------------------------------------

// Sends a Message to Server.
func MessageSend(conn *net.TCPConn, messageStr *string, verbose bool) error {

	var bytesSent int
	var err error
	var messageBytes []byte

	messageBytes = []byte(*messageStr)

	if verbose {
		log.Println(*messageStr)
	}

	bytesSent, err = conn.Write(messageBytes)
	if err != nil {
		return err
	}

	err = message.CheckLength(messageBytes, bytesSent)
	if err != nil {
		return err
	}

	return nil
}

//------------------------------------------------------------------------------

// Sends Ready-Message to Server.
func ReadySend(conn *net.TCPConn, verbose bool) error {

	var err error
	var msg string
	var responseKey string

	responseKey, err = ResponseKey()
	if err != nil {
		return err
	}

	msg = message.TYPE_3 + " " +
		message.PARAM_KEY + "=" + responseKey + message.END

	err = MessageSend(conn, &msg, verbose)

	if err != nil {
		return err
	}

	return nil
}

//------------------------------------------------------------------------------

func ResponseKey() (string, error) {

	var err error
	var hexStr string
	var i int
	var parts []string
	var result string
	var sha1sum [sha1.Size]byte
	var sha1sumBA []byte
	var signature string
	var tmp string
	var x string

	tmp = config.ServerKey + config.ProductKey
	sha1sum = sha1.Sum([]byte(tmp))
	sha1sumBA = make([]byte, sha1.Size)
	for i = range sha1sum {
		sha1sumBA[i] = sha1sum[i]
	}
	hexStr = hex.EncodeToString(sha1sumBA)
	signature = hexStr

	parts = strings.Split(config.ProductKey, message.X_SEPARATOR)
	if len(parts) < 1 {
		err = errors.New("Bad product key.")
		return "", err
	}
	x = parts[0]

	result = x + message.X_SEPARATOR + signature

	return result, nil
}

//------------------------------------------------------------------------------
