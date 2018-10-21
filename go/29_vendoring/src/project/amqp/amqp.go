// amqp.go.

package amqp

import (
	"github.com/streadway/amqp"
)

// Connects to AMQP Server.
func MyConnect(address string) (*amqp.Connection, error) {

	var conn *amqp.Connection
	var err error

	conn, err = amqp.Dial(address)

	return conn, err
}

// DisConnects from AMQP Server.
func MyDisConnect(conn *amqp.Connection) (error) {

	var err error

	err = conn.Close()

	return err
}

// Tests Compiler choosing a Function from local Package.
// We have declared this Package's name as 'amqp'.
// At the same Time, we import an external Package with the same 'amqp' Name.
// This is a bad Programming Style, but unfortunately it is often used nowadays.
// This Function tests that Compiler can differentiate local 'amqp' Package from
// the external 'amqp' Package.
func CompilerTest(address string) (*amqp.Connection, error) {

	var conn *amqp.Connection
	var err error

	conn, err = MyConnect(address)

	return conn, err
}
