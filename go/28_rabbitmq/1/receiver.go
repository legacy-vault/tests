// receiver.go.

package main

import (
	"github.com/streadway/amqp"
)

import (
	"log"
)

func main() {

	var channel *amqp.Channel
	var conn *amqp.Connection
	var err error
	var messages <-chan amqp.Delivery
	var msgAutoAck bool
	var msgConsumer string
	var msgIsExclusive bool
	var msgNoLocal bool
	var msgNoWait bool
	var msgQueue string
	var queue amqp.Queue
	var quitChan chan bool

	// Connect to the Server.
	conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// Open a Channel.
	channel, err = conn.Channel()
	if err != nil {
		log.Fatalln(err)
	}
	defer channel.Close()

	// Declare a Queue.
	queue, err = channel.QueueDeclare(
		"The Queue",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Prepare to read.
	msgQueue = queue.Name
	msgConsumer = ""
	msgAutoAck = true
	msgIsExclusive = false
	msgNoLocal = false
	msgNoWait = false

	// "Connect" to Channel.
	messages, err = channel.Consume(msgQueue,
		msgConsumer,
		msgAutoAck,
		msgIsExclusive,
		msgNoLocal,
		msgNoWait,
		nil)
	if err != nil {
		log.Fatalln(err)
	}

	quitChan = make(chan bool)

	go receiveMessages(messages)

	<-quitChan
}

func receiveMessages(messagesChannel <-chan amqp.Delivery) {

	var msg amqp.Delivery

	for msg = range messagesChannel {
		log.Printf("Incoming Message.\r\n%v\r\n%v\r\n%+v\n\n",
			msg.Body, string(msg.Body), msg)
	}
}
