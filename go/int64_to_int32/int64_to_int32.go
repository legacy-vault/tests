// int64_to_int32

package main

import (
	"fmt"
	"math"
)

func main() {

	var a64 int64
	var a32 int32

	fmt.Println("Hello.")

	a64 = math.MaxInt64
	if a64 > math.MaxInt32 {
		fmt.Println("too big")
	} else {
		a32 = (int32)(a64)
	}

	fmt.Println(a64, a32)
}
