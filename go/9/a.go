// a.go

package main

import "fmt"

type TA = uint64
type TB = uint64
type TC = TA

func main() {

	var a TA
	var b TB
	var c TC
	var x uint64

	a = 1
	b = 2
	c = 3
	x = 100

	a = c

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("x:", x)
}
