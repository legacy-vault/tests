// sender.go.

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
	var msg amqp.Publishing
	var msgBA []byte
	var msgExchange string
	var msgDeliveryIsImmediate bool
	var msgDeliveryIsMandatory bool
	var msgRouteKey string
	var msgStr string
	var queue amqp.Queue

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

	// Prepare a Message.
	msgStr = "Hello World"
	msgBA = []byte(msgStr)
	msg = amqp.Publishing{}
	msg.Body = msgBA
	msg.ContentType = "text/plain"
	//
	msgExchange = ""
	msgRouteKey = queue.Name
	msgDeliveryIsMandatory = false
	msgDeliveryIsImmediate = false

	// Send the Message.
	err = channel.Publish(
		msgExchange,
		msgRouteKey,
		msgDeliveryIsMandatory,
		msgDeliveryIsImmediate,
		msg)
	if err != nil {
		log.Fatalln(err)
	}
}
