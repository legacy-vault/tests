package main

import (
	"fmt"
	"sort"
)

type Test struct {
	Array         []int
	ForwardSorter func([]int)
	ReverseSorter func([]int)
}

func main() {
	var tests []Test = prepareTests()
	runTests(tests)
}

func prepareTests() []Test {

	var tests []Test = make([]Test, 0)
	var test Test

	// Test #1. Simple Sorting.
	test = Test{
		Array:         getInitialArray(),
		ForwardSorter: sort.Ints,
	}
	tests = append(tests, test)

	// Test #2. Simple Reverse Sorting after standard Sorting.
	test = Test{
		Array: getInitialArray(),
		ForwardSorter: func(x []int) {
			sort.Sort(sort.IntSlice(x))
		},
		ReverseSorter: func(x []int) {
			sort.Sort(sort.Reverse(sort.IntSlice(x)))
		},
	}
	tests = append(tests, test)

	// Test #3. Simple Reverse Sorting without standard Sorting.
	test = Test{
		Array: getInitialArray(),
		ReverseSorter: func(x []int) {
			sort.Sort(sort.Reverse(sort.IntSlice(x)))
		},
	}
	tests = append(tests, test)

	// Test #4. Custom Sorter.
	test = Test{
		Array: getInitialArray(),
		ForwardSorter: func(x []int) {
			sort.Sort(MySorter(x))
		},
		ReverseSorter: func(x []int) {
			sort.Sort(sort.Reverse(MySorter(x)))
		},
	}
	tests = append(tests, test)

	// Test #5. Custom Sorter without forward Sorting.
	test = Test{
		Array: getInitialArray(),
		ReverseSorter: func(x []int) {
			sort.Sort(sort.Reverse(MySorter(x)))
		},
	}
	tests = append(tests, test)

	return tests
}

func getInitialArray() []int {
	var x []int = make([]int, 10)
	for i := 1; i <= 10; i++ {
		x[i-1] = i
	}
	x[1] = 99
	return x
}

func runTests(tests []Test) {
	for i, test := range tests {

		fmt.Printf("Test #%v.\r\n", i+1)

		fmt.Println("Initial Array:")
		fmt.Println(test.Array)

		if test.ForwardSorter != nil {
			test.ForwardSorter(test.Array)
			fmt.Println("Sorted Array:")
			fmt.Println(test.Array)
		}

		if test.ReverseSorter != nil {
			test.ReverseSorter(test.Array)
			fmt.Println("Array sorted reversly:")
			fmt.Println(test.Array)
		}

		fmt.Println()
	}
}
