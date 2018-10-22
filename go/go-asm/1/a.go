// a.go.

package main

import "fmt"

func main() {

	x := 40
	y := 2
	z := add(x, y)

	fmt.Println(z)
}

func add(x int, y int) int {
	return x+y
}
