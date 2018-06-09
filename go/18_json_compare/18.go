// 18.go.

package main

import (
	jc "./json_comparator"
)

// Reference Structure.
type Point struct {
	X int
	Y int
}
type ReferenceData struct {
	ArrayA []int16
	ArrayB []Point
	FlagF  bool
	G      bool
	Person struct {
		Name     string
		Nickname string
		Age      uint8
	}
	Hardware struct {
		Vendor string
		Year   uint16
	}
}

func main() {

	var ba_1 []byte
	var ba_2 []byte
	var ba_3 []byte
	var ba_4 []byte
	var cfg jc.Config
	var err error
	var refObj_1 ReferenceData
	var refObj_2 ReferenceData
	var result int
	var str_1 string
	var str_2 string
	var str_3 string
	var str_4 string
	var verbose bool

	// Prepare Data.
	refObj_1.ArrayA = []int16{1, 2, 3}
	refObj_1.ArrayB = []Point{{X: 6, Y: 7}, {X: 8, Y: 9}}
	refObj_1.FlagF = true
	refObj_1.Person.Name = "Василий"
	refObj_1.Person.Age = 25
	refObj_1.Person.Nickname = "Вася"
	refObj_1.Hardware.Vendor = "AMD"
	refObj_1.Hardware.Year = 2012
	refObj_2 = refObj_1

	str_1 = `
{
	"ArrayA":[1,2,3],
	"ArrayB":
	[
		{"X":6,"Y":7},
		{"X":8,"Y":9}
	],
	"FlagF":true,
	"G":false,
	"Person":
		{"Name":"Василий","Age":25,"Nickname":"Вася"},
	"Hardware":
		{"Vendor":"AMD","Year":2012}
}`
	str_2 = `
{
	"ArrayA":[1,2,3],
	"ArrayB":
	[
		{"X":6,"Y":7},
		{"X":8,"Y":9}
	],
	"FlagF":null,
	"G":false,
	"Person":
		{"Name":"Василий","Age":25,"Nickname":"Вася"},
	"Hardware":
		{"Vendor":"AMD","Year":2012}
}`
	str_3 = `
{
	"ArrayA":[1,2,3],
	"ArrayB":
	[
		{"X":6,"Y":7},
		{"X":8,"Y":9}
	],
	"FlagF":null,
	"G":false,
	"Person":
		{"Name":"Василий","Age":25,"Nickname":"Вася"},
	"Hardware":
		{"Vendor":"AMD","Year":2012},
	"SomethingNew":123
}`
	str_4 = `
{
	"ArrayA":[1,2,3],
	"ArrayB":
	[
		{"X":6,"Y":7},
		{"X":8,"Y":9}
	],
	"FlagF":true,
	"G":false,
	"Person":
		{"Name":"Василий"},
	"Hardware":
		{"Vendor":"AMD","Year":2012}
}`
	ba_1 = []byte(str_1)
	ba_2 = []byte(str_2)
	ba_3 = []byte(str_3)
	ba_4 = []byte(str_4)

	// 1. Fully equal Objects.
	cfg = jc.CONFIG_SLICE_COMPAR_FULL
	verbose = false
	result, err = jc.CompareBytesWithObject(ba_1, refObj_1, cfg, verbose)
	jc.PrintResult(result, err)

	// 2. Partially equal Objects (only first Elements of Slices are checked).
	cfg = jc.CONFIG_DEFAULT
	verbose = false
	result, err = jc.CompareBytesWithObject(ba_1, refObj_1, cfg, verbose)
	jc.PrintResult(result, err)

	// 3. JSON with 'null' Value.
	cfg = jc.CONFIG_SLICE_COMPAR_FULL
	verbose = false
	result, err = jc.CompareBytesWithObject(ba_2, refObj_2, cfg, verbose)
	jc.PrintResult(result, err)

	// 4. Similar Objects, full Comparison.
	cfg = jc.CONFIG_SLICE_COMPAR_FULL
	verbose = false
	refObj_1.ArrayA[2] = 33
	result, err = jc.CompareBytesWithObject(ba_1, refObj_1, cfg, verbose)
	jc.PrintResult(result, err)
	// Revert Changes.
	refObj_1.ArrayA[2] = 3

	// 5. Similar Objects, partial Comparison of Slices.
	cfg = jc.CONFIG_DEFAULT
	verbose = false
	refObj_1.ArrayA[0] = 11
	result, err = jc.CompareBytesWithObject(ba_1, refObj_1, cfg, verbose)
	jc.PrintResult(result, err)
	// Revert Changes.
	refObj_1.ArrayA[0] = 1

	// 6. JSON Object has more Fields than Reference Object.
	cfg = jc.CONFIG_DEFAULT
	verbose = false
	result, err = jc.CompareBytesWithObject(ba_3, refObj_1, cfg, verbose)
	jc.PrintResult(result, err)

	// 7. JSON Object has less Fields than Reference Object.
	cfg = jc.CONFIG_DEFAULT
	verbose = false
	result, err = jc.CompareBytesWithObject(ba_4, refObj_1, cfg, verbose)
	jc.PrintResult(result, err)
}
