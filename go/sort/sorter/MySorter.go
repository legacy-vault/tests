package main

// MySorter Type implements a custom Sorting of an Array of Integers.
// Actually, it is the standard Sorting, wrapped into the 'sort.Interface'
// Interface just for Demonstration Purposes.
type MySorter []int // Implements 'sort.Interface' Interface.

func (ms MySorter) Len() int {
	return len(ms)
}

func (ms MySorter) Less(i, j int) bool {
	return (ms[i] < ms[j])
}

func (ms MySorter) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}
