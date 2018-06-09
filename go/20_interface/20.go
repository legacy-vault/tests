// 20.go.

// How to check whether an Object implements an Interface or not?

package main

import "fmt"

type T struct {
	Size int
}

type I1 interface {
	GetSize() int
}

type I2 interface {
	GetWeight() int
}

func (obj T) GetSize() int {

	var size int

	size = obj.Size

	return size
}

func main() {

	var implements bool
	var obj T
	var _ I1 = (*T)(nil) // Compile-Time Check.
	//var _ I2 = (*T)(nil) // Compile-Time Check. This fails to compile.

	_, implements = interface{}(obj).(I1)
	fmt.Println(implements)

	_, implements = interface{}(obj).(I2)
	fmt.Println(implements)
}
