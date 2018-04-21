// acestreamdirector/config.go

package config

//------------------------------------------------------------------------------

import "net"

//------------------------------------------------------------------------------

const API_VERSION_MAJOR = "3"

const CONNECTION_TYPE = "tcp"

const PRODUCT_KEY_PUBLIC = "kjYX790gTytRaXV04IvC-xZH3A18sj5b1Tf3I-J5XVS1xsj-j0797KwxxLpBl26HPvWMm"

const SERVER_HOST_DEFAULT = "localhost"
const SERVER_PORT_DEFAULT = 62062

const VERBOSE_DEFAULT = true

const INT_SIZE = 32 // Bytes in int.

// API : http://wiki.acestream.org/wiki/index.php/Engine_API

//------------------------------------------------------------------------------

var ProductKey string
var ConnAddress net.TCPAddr

var ServerPortHTTP int
var ServerKey string

var Verbose bool

//------------------------------------------------------------------------------

// Initializes Program's Configuration.
func Init() {

	Verbose = VERBOSE_DEFAULT

	ConnAddress = net.TCPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: SERVER_PORT_DEFAULT,
	}

	ServerPortHTTP = 0

	ProductKey = PRODUCT_KEY_PUBLIC
}

//------------------------------------------------------------------------------
