// main.go.

package main

import "fmt"

func main() {

	a := 40
	b := 2
	c := add(a, b)
	d := add(a, c)
	e := add(d, -c)

	fmt.Println(e)
}
