// pointer_stupidity.go.

// In 'Go' Programming Language,
// Pointers are passed into Function as Values, not as real Pointers,
// as it may first seem by Analogy with such Languages as C, C++ and Java.
//
// To implement the same Functionality as in common Languages, we must make
// a small Trick. Two tricky Lines are marked with the '!' in Comments.

package main

import "fmt"

type ClassX struct {
	Name string
	Age  int
}

func main() {
	var obj *ClassX
	obj = new(ClassX) //!
	initApp(obj)
	fmt.Println(obj)
	return
}

func initApp(x *ClassX) {
	tmp := NewClassXObject()
	*x = *tmp //!
	fmt.Println("tmp:", tmp)
	fmt.Println("x:", x)
}

func NewClassXObject() *ClassX {
	x := new(ClassX)
	x.init()
	return x
}

func (o *ClassX) init() {
	o.Age = 123
	o.Name = "John"
}
