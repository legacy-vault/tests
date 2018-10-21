// testing.go

package main

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {

	fmt.Println("Test A.")

}

func TestB(t *testing.T) {

	var s int

	s = sum(1, 2)
	if s != 3 {
		t.Fail()
	}

}
