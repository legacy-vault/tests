// a.go

package main

import "fmt"

func main() {

	var a int
	var i int
	var j int

	for a = 6; a <= 8; a++ {

		fmt.Println(">>> a:", a)
	Label:

		for i = 1; i <= 3; i++ {

			fmt.Println("> i:", i)

			for j = i; j <= i+2; j++ {

				fmt.Println("i:", i, "j:", j)

				if (i == 2) && (j == 3) {
					break Label
				}

			}
		}
	}

}
