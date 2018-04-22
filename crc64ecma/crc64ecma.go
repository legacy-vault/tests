// crc64ecma.go

/*
	CRC 64 ECMA.

	64-Bit Cyclic Redundancy Check using the Polynomial Table provided by the
	European Computer Manufacturers Association (ECMA).

	Notes:

		1.	This Program can read only small Files which can be placed in the
			Memory.

	Version:	0.1.
	Date:		2018-04-22.
	Author:		McArcher.

*/

//=============================================================================|

package main

//=============================================================================|

import "hash/crc64"
import "io/ioutil"
import "fmt"
import "log"
import "os"
import "strconv"

//=============================================================================|

const MSG_1 = "Usage: crc64ecma <file_1> <file_2> ..."

const MSG_ERR_1 = "Checksum Calculation Error. File:"

//=============================================================================|

// Calculates CRC-64 Check Sum and returns it as a Hexademical String.
func CalculateSumStringHex(filePath string) (string, error) {

	var crcSum uint64
	var crcTable *crc64.Table
	var fileData []byte
	var err error
	var result string

	result = ""

	// Read File.
	fileData, err = ioutil.ReadFile(filePath)
	if err != nil {
		return result, err
	}

	// Prepare Polynomial Table.
	crcTable = crc64.MakeTable(crc64.ECMA)

	// Calculate Check Sum.
	crcSum = crc64.Checksum(fileData, crcTable)

	// Uint64 -> String.
	result = strconv.FormatUint(crcSum, 16)

	return result, nil
}

//=============================================================================|

// Program's Entry Point.
func main() {

	var arg string
	var clArgumentsCount int
	var err error
	var i int
	var num int
	var sum string

	clArgumentsCount = len(os.Args)
	if clArgumentsCount <= 1 {
		fmt.Println(MSG_1)
		return
	}

	num = clArgumentsCount - 1
	for i = 1; i <= num; i++ {

		// Get Argument.
		arg = os.Args[i]

		// Calculate Check Sum.
		sum, err = CalculateSumStringHex(arg)
		if err != nil {
			log.Println(MSG_ERR_1, arg, err)
			continue
		}

		fmt.Println(sum, arg)
	}

}

//=============================================================================|
