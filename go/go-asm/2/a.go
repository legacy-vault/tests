// a.go.

package main

import "fmt"

type ObjectX struct {
	FieldY int
}

func main() {

	obj := ObjectX{FieldY: 123}
	a := 40
	b := 2
	c := add(a, b)
	d := mul(c, obj)

	fmt.Println(d)
}

func add(x int, y int) int {
	return x + y
}

func mul(x int, obj ObjectX) int {
	return x * obj.FieldY
}
