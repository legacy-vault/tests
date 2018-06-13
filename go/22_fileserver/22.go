// 22.go.

/*

	This a simple Example of Golang's built-in File-Server Usage.

	The HTTP Server combines File Handler and ordinary Handler.
	Moreover, a simple Counter collects Read Counts for each requested File.

	To show Read Count: set the 'act' Parameter to '2'.
	Example:
		http://localhost:12345/files/dir/2.txt?act=2

	To get a File: either set the 'act' Parameter to '1' or do not set it.
	Example:
		http://localhost:12345/files/dir/2.txt?act=1
		http://localhost:12345/files/dir/2.txt

*/

package main

import (
	"net/http"
	"time"
	"log"
	"strings"
	"fmt"
	"net/url"
	"strconv"
)

const FILESERVER_ROOT_PATH = "./files_storage/"
const FILESERVER_URL_PREFIX = "/files/"

const URL_PARAM_ACTION = "act"

const ACTION_FILE_GET = 1
const ACTION_FILE_STAT_SHOW = 2
const ACTION_DEFAULT = ACTION_FILE_GET

var fileHandler http.Handler
var fileReadCounter map[string]uint64

func main() {

	var err error
	var srv *http.Server

	fileReadCounter = make(map[string]uint64)

	srv = &http.Server{}
	srv.Addr = "0.0.0.0:12345"
	srv.Handler = http.HandlerFunc(httpHandler)
	srv.ReadTimeout = time.Second * 60
	srv.WriteTimeout = time.Second * 60

	// Create a Handler which serves the Files' Storage Directory.
	fileHandler = http.FileServer(http.Dir(FILESERVER_ROOT_PATH))

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func httpHandler(w http.ResponseWriter, r *http.Request) {

	var actionType uint64
	var actionTypeStr string
	var actionTypeStrs []string
	var buf string // Using this only for Simplicity.
	var count uint64
	var err error
	var exists bool
	var pathForFileServer string
	var reqURLPath string
	var reqURLParams url.Values

	// Request's URL & Parameters.
	reqURLPath = r.URL.Path
	reqURLParams, err = url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Fatal(err)
	}

	// Read 'Action Type' Parameter.
	// If Reading fails, then fall back to the default Action.
	actionTypeStrs, exists = reqURLParams[URL_PARAM_ACTION]
	if exists {

		// 'Action Type' Parameter is set in Request.
		actionTypeStr = actionTypeStrs[0]
		actionType, err = strconv.ParseUint(actionTypeStr, 10, 64)
		if err != nil {
			// 'Action Type' Parameter is unrecognized or damaged.
			actionType = ACTION_DEFAULT
		}

	} else {

		// 'Action Type' Parameter is not set in Request.
		actionType = ACTION_DEFAULT
	}

	// Check URL's Path.
	if (strings.HasPrefix(reqURLPath, FILESERVER_URL_PREFIX)) {

		// URL matches the File Server Path.
		pathForFileServer = strings.TrimPrefix(reqURLPath, FILESERVER_URL_PREFIX)

		switch actionType {

		case ACTION_FILE_GET:

			// Change Read Counter.
			count, exists = fileReadCounter[pathForFileServer]
			if !exists {
				fileReadCounter[pathForFileServer] = 0
			}
			fileReadCounter[pathForFileServer]++

			// Change URL's Path for the File Server to recognize it.
			r.URL.Path = pathForFileServer

			// Get the File.
			fileHandler.ServeHTTP(w, r)

		case ACTION_FILE_STAT_SHOW:
			count, exists = fileReadCounter[pathForFileServer]
			if exists {

				// Counter exists.
				buf = "File [" + pathForFileServer + "] has been requested [" +
					strconv.FormatUint(count, 10) + "] times."

				fmt.Fprint(w, buf)

			} else {

				// Counter is empty.
				buf = "File [" + pathForFileServer + "] has not been" +
					" requested."

				fmt.Fprint(w, buf)
			}

		default:
			fmt.Fprint(w, "Action Type is not supported.")
		}

	} else {

		// URL does not match the File Server Path.

		// Send Reply.
		fmt.Fprint(w, "You have requested an ordinary Resource :-)")
	}
}
