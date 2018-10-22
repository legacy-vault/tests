// b.go.

package main

import "fmt"

type ObjectX struct {
	FieldY int
}

func main() {

	var a int
	var b int
	var c int
	var d int
	var obj ObjectX

	obj.FieldY = 123
	a = 40
	b = 2
	c = add(a, b)
	d = mul(c, obj)

	fmt.Println(d)
}

func add(x int, y int) int {

	var sum int

	sum = x + y

	return sum
}

func mul(x int, obj ObjectX) int {

	var result int

	result = x * obj.FieldY

	return result
}
