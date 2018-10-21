// main.go.

// CRC Guess Test.

// Date: 2018-09-22.

package main

import (
	"hash/crc32"
	"log"
)

// Program's Entry Point.
func main() {

	var ba []byte
	var baElementsCount int
	var combinationsFoundCount uint64
	var crcTable *crc32.Table
	var crcSum uint32
	var crcSumSearched uint32
	var goOn bool
	var i int

	log.Println("Start.")

	baElementsCount = 5
	crcSumSearched = 16521337 //!

	crcTable = crc32.IEEETable
	ba = make([]byte, baElementsCount)
	for i = 0; i < baElementsCount; i++ {
		ba[i] = 0
	}
	goOn = true

	for goOn {

		// Increase Array Value.
		goOn = increaseBA(ba, baElementsCount-1)
		crcSum = crc32.Checksum(ba, crcTable)
		if crcSum == crcSumSearched {
			combinationsFoundCount++
			log.Println(combinationsFoundCount, ". BA:", ba) //!
		}
	}

	log.Println("End.")
}

func increaseBA(ba []byte, elementIdx int) bool {

	if ba[elementIdx] != 255 {

		// No Digit Overflow.
		ba[elementIdx]++
		return true

	} else {

		// Digit Overflow.
		if elementIdx > 0 {
			ba[elementIdx] = 0
			return increaseBA(ba, elementIdx-1)
		} else {
			return false
		}
	}
}
