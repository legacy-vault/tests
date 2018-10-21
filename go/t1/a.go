package main

import "fmt"

//import "time"

var jobs int
var ch chan int

func main() {

	ch = make(chan int)

	jobs++
	go x()

	jobs++
	go x()

	jobs++
	go x()

	//time.Sleep(time.Second * 5)
	for jobs > 0 {
		<-ch
		jobs--
	}
}

func x() {

	fmt.Println("Hello")
	ch <- 1
}
