// duplicate_int.go

/*

	Search for Duplicates in the large Array of int.

	Comparison of two Methods:

	1. Sorting by QuickSort and Viewing Neighbours.
	2. Writing to Map and checking Existence of an Element.

	Creator:		McArcher.
	Date:			2018-01-31.

*/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

//-----------------------------------------------------------------------------|

func main() {

	var i int
	var arr_len int
	var arr []int    // Original Array, not sorted.
	var a1, a2 []int // Arrays for Tests.
	var duplicates []int
	var duplicates_count int
	var time_start time.Time
	var time_end time.Time
	var time1, time2 time.Duration
	var copied_count int

	arr_len = 1000 * 1000 * 10 // 10 Millions.
	//arr_len = 100             ///
	rand.Seed(time.Now().UTC().UnixNano())

	// Create Array of random int.
	arr = make([]int, arr_len)
	for i = 0; i < arr_len; i++ {
		arr[i] = rand.Int() / 1000000 // Narrow the Dispersion.
	}
	// Artificial Duplicates.
	arr[1] = arr[0]
	arr[arr_len-1] = arr[2]
	//fmt.Println("Array:", arr) ///

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

	// Reset Counter.
	duplicates = make([]int, 0)

	// Test #1.
	time_start = time.Now()
	test_1(a1, &duplicates, &duplicates_count)
	time_end = time.Now()
	time1 = time_end.Sub(time_start)
	//
	fmt.Println("Test #1:", time1, duplicates_count, duplicates)

	// Reset Counter.
	duplicates = make([]int, 0)

	// Test #2.
	time_start = time.Now()
	test_2(a2, &duplicates, &duplicates_count)
	time_end = time.Now()
	time2 = time_end.Sub(time_start)
	//
	fmt.Println("Test #2:", time2, duplicates_count, duplicates)
}

//-----------------------------------------------------------------------------|

func test_1(a []int, duplicates *[]int, duplicates_count *int) {

	qsort_std(a, 0, len(a)-1)
	*duplicates_count = find_duplicates(a, duplicates)
}

//-----------------------------------------------------------------------------|

func test_2(a []int, duplicates *[]int, duplicates_count *int) {

	*duplicates_count = find_duplicates_by_map(a, duplicates)
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

	var i, j, p int
	var p_val int

	p = (lo + hi) / 2
	p_val = a[p]

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

func find_duplicates(src []int, duplicates *[]int) int {

	var k, k_max, dup_count int
	var element int

	k_max = len(src) - 1

	for k = 1; k < k_max; k++ {

		element = src[k]
		if src[k-1] == element {

			*duplicates = append(*duplicates, element)
		}
	}

	dup_count = len(*duplicates)

	return dup_count
}

//-----------------------------------------------------------------------------|

func find_duplicates_by_map(src []int, duplicates *[]int) int {

	var m map[int]bool
	var element int
	var dup_count int
	var map_val bool // => same time
	//var exists bool // => same time

	m = make(map[int]bool)

	for _, element = range src {

		/*
			// Check by Existence.
			_, exists = m[element]

			if exists {

				*duplicates = append(*duplicates, element)

			} else {

				m[element] = true
			}
		*/
		/**/
		// Check by Value.
		map_val, _ = m[element]

		if map_val {

			*duplicates = append(*duplicates, element)

		} else {

			m[element] = true
		}
		/**/
	}

	dup_count = len(*duplicates)

	return dup_count
}

//-----------------------------------------------------------------------------|
