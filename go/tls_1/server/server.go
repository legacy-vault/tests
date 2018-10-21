// server.go

package main

import (
	// "fmt"
	// "io"
	"log"
	"net/http"
)

//------------------------------------------------------------------------------

const SERVER_CERTIFICATE_PATH = "../key/1/localhost.crt"
const SERVER_KEY_PATH = "../key/1/localhost.key"
const SERVER_LISTEN_HOST = ""
const SERVER_LISTEN_PORT = "2000"

//------------------------------------------------------------------------------

var server_address string

//------------------------------------------------------------------------------

func ServeFolderRoot(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("HTTPS works!\r\n"))
	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}

//------------------------------------------------------------------------------

func main() {

	var err error

	server_address = SERVER_LISTEN_HOST + ":" + SERVER_LISTEN_PORT

	http.HandleFunc("/", ServeFolderRoot)

	// Start Server.
	err = http.ListenAndServeTLS(server_address,
		SERVER_CERTIFICATE_PATH, SERVER_KEY_PATH, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//------------------------------------------------------------------------------
