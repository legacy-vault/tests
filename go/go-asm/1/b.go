// b.go.

package main

import "fmt"

func main() {

	var x int
	var y int
	var z int
	
	x = 40
	y = 2
	z = add(x, y)

	fmt.Println(z)
}

func add(x int, y int) int {
	
	var sum int
	
	sum = x + y
	
	return sum
}
