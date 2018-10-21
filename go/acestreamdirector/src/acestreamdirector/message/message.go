// acestreamdirector/message.go

package message

//------------------------------------------------------------------------------

import "errors"
import "strings"

//------------------------------------------------------------------------------

const ASCII_CR = '\r' // DEC:13, HEX:0D.
const ASCII_LF = '\n' // DEC:10, HEX:0A.

const TYPE_1 = "HELLOBG"
const TYPE_2 = "HELLOTS"
const TYPE_3 = "READY"

const PARAM_BMODE = "bmode"
const PARAM_HTTP_PORT = "http_port"
const PARAM_KEY = "key"
const PARAM_VERSION = "version"
const PARAM_VERSION_CODE = "version_code"

const END = "\r\n"
const LAST_BYTE = ASCII_LF
const PRELAST_BYTE = ASCII_CR

const FIELDS_SEPARATOR = " "
const KEY_VALUE_SEPARATOR = "="
const VERSION_SEPARATOR = "."
const X_SEPARATOR = "-"

const LENGTH_MIN = 2

//------------------------------------------------------------------------------

type Parameter struct {
	Key   string
	Value string
}

type Message struct {
	Type       string
	Parameters []Parameter
}

//------------------------------------------------------------------------------

// Checks whether Message has been fully sent.
func CheckLength(message []byte, bytesSent int) error {

	var err error

	if len(message) != bytesSent {
		err = errors.New("Message has been sent partially!")
		return err
	}

	return nil
}

//------------------------------------------------------------------------------

// Parses raw Reply into Message Structure.
func Parse(str *string) (*Message, error) {

	var err error
	var i int
	var fieldsCount int
	var kv []string
	var key string
	var msg Message
	var msgType string
	var param Parameter
	var params []Parameter
	var paramsCount int
	var part string
	var parts []string
	var value string

	parts = strings.Split(*str, FIELDS_SEPARATOR)

	fieldsCount = len(parts)
	if fieldsCount < 1 {
		err = errors.New("No fields in a message!")
		return nil, err
	}

	for i = range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	msgType = parts[0]

	paramsCount = fieldsCount - 1
	params = make([]Parameter, paramsCount)

	for i = 1; i <= paramsCount; i++ {

		part = parts[i]

		kv = strings.Split(part, KEY_VALUE_SEPARATOR)
		if len(kv) != 2 {
			err = errors.New("Parse error in field [" + part + "].")
			return nil, err
		}

		key = kv[0]
		value = kv[1]

		param = Parameter{
			Key:   key,
			Value: value,
		}

		params[i-1] = param
	}

	msg = Message{
		Type:       msgType,
		Parameters: params,
	}

	return &msg, nil
}

//------------------------------------------------------------------------------

// Reads major Version from the full Version String.
func VersionMajor(ver string) string {

	var parts []string
	var result string

	parts = strings.Split(ver, VERSION_SEPARATOR)

	if len(parts) < 1 {
		return ""
	}

	result = parts[0]

	return result
}

//------------------------------------------------------------------------------
