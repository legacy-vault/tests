// a.go

package main

import "fmt"
import "unsafe"

type Struct1 struct {
	a uint8   // 1 Byte.
	b [5]byte // 5 Bytes.
	c [3]byte // 3 Bytes.
	d uint32  // 4 Bytes.
	e uint8   // 1 Byte.
	f uint16  // 2 Bytes.
}

func main() {

	var addr unsafe.Pointer
	var alignment uint8
	var offset uint8
	var s1 Struct1
	var s1BytesReservedInMemory uintptr
	var s1BytesActuallyUsed uintptr
	var size uintptr
	var varName string

	s1.a = 1
	s1.b = [5]byte{2, 2, 2, 2, 2}
	s1.c = [3]byte{3, 3, 3}
	s1.d = 4
	s1.e = 5
	s1.f = 6

	s1BytesReservedInMemory = unsafe.Sizeof(s1)
	s1BytesActuallyUsed = 0

	fmt.Println("VarName\tAddress\t\tOffset\tSize\tAlignment")

	// s1.a.
	varName = "s1.a"
	addr = unsafe.Pointer(&(s1.a))
	alignment = uint8(unsafe.Alignof(s1.a))
	offset = uint8(unsafe.Offsetof(s1.a))
	size = unsafe.Sizeof(s1.a)
	s1BytesActuallyUsed += size
	fmt.Printf("%v\t%v\t%v\t%v\t%v\r\n", varName, addr, offset, size, alignment)

	// s1.b.
	varName = "s1.b"
	addr = unsafe.Pointer(&(s1.b))
	alignment = uint8(unsafe.Alignof(s1.b))
	offset = uint8(unsafe.Offsetof(s1.b))
	size = unsafe.Sizeof(s1.b)
	s1BytesActuallyUsed += size
	fmt.Printf("%v\t%v\t%v\t%v\t%v\r\n", varName, addr, offset, size, alignment)

	// s1.c.
	varName = "s1.c"
	addr = unsafe.Pointer(&(s1.c))
	alignment = uint8(unsafe.Alignof(s1.c))
	offset = uint8(unsafe.Offsetof(s1.c))
	size = unsafe.Sizeof(s1.c)
	s1BytesActuallyUsed += size
	fmt.Printf("%v\t%v\t%v\t%v\t%v\r\n", varName, addr, offset, size, alignment)

	// s1.d.
	varName = "s1.d"
	addr = unsafe.Pointer(&(s1.d))
	alignment = uint8(unsafe.Alignof(s1.d))
	offset = uint8(unsafe.Offsetof(s1.d))
	size = unsafe.Sizeof(s1.d)
	s1BytesActuallyUsed += size
	fmt.Printf("%v\t%v\t%v\t%v\t%v\r\n", varName, addr, offset, size, alignment)

	// s1.e.
	varName = "s1.e"
	addr = unsafe.Pointer(&(s1.e))
	alignment = uint8(unsafe.Alignof(s1.e))
	offset = uint8(unsafe.Offsetof(s1.e))
	size = unsafe.Sizeof(s1.e)
	s1BytesActuallyUsed += size
	fmt.Printf("%v\t%v\t%v\t%v\t%v\r\n", varName, addr, offset, size, alignment)

	// s1.f.
	varName = "s1.f"
	addr = unsafe.Pointer(&(s1.f))
	alignment = uint8(unsafe.Alignof(s1.f))
	offset = uint8(unsafe.Offsetof(s1.f))
	size = unsafe.Sizeof(s1.f)
	s1BytesActuallyUsed += size
	fmt.Printf("%v\t%v\t%v\t%v\t%v\r\n", varName, addr, offset, size, alignment)

	fmt.Println()
	fmt.Println("Size of Struct in Memory:", s1BytesReservedInMemory)
	fmt.Println("Size of useful Data in Memory:", s1BytesActuallyUsed)
}
