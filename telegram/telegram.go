// telegram.go

package main

// Telegram Sender.

//-----------------------------------------------------------------------------|

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

//-----------------------------------------------------------------------------|

// Basic Parameters.
const TELEGRAM_API_HOST = "api.telegram.org"
const TELEGRAM_API_FUNC_MESSAGE_SEND = "sendMessage"
const TELEGRAM_API_PATH = "/bot"
const TELEGRAM_API_PORT = "443"
const TELEGRAM_API_PROTOCOL = "https"
const TELEGRAM_API_PROTOCOL_VERSION_FULL = "HTTP/1.1"
const TELEGRAM_API_PROTOCOL_METHOD = "POST"
const TELEGRAM_API_REQUEST_USERAGENT = "Sender"

// Composite Parameters.
const TELEGRAM_API_ADDRESS = TELEGRAM_API_HOST + ":" + TELEGRAM_API_PORT

const TELEGRAM_BOT_HT_USERNAME = "@your_bot_alias_here"
const TELEGRAM_BOT_HT_TOKEN = "your_bot_token_here"

const TELEGRAM_CONNECTIONS_IDLE_MAX = 10

const TELEGRAM_TIMEOUT_CONNECTION_IDLE = 30
const TELEGRAM_TIMEOUT_EXPECT_CONTINUE = 30
const TELEGRAM_TIMEOUT_HTTP_GENERAL = 30
const TELEGRAM_TIMEOUT_RESPONSE_HEADER = 30

const SYMBOL_BACKSLASH = "\\"
const SYMBOL_BACKSLASH_DOUBLEQUOTE = "\\\""
const SYMBOL_BACKSLASH_X2 = "\\\\"
const SYMBOL_DOUBLEQUOTE = "\""

const MSG_ERROR_1 = "Sender is not set."
const MSG_ERROR_2 = "Recipient is not set."
const MSG_ERROR_3 = "Response Status is bad."

const DELIMITER_RESPONSE = "---"
const VERBOSE_OFF = false
const VERBOSE_ON = true

//-----------------------------------------------------------------------------|

func main() {

	var err error
	var msg string
	var ok bool
	var recipient string
	var verbose bool

	msg = "Test."
	recipient = "Vasya"
	verbose = true
	ok, err = TelegramMessageSendFromBot(msg, recipient, verbose)
	if err != nil {
		log.Println(err)
	}
	if ok == false {
		log.Println("Message was not sent.")
	}
}

//-----------------------------------------------------------------------------|

// Sends a Message from Telegram Bot into the specified Chat or Channel.
// If 'verboseMode' is set to TRUE, then it prints the sent Message and the
// Server's Response to it.
func TelegramMessageSend(
	message string,
	recipientID string,
	senderID string,
	verboseMode bool) (bool, error) {

	var recipientIDChecked string
	var clientConnectionAddress string
	var clientRequestURI string
	var err error
	var errMsg string
	var httpClient *http.Client
	var httpRequest *http.Request
	var httpRequestBody string
	var httpResponse *http.Response
	var httpResponseBodyBA []byte
	var httpResponseObj map[string]interface{}
	var httpResponseObj_ok interface{}
	var httpTransport *http.Transport
	var messageChecked string
	var tlsConfig *tls.Config

	// Check Length of Input Parameters.
	if len(senderID) == 0 {
		err = errors.New(MSG_ERROR_1)
		return false, err
	}
	if len(recipientID) == 0 {
		err = errors.New(MSG_ERROR_2)
		return false, err
	}

	clientConnectionAddress = TELEGRAM_API_ADDRESS

	clientRequestURI = TELEGRAM_API_PROTOCOL + "://" +
		clientConnectionAddress +
		TELEGRAM_API_PATH +
		senderID + "/" +
		TELEGRAM_API_FUNC_MESSAGE_SEND

	// TLS Configuration.
	tlsConfig = &tls.Config{
		ServerName:         TELEGRAM_API_HOST,
		InsecureSkipVerify: false,
	}

	// HTTP Transport.
	httpTransport = &http.Transport{
		DisableCompression:    true,
		ExpectContinueTimeout: TELEGRAM_TIMEOUT_EXPECT_CONTINUE * time.Second,
		IdleConnTimeout:       TELEGRAM_TIMEOUT_CONNECTION_IDLE * time.Second,
		MaxIdleConns:          TELEGRAM_CONNECTIONS_IDLE_MAX,
		ResponseHeaderTimeout: TELEGRAM_TIMEOUT_RESPONSE_HEADER * time.Second,
		TLSClientConfig:       tlsConfig,
	}

	// HTTP Client.
	httpClient = &http.Client{
		Transport: httpTransport,
		Timeout:   TELEGRAM_TIMEOUT_HTTP_GENERAL * time.Second,
	}

	// Process Input Data.
	messageChecked = strings.Replace(message,
		SYMBOL_BACKSLASH, SYMBOL_BACKSLASH_X2, -1)

	messageChecked = strings.Replace(messageChecked,
		SYMBOL_DOUBLEQUOTE, SYMBOL_BACKSLASH_DOUBLEQUOTE, -1)

	recipientIDChecked = strings.Replace(recipientID,
		SYMBOL_BACKSLASH, SYMBOL_BACKSLASH_X2, -1)

	recipientIDChecked = strings.Replace(recipientIDChecked,
		SYMBOL_DOUBLEQUOTE, SYMBOL_BACKSLASH_DOUBLEQUOTE, -1)

	// Prepare Request.
	httpRequestBody = "{" +
		"\"chat_id\":\"" + recipientIDChecked + "\"," +
		"\"text\":\"" + messageChecked + "\"" +
		"}"

	if verboseMode {
		log.Println(httpRequestBody)
	}

	httpRequest, err = http.NewRequest(TELEGRAM_API_PROTOCOL_METHOD,
		clientRequestURI, bytes.NewBufferString(httpRequestBody))

	httpRequest.Proto = TELEGRAM_API_PROTOCOL_VERSION_FULL
	httpRequest.Header.Add("User-Agent", TELEGRAM_API_REQUEST_USERAGENT)
	httpRequest.Header.Add("Content-Type", "application/json")

	// Send Request & get Response.
	httpResponse, err = httpClient.Do(httpRequest)
	if err != nil {
		return false, err
	}

	httpResponseBodyBA, err = ioutil.ReadAll(httpResponse.Body)
	if err != nil {

		// Clean up.
		httpResponse.Body.Close()

		return false, err
	}

	if verboseMode {
		log.Println(string(httpResponseBodyBA))
	}

	// Response -> JSON -> Object.
	err = json.Unmarshal(httpResponseBodyBA, &httpResponseObj)
	if err != nil {

		// Clean up.
		httpResponse.Body.Close()

		return false, err
	}

	// Check Status of Success.
	httpResponseObj_ok = httpResponseObj["ok"]
	if httpResponseObj_ok != true {

		// Clean up.
		httpResponse.Body.Close()

		errMsg = MSG_ERROR_3 + "\r\n" +
			DELIMITER_RESPONSE + "\r\n" +
			string(httpResponseBodyBA) + "\r\n" +
			DELIMITER_RESPONSE

		err = errors.New(errMsg)
		return false, err
	}

	// Full Clean up.
	err = httpResponse.Body.Close()
	if err != nil {
		return false, err
	}

	return true, nil
}

//-----------------------------------------------------------------------------|

func TelegramMessageSendFromBot(
	message string,
	recipient string,
	verboseMode bool) (bool, error) {

	var err error
	var sendResult bool

	sendResult, err = TelegramMessageSend(
		message,
		recipient,
		TELEGRAM_BOT_HT_TOKEN,
		verboseMode)

	if err != nil {
		return false, err
	}

	return sendResult, nil
}

//-----------------------------------------------------------------------------|
