// append.go

package main

import (
	"fmt"
)

func main() {

	var a, b []int

	a = make([]int, 0, 2)
	fmt.Println(a)
	a = append(a, 1)
	fmt.Println(a)
	a = append(a, 2)
	fmt.Println(a)
	a = append(a, 3)
	fmt.Println(a)

	b = make([]int, 0, 2)
	fmt.Println(b)
	f1(&b, 1)
	fmt.Println(b)
	f1(&b, 2)
	fmt.Println(b)
}

func f1(arr *[]int, x int) {

	*arr = append(*arr, x)
}
