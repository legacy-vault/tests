// a.go

package main

import "fmt"

func main() {

	var a int = 10
	fmt.Println("a-1:", a)

	for true {

		var a int

		a++
		fmt.Println("a-2:", a)

		break
	}

	fmt.Println("a-3:", a)

	{
		var a int
		a = 3
		fmt.Println("a-4:", a)
	}

	fmt.Println("a-5:", a)
}
