// 13.go

package main

import "fmt"
import "time"

func main() {

	var myChan chan string

	myChan = make(chan string, 10)

	// Start a Writer.
	go writer(myChan)

	// Start Reader.
	//go reader_a(myChan)
	//go reader_b(myChan)
	go reader_c(myChan)

	time.Sleep(time.Second * 5)
}

// Sends Messages to a Channel.
func writer(ch chan string) {

	fmt.Println("Writer Start.")
	ch <- "a"
	ch <- "b"
	close(ch)
	//ch <- "c" // Leads to an Error.
	fmt.Println("Writer Stop.")
}

// Reads Messages from a Channel.
func reader_a(ch chan string) {

	var msg string
	var msgIsReceived bool

	fmt.Println("Reader A Start.")

	msg, msgIsReceived = <-ch
	for msgIsReceived == true {

		// Process incoming Message.
		fmt.Println(msg)

		// Read next Message.
		msg, msgIsReceived = <-ch
	}

	fmt.Println("Reader A Stop.")
}

// Reads Messages from a Channel.
func reader_b(ch chan string) {

	var msg string

	fmt.Println("Reader B Start.")

	for msg = range ch {
		fmt.Println(msg)
	}

	fmt.Println("Reader B Stop.")
}

// Reads Messages from a Channel.
func reader_c(ch chan string) {

	var msg string
	var msgIsReceived bool

	fmt.Println("Reader C Start.")

	for true {

		select {

		case msg, msgIsReceived = <-ch:

			// This Branch is selected when:
			//	1. Message is received;
			//	2. Channel is closed.

			if msgIsReceived == true {
				fmt.Println(msg)
			}

		default:

			// This Branch is selected when:

			//	1. Channel is open and Message is not received.

			fmt.Println("Default Branch.")
		}

	}

	fmt.Println("Reader C Stop.")
}
