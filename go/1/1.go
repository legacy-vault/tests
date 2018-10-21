// 1.go

package main

import "fmt"

type bit bool
type array_of_bits []bit

func main() {

	var my_array array_of_bits

	my_array = make(array_of_bits, 8)

	my_array[0] = true
	my_array[7] = true

	fmt.Println(my_array)
}
