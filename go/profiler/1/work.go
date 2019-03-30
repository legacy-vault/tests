// work.go.

package main

import "log"

// Work.

func work() {

	var i int
	var x int

	for i = 1; i <= 10; i++ {
		x = square(i)
		log.Printf(
			"i=%v x=%v",
			i,
			x,
		)
	}
}

func square(
	x int,
) int {

	return x * x
}
