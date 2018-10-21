// server.go

/*

	Simple Queue Server.

	Version: 0.1.
	Date of Creation: 2017-06-01.
	Author: McArcher.

	A very simple Queue Server.

	The Queue is a FIFO List.
	In this Example, the Queue is made with Go's Channel.
	It is very easy and supports only few Methods.
	Available methods are:
		- put Item,
		- get Item,
		- show actual Size (current Number of Items in Queue),
		- close Connection.

	To do more complicated Tasks, Constructs other than Channels will be needed.
	The stored Item is an Array of Bytes in this Example.
	This Program is just a Test to demonstrate Capabilities of Golang.
	Note that Sizes in Requests and Replies are written in binary Mode.

	While this test Server uses only one Queue and each Operation (except 'size')
	modifies the Queue, it is not recommended to enable multiple simultaneous
	Connections to this Server. That is why it processes only a single Connection
	at a Time. All Connections are put into a Task Queue.

*/

//------------------------------------------------------------------------------

package main

import (
	//"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

//------------------------------------------------------------------------------

/*

	Protocol.

	In this test API the Protocol is very simple.
	In the Text below, "BB" means two binary Bytes, "BBBB" is four binary Bytes.

	---------

	1.	Put (push, add) Element to the Queue.

		Client:
				'p' + ' ' + BB + ' ' + <Value>.
				<Value> is a UTF-8 string. "BB" is Size in Bytes (uint16).
		Server:
				'k' when OK,
				'f' if Queue is full,
				'r' on Error.

	---------

	2.	Get (read, pop) Element from the Queue.

		Client:
				'g'.
		Server:
				'e' if Queue is empty,
				'r' on Error,
				'i' + ' ' + BB + ' ' + <Value>.

	---------

	3.	Show the actual Size of the Queue.

		Client:
				's'.
		Server:
				'r' on Error,
				's' + ' ' + BBBB. BBBB is Size (uint64).

	---------

	4.	Close Connection (quit).

		Client:
				'q'.
		Server:
				'd'.

	---------
*/

//------------------------------------------------------------------------------

const PORT_DELIMITER = ":"
const CLIENT_IPA_ALLOWED = "127.0.0.1"
const QUEUE_MAX_SIZE = 16                  // Must be <= max_uint64 (2^64 - 1)
const QUEUE_SAVE_INTERVAL = 60             // Auto-Save Interval in Seconds
const QUEUE_ITEM_DATA_MAXLEN = 1024*64 - 1 // Must be <= max_uint16 (2^16 - 1)
const BUFFER_SIZE = QUEUE_ITEM_DATA_MAXLEN + 5

// Client, being idle for this Time, is disconnected. Interval is in Seconds.
const CONNECTION_IDLE_TIMEOUT = 30

// Size of Queue of incomming Connections to Server
const TASK_QUEUE_SIZE = 128

const CMD_GET byte = 'g'              // Get Item from Queue
const CMD_PUT byte = 'p'              // Put Item to Queue
const CMD_SIZE byte = 's'             // Show actual (filled) Size of Queue
const CMD_CLOSE_CONNECTION byte = 'q' // Close current Connection
const CMD_SEPARATOR byte = ' '        // Separator in Client's Request

const REPLY_CLOSED_CONNECTION byte = 'd' // Closed current Connection
const REPLY_QUEUE_IS_EMPTY byte = 'e'    // Reply if Queue is empty
const REPLY_QUEUE_IS_FULL byte = 'f'     // Reply if Queue is full
const REPLY_ITEM byte = 'i'              // Reply when Item is shown
const REPLY_OK byte = 'k'                // Reply if OK
const REPLY_ERROR byte = 'r'             // Reply if Error
const REPLY_SIZE byte = 's'              // Reply when Size is shown
const REPLY_SEPARATOR byte = ' '         // Separator in Server's Reply

type tTask struct {
	RcvdTime      int64
	Connection    net.Conn
	ClientAddress string
	ClientPort    string
}

type tQueueItem struct {
	CreationTime int64
	//Data         string // Maximum Length is QUEUE_ITEM_MAX_LENGTH
	DataLength uint16
	Data       []byte
}

var TaskChan chan tTask
var Queue chan tQueueItem
var QueueSize uint64
var QueueLST int64 // Last Save Time

//------------------------------------------------------------------------------

func main() {

	var ConnectionAddress string
	var ConnectionHost = "127.0.0.1"
	var ConnectionPort = "5555"
	var ConnectionType = "tcp"

	// Create Channels
	TaskChan = make(chan tTask, TASK_QUEUE_SIZE)

	// Preparations
	QueueInit()

	// Start Connection Handler
	go ConnectionHandler()

	// Listen
	ConnectionAddress = ConnectionHost + PORT_DELIMITER + ConnectionPort
	Listen(ConnectionType, ConnectionAddress)
}

//------------------------------------------------------------------------------

func Listen(ConnectionType, ConnectionAddress string) {

	// Listens to incoming Connections.

	var ClientAddress, ClientIPA, ClientPort string
	var Connection net.Conn
	var err error
	var Listener net.Listener
	var PortDelimiterPos int
	var Task *tTask

	Task = new(tTask)

	// Create a Listener
	Listener, err = net.Listen(ConnectionType, ConnectionAddress)
	if err != nil {
		log.Println("Error. Can not create Listener at", ConnectionAddress, err) //
		return
	}

	log.Println("Started TCP Server at", ConnectionAddress) //

	// Accept Connections in Loop & send Tasks
	for {

		Connection, err = Listener.Accept()
		if err != nil {

			log.Println("Error. Can not accept Connection.", err) //
			return
		}

		// Filter Connection by IP Address
		ClientAddress = Connection.RemoteAddr().String()
		PortDelimiterPos = strings.Index(ClientAddress, PORT_DELIMITER)
		ClientIPA = ClientAddress[:PortDelimiterPos]
		ClientPort = ClientAddress[PortDelimiterPos+1:]

		//fmt.Println("Incomming Connection from IP:", ClientIPA, "; Port:", ClientPort) // dbg

		if ClientIPA != CLIENT_IPA_ALLOWED {
			log.Println("Achtung! Connection from dis-allowed IP Address:", ClientIPA) //
		}

		// Create Task
		Task.Connection = Connection
		Task.ClientAddress = ClientIPA
		Task.ClientPort = ClientPort
		Task.RcvdTime = time.Now().Unix() // simulate some Activity

		// Send Task
		TaskChan <- *Task
	}
}

//------------------------------------------------------------------------------

func ConnectionHandler() {

	// Handles Connections.
	// If no Errors occur, then Connection stays alive (not closed).

	var Task tTask
	var ClientAddress, ClientIPA, ClientPort string
	var RequestBuffer []byte
	var Connection net.Conn
	//var RcvdTime int64
	var err error
	var neterr net.Error
	var ReplyLength, RequestLength int
	var Reply []byte
	var NowUnix int64
	var cmd byte
	var DeadLine, NowTime time.Time
	var IdleTimeout time.Duration
	var ok bool

	RequestBuffer = make([]byte, BUFFER_SIZE)
	IdleTimeout = time.Second * CONNECTION_IDLE_TIMEOUT

	for {

		// Read Task
		Task = <-TaskChan
		Connection = Task.Connection
		//RcvdTime = Task.RcvdTime
		ClientIPA = Task.ClientAddress
		ClientPort = Task.ClientPort
		ClientAddress = ClientIPA + PORT_DELIMITER + ClientPort

		// Refresh Idle Deadline
		NowTime = time.Now()
		NowUnix = NowTime.Unix()
		DeadLine = NowTime.Add(IdleTimeout)
		Connection.SetDeadline(DeadLine)

		// Process a Connected Client
		for {

			// Save Queue if needed
			if NowUnix-QueueLST > QUEUE_SAVE_INTERVAL {
				QueueSave()
			}

			// Read Request
			RequestLength, err = Connection.Read(RequestBuffer)

			if err != nil {

				// Client ended Connection
				if err == io.EOF {

					log.Println("Warning. Connection terminated by Client (",
						ClientAddress, ").") //
					// Abort the same way as Client
					Connection.Close() // Brute Close
					break
				}

				// Check Idle Timeout
				neterr, ok = err.(net.Error)
				if ok && neterr.Timeout() {

					log.Println("Warning. Connection timed out (",
						ClientAddress, ").") //
					Connection.Close() // Brute Close
					break
				}

				log.Println("Achtung! Broken Connection (",
					ClientAddress, ").", err) //
				CloseConnection(Connection)
				break
			}
			if RequestLength == 0 {

				log.Println("Error. Empty Request.") //
				CloseConnection(Connection)
				break
			}

			// Process Request
			cmd = RequestBuffer[0]

			// 1. Action : Close Connection
			if cmd == CMD_CLOSE_CONNECTION {

				CloseConnection(Connection)
				break
			}

			// 2. Other Actions
			ProcessRequest(RequestBuffer, RequestLength, &Reply, NowUnix, cmd)

			// Refresh Idle Deadline
			NowTime = time.Now()
			NowUnix = NowTime.Unix()
			DeadLine = NowTime.Add(IdleTimeout)
			Connection.SetDeadline(DeadLine)

			//fmt.Println("Reply:", Reply) // dbg

			// Send Reply
			ReplyLength, err = Connection.Write(Reply)
			if err != nil {

				log.Println("Error. Can not write Reply.", err) //
				CloseConnection(Connection)
				break
			}
			if ReplyLength == 0 {

				log.Println("Error. Reply is empty.") //
				CloseConnection(Connection)
				break
			}
		}
	}
}

//------------------------------------------------------------------------------

func ProcessRequest(RequestBuffer []byte, RequestLength int, Reply *[]byte, Now int64, cmd byte) {

	// Processes incoming Request.

	var t1, t2 byte
	var ItemSize, i int
	var QueueItem *tQueueItem
	var ValueSize uint16

	// Process Request
	// Show Request
	//fmt.Println("Request:", string(*RequestBuffer)) // dbg

	// GET
	if cmd == CMD_GET {

		if RequestLength != 1 {

			log.Println("Protocol Error. 'Get' needs no Parameters.") //

			// Reply: Error
			*Reply = make([]byte, 1)
			(*Reply)[0] = REPLY_ERROR
			return
		}

		// Get Next Item from Queue
		if QueueSize == 0 {

			//fmt.Println("Warning. Queue is empty.") // dbg

			// Reply: Empty Queue
			*Reply = make([]byte, 1)
			(*Reply)[0] = REPLY_QUEUE_IS_EMPTY
			return
		}

		QueueItem = new(tQueueItem)
		*QueueItem = <-Queue
		QueueSize--

		// Reply: Item
		ItemSize = len(QueueItem.Data)

		*Reply = make([]byte, ItemSize+5)
		(*Reply)[0] = REPLY_ITEM
		(*Reply)[1] = REPLY_SEPARATOR
		(*Reply)[2] = byte(uint16(ItemSize) >> 8)
		(*Reply)[3] = byte(uint16(ItemSize) % 256)
		(*Reply)[4] = REPLY_SEPARATOR
		for i = 0; i < ItemSize; i++ {
			(*Reply)[5+i] = QueueItem.Data[i]
		}
		return
	}

	// PUT
	if cmd == CMD_PUT {

		// Read First Separator
		t1 = RequestBuffer[1]
		if t1 != CMD_SEPARATOR {

			log.Println("Protocol Error. Wrong Symbol after 'Put'.") //

			// Reply: Error
			*Reply = make([]byte, 1)
			(*Reply)[0] = REPLY_ERROR
			return
		}

		// Read Size of Value (uint16)
		t1 = RequestBuffer[2]
		t2 = RequestBuffer[3]
		ValueSize = uint16(t1)*256 + uint16(t2)
		if ValueSize > QUEUE_ITEM_DATA_MAXLEN {

			log.Println("Error. Too long Value requested.") //

			// Reply: Error
			*Reply = make([]byte, 1)
			(*Reply)[0] = REPLY_ERROR
			return
		}

		// Queue Size?
		if QueueSize == QUEUE_MAX_SIZE {

			//fmt.Println("Warning. The Queue is full.") // dbg

			// Reply: Full Queue
			*Reply = make([]byte, 1)
			(*Reply)[0] = REPLY_QUEUE_IS_FULL
			return
		}

		//fmt.Println("ValueSize:", ValueSize) // dbg

		// Read Second Separator
		t1 = RequestBuffer[4]
		if t1 != CMD_SEPARATOR {

			log.Println("Protocol Error. 'Put' bad Syntax.") //

			// Reply: Error
			*Reply = make([]byte, 1)
			(*Reply)[0] = REPLY_ERROR
			return
		}

		// Create an Item
		QueueItem = new(tQueueItem)
		QueueItem.CreationTime = Now
		QueueItem.DataLength = ValueSize
		QueueItem.Data = make([]byte, ValueSize)
		copy(QueueItem.Data, RequestBuffer[5:5+ValueSize])

		// Put an Item to Queue
		Queue <- *QueueItem
		QueueSize++

		// Reply: OK
		*Reply = make([]byte, 1)
		(*Reply)[0] = REPLY_OK
		return
	}

	// SIZE
	if cmd == CMD_SIZE {

		if RequestLength != 1 {

			log.Println("Protocol Error. 'Show' needs no Parameters.") //

			// Reply: Error
			*Reply = make([]byte, 1)
			(*Reply)[0] = REPLY_ERROR
			return
		}

		// Reply: Queue Size
		*Reply = make([]byte, 6)
		(*Reply)[0] = REPLY_SIZE
		(*Reply)[1] = REPLY_SEPARATOR
		(*Reply)[2] = byte(QueueSize >> 24)
		(*Reply)[3] = byte(QueueSize >> 16)
		(*Reply)[4] = byte(QueueSize >> 8)
		(*Reply)[5] = byte(QueueSize % 256)
		return
	}
}

//------------------------------------------------------------------------------

func QueueInit() {

	// Creates Queue.

	QueueSize = 0
	// Create Channel
	Queue = make(chan tQueueItem, QUEUE_MAX_SIZE)

	//fmt.Println("Loading Queue...") // dbg
	// Load Queue
	// ...
	// not yet implemented, as it depends on the Structure of 'tQueueItem'
	// ...
	//fmt.Println("Loading Done.") // dbg

	QueueLST = time.Now().Unix()
}

//------------------------------------------------------------------------------

func QueueSave() {

	// Saves the Queue from Memory to Disk.

	//fmt.Println("Saving Queue...") // dbg
	// Save Queue
	// ...
	// not yet implemented, as it depends on the Structure of 'tQueueItem'
	// ...
	//fmt.Println("Saving Done.") // dbg

	QueueLST = time.Now().Unix()
}

//------------------------------------------------------------------------------

func CloseConnection(Connection net.Conn) {

	// Closes Connection.

	var Reply []byte
	var ReplyLength int
	var err error

	Reply = make([]byte, 1)
	Reply[0] = REPLY_CLOSED_CONNECTION
	ReplyLength, err = Connection.Write(Reply)
	if err != nil {

		log.Println("Error. Can not write Reply.", err) //
	}
	if ReplyLength == 0 {

		log.Println("Error. Reply is empty.") //
	}

	Connection.Close()
}

//------------------------------------------------------------------------------
