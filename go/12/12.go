// 12.go

package main

import "fmt"
import "strconv"

func main() {

	var a uint8 = 3 / 2

	fmt.Println("a =", strconv.FormatUint(uint64(a), 10))
}
