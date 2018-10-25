// pointer_test_a_test.go.

// Here we are dealing with Golang's Pointers Handling.
// The main Aim is to feel the Difference between various Ways of Object
// Initialization in Go Language.

package main

import (
	"testing"
)

func BenchmarkTest1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test1()
	}
}

func BenchmarkTest2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test2()
	}
}
