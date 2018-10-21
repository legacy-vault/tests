// sort.go

/*

	Comparison of different Ways to sort an Array of "int".

	Methods compared:

		- built-in Golang's "sort.Ints" Function;
		- modified "Quick Sort" where Pivot is at right Edge;
		- standard "Quick Sort" where Pivot is in the Middle of Array.

	Results:
	Apart from being unable to compare unsigned Integers, the built-in Method
	is the slowest one when the Data in Array is random.

*/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

//-----------------------------------------------------------------------------|

func main() {

	var i int
	var arr_len int
	var arr []int        // Original Array, not sorted.
	var a1, a2, a3 []int // Array to be sorted.
	var time_start time.Time
	var time_end time.Time
	var time1, time2, time3 time.Duration
	var copied_count int

	arr_len = 1000 * 1000 * 10 // 10 Millions.

	// Create Array of random int.
	arr = make([]int, arr_len)
	for i = 0; i < arr_len; i++ {
		arr[i] = rand.Int()
	}

	// Prepare Arrays.
	a1 = make([]int, arr_len)
	copied_count = copy(a1, arr)
	if copied_count != arr_len {
		fmt.Println("Error.")
		os.Exit(1)
	}
	a2 = make([]int, arr_len)
	copied_count = copy(a2, arr)
	if copied_count != arr_len {
		fmt.Println("Error.")
		os.Exit(1)
	}
	a3 = make([]int, arr_len)
	copied_count = copy(a3, arr)
	if copied_count != arr_len {
		fmt.Println("Error.")
		os.Exit(1)
	}

	// Test #1.
	time_start = time.Now()
	//
	sort.Ints(a1)
	//
	time_end = time.Now()
	time1 = time_end.Sub(time_start)

	// Test #2.
	time_start = time.Now()
	//
	qsort(a2)
	//
	time_end = time.Now()
	time2 = time_end.Sub(time_start)

	// Test #3.
	time_start = time.Now()
	//
	qsort_std(a2, 0, arr_len-1)
	//
	time_end = time.Now()
	time3 = time_end.Sub(time_start)

	// Report.
	fmt.Println("Duration")
	fmt.Println("Test #1:", time1)
	fmt.Println("Test #2:", time2)
	fmt.Println("Test #3:", time3)
}

//-----------------------------------------------------------------------------|

func qsort_std(a []int, lo, hi int) {

	if len(a) < 2 {
		return
	}

	if lo < hi {

		var p int // Index

		p = qsort_partition(a, lo, hi)
		qsort_std(a, lo, p-1)
		qsort_std(a, p, hi)
	}
}

//-----------------------------------------------------------------------------|

func qsort_partition(a []int, lo, hi int) int {

	//fmt.Println("lo=", lo, "hi=", hi) ///

	var i, j, p int
	var p_val int

	p = (lo + hi) / 2
	p_val = a[p]

	//fmt.Println("i=", i, "p=", p, "j=", j) ///
	//fmt.Println("p=", p, "pv=", p_val) ///

	i = lo
	j = hi

	for i <= j {

		for a[i] < p_val {
			i = i + 1
		}
		for a[j] > p_val {
			j = j - 1
		}
		if i <= j {
			a[i], a[j] = a[j], a[i]
			i = i + 1
			j = j - 1
		}
	}

	return i
}

//-----------------------------------------------------------------------------|

func qsort(a []int) []int {

	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// Pick a pivot
	pivotIndex := rand.Int() % len(a)

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	qsort(a[:left])
	qsort(a[left+1:])

	return a
}

//-----------------------------------------------------------------------------|
