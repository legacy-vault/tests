// crc.go

package main

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"math"
	"strconv"
)

var ba []byte
var hash uint32
var h uint32
var hash_hex string
var a uint32
var a_max uint32
var aa []byte
var collisions [][]byte

func main() {

	fmt.Println("Hello.")

	ba = []byte("Test") // 4 Bytes = 32 Bit.

	hash = crc32.ChecksumIEEE(ba)
	hash_hex = strconv.FormatUint(uint64(hash), 16)

	fmt.Println(hash_hex)

	// Find Arrays with the same Hash Sum and Length.
	a_max = math.MaxUint32
	aa = make([]byte, 4)
	fmt.Println(a_max)
	for a = 0; a <= a_max; a++ {

		binary.BigEndian.PutUint32(aa, a)
		fmt.Println(aa, len(collisions))
		h = crc32.ChecksumIEEE(aa)

		if h == hash {
			collisions = append(collisions, aa)
			fmt.Println("New Collision Found.", aa)

		}
	}
	fmt.Println("All Collisions:", collisions)
	fmt.Println("Number of Collisions:", len(collisions))
}
