// receiver.go.

package main

import (
	"github.com/streadway/amqp"
)

import (
	"log"
	"fmt"
	"time"
)

func main() {

	const ExchangeName = "x1"

	var channel *amqp.Channel
	var conn *amqp.Connection
	var err error
	var messages <-chan amqp.Delivery
	var queueA amqp.Queue
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

	// Declare an Exchange.
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

	// "Connect" to Channel's Queue.
	messages, err = channel.Consume(
		queueA.Name,
		"",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Fatalln(err)
	}

	quitChan = make(chan bool)

	go receiveMessages(messages)

	<-quitChan
}

func receiveMessages(messagesChannel <-chan amqp.Delivery) {

	var err error
	var msg amqp.Delivery

	for msg = range messagesChannel {
		log.Printf("Incoming Message.\r\n%v\r\n%v\r\n%+v\n\n",
			msg.Body, string(msg.Body), msg)

		// Simulate some Work.
		time.Sleep(time.Second * 15) //!

		// Acknowledge Message Reception.
		err = msg.Ack(false)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
