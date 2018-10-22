// a.go.

package main

import "fmt"

func main() {

	var i int
	var iMax int
	var j int
	var jMax int
	var sum int

	iMax = 9
	jMax = 5

	for j = 1; j <= jMax; j++ {
		for i = 1; i <= iMax; i++ {
			sum = add(sum, i)
		}
		sum = add(sum, j)
	}

	fmt.Println(sum)
}

func add(a, b int) int {
	return a + b
}
