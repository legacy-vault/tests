// endian.go

/*
	Functions to detect 'Endianness' of the C.P.U.

	We cannot truly detect the Endianness while we have no Methods to get the
	real Bits Order inside a Byte. Most of the modern Hardware provides no
	C.P.U. Instructions to know the real Bits Order. All we can do here and
	now: we can detect the Order of Bytes inside multi-Byte Integer Variables.
*/

package main

import "fmt"
import "unsafe"

const BYTE_ORDER_BIG_ENDIAN = 1
const BYTE_ORDER_LITTLE_ENDIAN = 2
const BYTE_ORDER_UNKNOWN_ENDIAN = 4

func main() {

	var endian uint8
	var verbose bool

	verbose = false
	endian = ByteOrder(verbose)
	fmt.Println("endian:", endian)
}

// Detects Byte Order in 32-bit Unsigned Integer.
func ByteOrder(verbose bool) uint8 {

	const BYTE_BE_0000_0001 uint8 = 1
	const BYTE_BE_0000_0011 uint8 = 3
	const BYTE_BE_0000_0111 uint8 = 7
	const BYTE_BE_0000_1111 uint8 = 15

	var a uint32
	var a1 uint8
	var a2 uint8
	var a3 uint8
	var a4 uint8

	var ptrA unsafe.Pointer
	var ptrA1 unsafe.Pointer
	var ptrA2 unsafe.Pointer
	var ptrA3 unsafe.Pointer
	var ptrA4 unsafe.Pointer

	var b uint8

	a = uint32(BYTE_BE_0000_0001) +
		uint32(BYTE_BE_0000_0011)*256 +
		uint32(BYTE_BE_0000_0111)*65536 +
		uint32(BYTE_BE_0000_1111)*16777216

	ptrA = unsafe.Pointer(&a)
	ptrA1 = unsafe.Pointer(uintptr(ptrA) + 0*unsafe.Sizeof(b))
	ptrA2 = unsafe.Pointer(uintptr(ptrA) + 1*unsafe.Sizeof(b))
	ptrA3 = unsafe.Pointer(uintptr(ptrA) + 2*unsafe.Sizeof(b))
	ptrA4 = unsafe.Pointer(uintptr(ptrA) + 3*unsafe.Sizeof(b))

	a1 = *((*uint8)(ptrA1))
	a2 = *((*uint8)(ptrA2))
	a3 = *((*uint8)(ptrA3))
	a4 = *((*uint8)(ptrA4))

	if (a1 == BYTE_BE_0000_0001) &&
		(a2 == BYTE_BE_0000_0011) &&
		(a3 == BYTE_BE_0000_0111) &&
		(a4 == BYTE_BE_0000_1111) {

		// Byte Order is 'Little Endian'.
		return BYTE_ORDER_LITTLE_ENDIAN

	} else if (a1 == BYTE_BE_0000_1111) &&
		(a2 == BYTE_BE_0000_0111) &&
		(a3 == BYTE_BE_0000_0011) &&
		(a4 == BYTE_BE_0000_0001) {

		// Byte Order is 'Big Endian'.
		return BYTE_ORDER_BIG_ENDIAN

	} else {

		if verbose {

			// Report.
			fmt.Println("a:", a)
			fmt.Println("a1:", a1)
			fmt.Println("a2:", a2)
			fmt.Println("a3:", a3)
			fmt.Println("a4:", a4)
		}

		// Byte Order is UnKnown.
		return BYTE_ORDER_UNKNOWN_ENDIAN
	}
}
