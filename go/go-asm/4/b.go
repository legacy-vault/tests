// b.go.

package main

import "fmt"

func main() {

	var sum int

	for j := 1; j <= 5; j++ {
		for i := 1; i <= 9; i++ {
			sum = add(sum, i)
		}
		sum = add(sum, j)
	}

	fmt.Println(sum)
}

func add(a, b int) int {
	return a + b
}
