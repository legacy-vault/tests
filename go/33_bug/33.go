// 33.go.

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}
type GroupOfPeople struct {
	Name   string
	People []*Person
}
type Union struct {
	ID    int
	Group *GroupOfPeople
}
type Record struct {
	UnionField *Union
	Type       int
}

func main() {

	var aPerson *Person
	var people []*Person
	var aGroup *GroupOfPeople
	var aUnion *Union
	var aRecord *Record

	aPerson = &Person{
		Name: "John",
		Age:  10,
	}
	people = []*Person{aPerson}
	aGroup = &GroupOfPeople{
		Name:   "A Group",
		People: people,
	}
	aUnion = &Union{
		ID:    123,
		Group: aGroup,
	} // John is inside 'aUnion'.
	aRecord = &Record{
		UnionField: aUnion,
		Type:       666,
	}
	// John is NOT inside 'aRecord'.
	// WHY ?!
	fmt.Printf("aUnion: %+v.\r\n", aUnion.Group.People[0])
	fmt.Printf("aRecord: %+v.\r\n", aRecord.UnionField.Group.People[0])
}
