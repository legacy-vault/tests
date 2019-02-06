// ftp_retr_stress.go.

package main

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"time"
	//"io/ioutil"
	"log"
	"os"
)

func main() {

	var err error

	err = TestFtpRetr()
	if err != nil {
		log.Println(err)
	}
}

func TestFtpRetr() error {

	const Env_FtpAddress = "FTP_ADDRESS"

	var address string
	var connection *ftp.ServerConn
	var err error
	var fileContents []byte
	var logoutIsRequired bool
	var password string
	var remotePath string
	var username string

	address = os.Getenv(Env_FtpAddress)
	username = "..."
	password = "..."
	logoutIsRequired = false
	remotePath = "..."

	connection, err = ftp.Connect(address)
	if err != nil {
		return err
	}

	err = connection.Login(username, password)
	if err != nil {
		return err
	}

	err = RetrStress(
		connection,
		10,
		10,
		remotePath,
	)
	if err != nil {
		return err
	}

	if logoutIsRequired {
		err = connection.Logout()
		if err != nil {
			return err
		}
	}

	err = connection.Quit()
	if err != nil {
		return err
	}

	fmt.Println(remotePath, len(fileContents), "Bytes")

	return nil
}

func RetrStress(
	connection *ftp.ServerConn,
	testCount int,
	retryCountMax int,
	remotePath string,
) error {

	var err error
	var fileStream *ftp.Response
	var retry bool
	var retryCounter int

	for i := 1; i <= testCount; i++ {

		fileStream, err = connection.Retr(remotePath)
		if err != nil {
			return err
		}

		err = fileStream.Close()
		if err != nil {
			return err
		}

		continue

		/*
			fileContents, err = ioutil.ReadAll(fileStream)
			if err != nil {
				return err
			}
		*/

		retry = false
		retryCounter = 0
		err = fileStream.Close()
		if err != nil {
			retry = true
		}
		for retry {
			retryCounter++
			err = fileStream.Close()
			if err == nil {
				break
			}
			if retryCounter > retryCountMax {
				return err
			}
			log.Println("Sleeping...")
			time.Sleep(time.Millisecond * 1)
		}
	}

	return nil
}
