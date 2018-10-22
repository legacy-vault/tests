// server.go.

package main

import (
	//"io"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
	//"github.com/gorilla/websocket"
)

func main() {

	var err error

	//http.HandleFunc("/ws", wsHandler)
	http.Handle("/test", websocket.Handler(handlerTest))
	err = http.ListenAndServe("0.0.0.0:2000", nil)
	if err != nil {
		log.Fatalln("ListenAndServe: " + err.Error())
	}
}

func handlerTest(ws *websocket.Conn) {

	var buffer []byte
	var bytesRead int
	var err error
	var msg []byte
	var reply []byte

	// Read a Message.
	buffer = make([]byte, 128)
	bytesRead, err = ws.Read(buffer)
	if err != nil {
		log.Println(err)
	}
	msg = make([]byte, bytesRead)
	copy(msg, buffer[:bytesRead])
	fmt.Println(bytesRead, "Bytes Request:", msg)

	// Prepare the Reply.
	reply = make([]byte, bytesRead+1)
	copy(reply, msg)
	reply[bytesRead] = '*'
	fmt.Println(bytesRead+1, "Bytes Response:", reply)

	// Write a Message.
	_, err = ws.Write(reply)
	if err != nil {
		log.Println(err)
	}
}
