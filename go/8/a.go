// a.go

package main

import "fmt"

type Animal struct {
	Name string
	Size uint8
}

type Cat struct {
	Animal
	HairLength uint8
	Name       string
}

var MyCat Cat

func main() {

	MyCat.Name = "Pussy Cat"
	MyCat.Animal.Name = "Cat"
	MyCat.HairLength = 10

	fmt.Println("MyCat.Name:", MyCat.Name)               //
	fmt.Println("MyCat.Animal.Name:", MyCat.Animal.Name) //

	return
}
