// sender.go.

package main

import (
	"github.com/streadway/amqp"
)

import (
	"log"
	"time"
	"fmt"
)

func main() {

	const ExchangeName = "x1"

	var channel *amqp.Channel
	var conn *amqp.Connection
	var err error
	var failedQueues chan string
	var failedQueuesChanSize uint64
	var msg amqp.Publishing
	var msgBA []byte
	var msgStr string
	var queueA amqp.Queue
	var returnedMessages chan amqp.Return
	var returnedMessagesChanSize uint64
	var sentMessages chan amqp.Confirmation
	var sentMessagesChanSize uint64
	var quitChan chan bool

	// Connect to the Server.
	fmt.Print("Connecting to RabbitMQ AMQP Server...")
	conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Done.")
	defer conn.Close()

	// Open a Channel.
	channel, err = conn.Channel()
	if err != nil {
		log.Fatalln(err)
	}
	defer channel.Close()

	// Configure auxiliary Handlers...
	err = channel.Confirm(false)
	if err != nil {
		log.Fatalln(err)
	}

	// 1. Handler for returned Messages.
	returnedMessagesChanSize = 100
	returnedMessages = make(chan amqp.Return, returnedMessagesChanSize)
	returnedMessages = channel.NotifyReturn(returnedMessages)
	go returnedMessagesHandler(returnedMessages)

	// 2. Handler for sent Messages.
	sentMessagesChanSize = 100
	sentMessages = make(chan amqp.Confirmation, sentMessagesChanSize)
	sentMessages = channel.NotifyPublish(sentMessages)
	go sentMessagesHandler(sentMessages)

	// 3. Handler for Queues which have been closed.
	failedQueuesChanSize = 100
	failedQueues = make(chan string, failedQueuesChanSize)
	failedQueues = channel.NotifyCancel(failedQueues)
	go closedQueuesHandler(failedQueues)

	quitChan = make(chan bool)

	// Delete a fake non-existent Exchange.
	// According to the Documentation, this should return an Error,
	// but by some strange Reason, it does not...
	fmt.Print("Deleting non-existent Exchange...")
	err = channel.ExchangeDelete("fake_exchange", false, false)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Done.")

	// Delete an old Exchange.
	fmt.Print("Deleting an old Exchange...")
	err = channel.ExchangeDelete(ExchangeName, false, false)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Done.")

	//!
	time.Sleep(time.Second * 15) //!

	// Declare a completely new Exchange.
	err = channel.ExchangeDeclare(
		ExchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Declare a Queue.
	queueA, err = channel.QueueDeclare(
		"Q1",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Bind Queue to Exchange.
	err = channel.QueueBind(
		queueA.Name,
		"news.*",
		ExchangeName,
		false,
		nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Prepare a Message.
	msgStr = "Good News! Cows can fly!"
	msgBA = []byte(msgStr)
	msg = amqp.Publishing{}
	msg.Body = msgBA
	msg.ContentType = "text/plain"
	msg.DeliveryMode = amqp.Persistent
	msg.Timestamp = time.Now()

	// Send the Message to Exchange.
	fmt.Println("Sending a Message...")
	err = channel.Publish(
		ExchangeName,
		"news.good",
		true,
		false,
		msg)
	if err != nil {
		log.Fatalln(err)
	}

	<-quitChan
}

// Handles Queues which have been closed.
func closedQueuesHandler(ch chan string) {

	var queue string

	fmt.Println("Closed Queues Handler has started.")

	for {
		queue = <-ch
		log.Printf("Closed Queue: %+v.\r\n", queue)
	}
}

// Handles Messages which failed to be sent.
func returnedMessagesHandler(ch chan amqp.Return) {

	var ret amqp.Return

	fmt.Println("Returned Messages Handler has started.")

	for {
		ret = <-ch
		log.Printf("Returned Message: %+v.\r\n", ret)
	}
}

// Handles Messages which were successfully sent.
func sentMessagesHandler(ch chan amqp.Confirmation) {

	var cnfrm amqp.Confirmation

	fmt.Println("Sent Messages Handler has started.")

	for {
		cnfrm = <-ch
		log.Printf("Sent Message: %+v.\r\n", cnfrm)
	}
}
