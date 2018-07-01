// 34.go.

// Example of Delaying Messages.

// If you send a Message with short Body to Receiver,
// it Acks the Message.
// Messages whose body is longer than 3 Symbols
// are postponed to Delay Channel.
// Delay Channel has a TTL and Dead-Letter Setting
// which send the Letter back to its Origin using the Routing Key.
// Long Letters start moving in a Circle (just to illustrate the Effect).

// 2018-07-02.

package main

import (
	"github.com/streadway/amqp"
)

import (
	"log"
	"time"
)

const ExchangeNameRcv = "receiver"
const QueueNamePrefixRcv = "receiver-"
const QueueNameRcvA = "receiver-a"
const QueueNameRcvB = "receiver-b"
const ExchangeNameDelay = "delay"
const QueueNameDelayA = "delay-a"

func main() {

	var channelRcv *amqp.Channel
	var channelDelay *amqp.Channel
	var conn *amqp.Connection
	var dsn string
	var err error
	var messagesRcvA, messagesRcvB <-chan amqp.Delivery
	var paramsDelayA amqp.Table
	var returnedMessagesRcv chan amqp.Return
	var returnedMessagesDelay chan amqp.Return
	var queueRcvA, queueRcvB, queueDelayA amqp.Queue
	var quitChan chan bool
	var sentMessagesRcv chan amqp.Confirmation
	var sentMessagesDelay chan amqp.Confirmation

	// Connect to the Server.
	dsn = "amqp://guest:guest@localhost:5672/"
	conn, err = amqp.Dial(dsn)
	checkError(err)
	defer conn.Close()

	// Channel.
	// 1. Rcv.
	channelRcv, err = conn.Channel()
	checkError(err)
	defer channelRcv.Close()
	err = channelRcv.Confirm(false)
	checkError(err)
	// 2. Delay.
	channelDelay, err = conn.Channel()
	checkError(err)
	defer channelDelay.Close()
	err = channelDelay.Confirm(false)
	checkError(err)

	// Go-Channels.
	// 1. Returned Messages.
	returnedMessagesRcv = make(chan amqp.Return, 100)
	returnedMessagesRcv =
		channelRcv.NotifyReturn(returnedMessagesRcv)
	returnedMessagesDelay = make(chan amqp.Return, 100)
	returnedMessagesDelay =
		channelDelay.NotifyReturn(returnedMessagesDelay)
	go returnedMessagesHandler(returnedMessagesRcv, "RcvChannel")
	go returnedMessagesHandler(returnedMessagesDelay, "DelayChannel")
	// 2. Sent Messages.
	sentMessagesRcv = make(chan amqp.Confirmation, 100)
	sentMessagesRcv = channelRcv.NotifyPublish(sentMessagesRcv)
	sentMessagesDelay = make(chan amqp.Confirmation, 100)
	sentMessagesDelay = channelRcv.NotifyPublish(sentMessagesDelay)
	go sentMessagesHandler(sentMessagesRcv, "RcvChannel")
	go sentMessagesHandler(sentMessagesDelay, "DelayChannel")

	// Exchange.
	// 1. Rcv.
	err = channelRcv.ExchangeDeclare(
		ExchangeNameRcv,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	checkError(err)
	// 2. Delay.
	err = channelDelay.ExchangeDeclare(
		ExchangeNameDelay,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	checkError(err)

	// Queue.
	// 1. Rcv.
	queueRcvA, err = channelRcv.QueueDeclare(
		QueueNameRcvA,
		true,
		false,
		false,
		false,
		nil,
	)
	checkError(err)
	queueRcvB, err = channelRcv.QueueDeclare(
		QueueNameRcvB,
		true,
		false,
		false,
		false,
		nil,
	)
	checkError(err)
	// 2. Delay.
	paramsDelayA = make(amqp.Table)
	paramsDelayA["x-message-ttl"] = int32(30000)
	paramsDelayA["x-dead-letter-exchange"] = ExchangeNameRcv
	queueDelayA, err = channelDelay.QueueDeclare(
		QueueNameDelayA,
		true,
		false,
		false,
		false,
		paramsDelayA,
	)
	checkError(err)

	// Queue Binding.
	// 1. Rcv.
	err = channelRcv.QueueBind(
		queueRcvA.Name,
		queueRcvA.Name,
		ExchangeNameRcv,
		false,
		nil,
	)
	checkError(err)
	err = channelRcv.QueueBind(
		queueRcvB.Name,
		queueRcvB.Name,
		ExchangeNameRcv,
		false,
		nil,
	)
	checkError(err)
	// 2. Delay.
	err = channelDelay.QueueBind(
		queueDelayA.Name,
		"#",
		ExchangeNameDelay,
		false,
		nil,
	)
	checkError(err)

	// Consume.
	messagesRcvA, err = channelRcv.Consume(
		queueRcvA.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	checkError(err)
	messagesRcvB, err = channelRcv.Consume(
		queueRcvB.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	checkError(err)

	// Manager.
	quitChan = make(chan bool)
	go receiveMessages(messagesRcvA, "a", channelDelay)
	go receiveMessages(messagesRcvB, "b", channelDelay)
	<-quitChan
}

// Receives Messages.
func receiveMessages(
	chIn <-chan amqp.Delivery,
	name string,
	chOut *amqp.Channel,
) {

	var err error
	var msg amqp.Delivery
	var routeKey string
	var str string

	routeKey = QueueNamePrefixRcv + name
	log.Println("Receiver", name, "has started")

	for msg = range chIn {

		str = string(msg.Body)

		log.Printf(
			"[Receiver %s] Incoming Message: %s\r\n%+v\r\n",
			name,
			str,
			msg,
		)
		// Simulate some Work.
		time.Sleep(time.Second * 15)

		if len(str) > 3 {

			// Postpone long Message.
			err = moveMessage(
				chOut,
				ExchangeNameDelay,
				routeKey,
				&msg,
			)
			checkError(err)

		} else {

			// Acknowledge Message Reception.
			err = msg.Ack(false)
			checkError(err)
			log.Println("Message Ack.")
		}
	}

	log.Println("Message Receiver has stopped.")
}

// Handles Messages which were successfully sent.
func sentMessagesHandler(
	ch chan amqp.Confirmation,
	name string,
) {
	var cnfrm amqp.Confirmation

	for cnfrm = range ch {

		log.Printf(
			"[%s] Message has been sent: %+v.\r\n",
			name,
			cnfrm,
		)
	}
}

// Handles Messages which failed to be sent.
func returnedMessagesHandler(
	ch chan amqp.Return,
	name string,
) {
	var ret amqp.Return

	for ret = range ch {

		log.Printf(
			"[%s] Message has been returned: %+v.\r\n",
			name,
			ret,
		)
	}
}

// Checks an Error.
func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// Moves Message to another Exchange.
func moveMessage(
	ch *amqp.Channel,
	xchangeName string,
	routeKey string,
	msg *amqp.Delivery,
) error {

	var err error
	var pub amqp.Publishing

	// Prepare a Letter.
	pub.Body = msg.Body
	pub.ContentType = "text/plain"
	pub.DeliveryMode = amqp.Persistent
	pub.Timestamp = time.Now()

	// Send the Letter.
	err = ch.Publish(
		xchangeName,
		routeKey,
		true,
		false,
		pub,
	)
	checkError(err)

	// Wait for Delivery Confirmation.
	// Not implemented.

	// Acknowledge.
	err = msg.Ack(false)
	checkError(err)
	log.Println("Message Ack.")

	return nil
}
