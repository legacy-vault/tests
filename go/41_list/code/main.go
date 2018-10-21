// main.go.

// List Test :: Main File.

package main

import (
	"container/list"
	"fmt"
)

// Program's Entry Point.
func main() {

	var aList *list.List
	var aListStr string

	aList = list.New()
	aList.PushBack("A")
	aList.PushBack("B")
	aList.PushBack("C")

	aListStr = listToSnake(*aList)
	fmt.Println(aListStr)

	return
}
