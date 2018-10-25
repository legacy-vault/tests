// pointer_test_b.go.

// Here we are dealing with Golang's Pointers Handling.
// The main Aim is to feel the Difference between various Ways of Object
// Initialization in Go Language.

package main

import (
	"fmt"
	"time"
)

type ClassX struct {
	Name             string
	Age              int
	Hobbies          []string
	Date             time.Time
	NeighbourObjects []*ClassX
	Size             uint64
}

func main() {

	test1()
	fmt.Println()
	test2()

	return
}

// Variant I.
func test1() {
	var obj_1 *ClassX
	obj_1 = new(ClassX)
	initApp_1(obj_1)
	//fmt.Println(obj_1)
}

// Variant II.
func test2() {
	var obj_2 *ClassX
	initApp_2(&obj_2)
	//fmt.Println(obj_2)
}

func initApp_1(x *ClassX) {
	tmp := NewClassXObject()
	*x = *tmp
}

func initApp_2(x **ClassX) {
	tmp := NewClassXObject()
	*x = tmp
}

func NewClassXObject() *ClassX {
	x := new(ClassX)
	x.init()
	return x
}

func (o *ClassX) init() {
	o.Name = "John"
	o.Age = 123
	o.Hobbies = []string{
		"Singing",
		"Running",
		"Music Listening",
		"Music Composing",
		"Car Fixing",
		"Speaking with Cats",
	}
	o.Date = time.Now()
	o.NeighbourObjects = make([]*ClassX, 20)
	o.Size = 12345
}
