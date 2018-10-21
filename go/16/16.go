// 16.go.

package main

import "fmt"

func main() {

	var ch chan int
	var i int
	var n int
	var rcvd bool

	ch = make(chan int, 100)
	ch <- 99
	i = <-ch
	fmt.Printf("i=[%v].\r\n", i)

	// No Data in Channel selects 'default' Branch.
	ch <- 1
	ch <- 2
	//ch <- 3

	for i = 1; i <= 3; i++ {

		select {
		case n, rcvd = <-ch:
			if rcvd {
				fmt.Printf("rcvd=[%v] n=[%v].\r\n", rcvd, n)
			} else {
				fmt.Printf("rcvd=[%v].\r\n", rcvd)
			}

		default:
			fmt.Printf("no data.\r\n")
		}
	}

	// Closed Channel selects first Branch.
	ch <- 1
	ch <- 2
	//ch <- 3
	close(ch)

	for i = 1; i <= 3; i++ {

		select {
		case n, rcvd = <-ch:
			if rcvd {
				fmt.Printf("rcvd=[%v] n=[%v].\r\n", rcvd, n)
			} else {
				fmt.Printf("rcvd=[%v].\r\n", rcvd)
			}

		default:
			fmt.Printf("no data.\r\n")
		}
	}
}
