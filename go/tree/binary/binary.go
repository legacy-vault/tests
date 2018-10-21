// binary.go

package main

import (
	"fmt"
)

type treeNode struct {
	originalIndex   uint64 // Index of the Object in the DataBase.
	content         string // Text of an Object.
	parentNode      *treeNode
	leftNode        *treeNode
	rightNode       *treeNode
	duplicatesCount uint64 // Number of Objects with the same Text.
}

type tree struct {
	root       *treeNode
	nodesCount uint64
}

func main() {

	fmt.Println("") //
}
