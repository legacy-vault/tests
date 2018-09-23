// main.go.

// Minio Database Test.

// Date: 2018-09-14.

package main

import (
	"log"
	"os"
	myMinio "test/9_minio/minio"
)

func main() {

	var err error

	err = myMinio.TestA()
	check_error(err)
}

func check_error(e error) {

	if e != nil {
		log.Println(e.Error())
		os.Exit(1)
	}
}
