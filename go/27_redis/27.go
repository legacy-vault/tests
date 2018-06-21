// 27.go.

package main

import (
	"github.com/gomodule/redigo/redis"
)

import (
	"log"
	"time"
	"fmt"
)

func main() {

	var conn redis.Conn
	var err error

	// Connection Options in "Redigo" are ugly... =)
	var opt1 redis.DialOption
	var opt2 redis.DialOption
	var opt3 redis.DialOption
	var opt4 redis.DialOption

	var reply interface{}
	var replyInt int64

	// Configure Connection.
	opt1 = redis.DialConnectTimeout(time.Second * 60)
	opt2 = redis.DialKeepAlive(time.Second * 60)
	opt3 = redis.DialReadTimeout(time.Second * 60)
	opt4 = redis.DialWriteTimeout(time.Second * 60)

	// Connect.
	conn, err = redis.Dial("tcp", "localhost:6379",
		opt1, opt2, opt3, opt4)
	if err != nil {
		log.Fatal(err)
	}

	// 'SET x 0'.
	reply, err = conn.Do("SET", "x", 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("reply=%+v.\r\n", reply)

	// 'INCR x'.
	reply, err = conn.Do("INCR", "x")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("reply=%+v.\r\n", reply)

	// 'INCRBY x 10'.
	reply, err = conn.Do("INCRBY", "x", 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("reply=%+v.\r\n", reply)

	// 'GET x'.
	replyInt, err = redis.Int64(conn.Do("GET", "x"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("replyInt=%+v.\r\n", replyInt)

	// 'DEL x' => 1.
	reply, err = conn.Do("DEL", "x")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("reply=%+v.\r\n", reply)

	// 'DEL x' => 0.
	reply, err = conn.Do("DEL", "x")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("reply=%+v.\r\n", reply)

	// Disconnect.
	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}
