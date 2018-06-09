// json_comparator.go.

// This is a simple Comparator of JSON Byte Array with Structure.
// It can perform a Comparison in two Modes:
// 		1. full;
// 		2. partial.
// In full Mode, Slices (Arrays) are fully compared, but
// in partial Mode, only first Elements of Slices are compared.
// Maps (Structures) are always fully compared.

package json_comparator

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"encoding/json"
)

type Config uint8

const CONFIG_DEFAULT = 0
const CONFIG_SLICE_COMPAR_FULL = 1

const RESULT_EQUAL = 0
const RESULT_SIMILAR = 1
const RESULT_UNKNOWN = 2
const ERROR_JSON_UNMARSHAL = 3
const ERROR_JSON_MARSHAL = 4
const ERROR_TYPE_MISMATCH = 5
const ERROR_MAP_TYPE = 6
const ERROR_MAP_KEYS_MISMATCH = 7
const ERROR_TYPE_CONVERTION = 8
const ERROR_TYPE_UNSUPPORTED = 9
const ERROR_REFERENCE_OBJECT = 10
const ERROR_SLICE_COMPAR_FULL_FAIL = 11
const ERROR_INTERNAL = 12

const MSG_DIFFERENT = "Objects are different."
const MSG_EQUAL = "Objects are equal."
const MSG_SIMILAR = "Objects have similar Structure, " +
	"but the Values are different."
const MSG_UNKNOWN = "JSON Object has an unknown Field ('null'), " +
	"and can not be compared. However, no Errors have been found."
const MSG_ERROR_INTERNAL = "INTERNAL_ERROR"

// Compares Boolean Variables.
func CompareBool(
	iA interface{},
	iB interface{},
	verbose bool) (int, error) {

	var err error
	var msg string
	var ok bool
	var bA bool
	var bB bool

	if verbose {
		fmt.Printf("Comparing Boolean Variables: '%+v' & '%+v'.\r\n",
			iA, iB)
	}

	bA, ok = iA.(bool)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Bool."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}
	bB, ok = iB.(bool)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Bool."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}

	if bA != bB {
		return RESULT_SIMILAR, nil
	}

	return RESULT_EQUAL, nil
}

// Compares JSON Bytes with Destination Object.
func CompareBytesWithObject(
	src []byte,
	dst interface{},
	config Config,
	verbose bool) (int, error) {

	var dstBA []byte
	var dstInterface interface{}
	var err error
	var result int
	var srcInterface interface{}

	err = json.Unmarshal(src, &srcInterface)
	if err != nil {
		if verbose {
			fmt.Println("<ERROR>")
		}
		return ERROR_JSON_UNMARSHAL, err
	}
	if verbose {
		fmt.Printf("%+v.\r\n", srcInterface)
	}

	dstBA, err = json.Marshal(dst)
	if err != nil {
		if verbose {
			fmt.Println("<ERROR>")
		}
		return ERROR_JSON_MARSHAL, err
	}
	err = json.Unmarshal(dstBA, &dstInterface)
	if err != nil {
		if verbose {
			fmt.Println("<ERROR>")
		}
		return ERROR_JSON_UNMARSHAL, err
	}
	if verbose {
		fmt.Printf("%+v.\r\n", dstInterface)
	}

	// Compare Interfaces.
	result, err = CompareInterfaces(srcInterface, dstInterface, config, verbose)
	return result, err
}

// Compares Complex64 Variables.
func CompareComplex64(
	iA interface{},
	iB interface{},
	verbose bool) (int, error) {

	var err error
	var msg string
	var ok bool
	var c64A complex64
	var c64B complex64

	if verbose {
		fmt.Printf("Comparing Complex64 Variables: '%+v' & '%+v'.\r\n",
			iA, iB)
	}

	c64A, ok = iA.(complex64)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Complex64."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}
	c64B, ok = iB.(complex64)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Complex64."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}

	if c64A != c64B {
		return RESULT_SIMILAR, nil
	}

	return RESULT_EQUAL, nil
}

// Compares Complex128 Variables.
func CompareComplex128(
	iA interface{},
	iB interface{},
	verbose bool) (int, error) {

	var err error
	var msg string
	var ok bool
	var c128A complex128
	var c128B complex128

	if verbose {
		fmt.Printf("Comparing Complex128 Variables: '%+v' & '%+v'.\r\n",
			iA, iB)
	}

	c128A, ok = iA.(complex128)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Complex128."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}
	c128B, ok = iB.(complex128)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Complex128."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}

	if c128A != c128B {
		return RESULT_SIMILAR, nil
	}

	return RESULT_EQUAL, nil
}

// Compares Float32 Variables.
func CompareFloat32(
	iA interface{},
	iB interface{},
	verbose bool) (int, error) {

	var err error
	var msg string
	var ok bool
	var f32A float32
	var f32B float32

	if verbose {
		fmt.Printf("Comparing Float32 Variables: '%+v' & '%+v'.\r\n",
			iA, iB)
	}

	f32A, ok = iA.(float32)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Float32."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}
	f32B, ok = iB.(float32)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Float32."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}

	if f32A != f32B {
		return RESULT_SIMILAR, nil
	}

	return RESULT_EQUAL, nil
}

// Compares Float64 Variables.
func CompareFloat64(
	iA interface{},
	iB interface{},
	verbose bool) (int, error) {

	var err error
	var msg string
	var ok bool
	var f64A float64
	var f64B float64

	if verbose {
		fmt.Printf("Comparing Float64 Variables: '%+v' & '%+v'.\r\n",
			iA, iB)
	}

	f64A, ok = iA.(float64)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Float64."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}
	f64B, ok = iB.(float64)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Float64."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}

	if f64A != f64B {
		return RESULT_SIMILAR, nil
	}

	return RESULT_EQUAL, nil
}

// Compares 'int' Variables.
func CompareInt(
	iA interface{},
	iB interface{},
	verbose bool) (int, error) {

	var err error
	var msg string
	var ok bool
	var intA int
	var intB int

	if verbose {
		fmt.Printf("Comparing 'int' Variables: '%+v' & '%+v'.\r\n",
			iA, iB)
	}

	intA, ok = iA.(int)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Int."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}
	intB, ok = iB.(int)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Int."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}

	if intA != intB {
		return RESULT_SIMILAR, nil
	}

	return RESULT_EQUAL, nil
}

// Compares 'int8' Variables.
func CompareInt8(
	iA interface{},
	iB interface{},
	verbose bool) (int, error) {

	var err error
	var msg string
	var ok bool
	var int8A int8
	var int8B int8

	if verbose {
		fmt.Printf("Comparing 'int8' Variables: '%+v' & '%+v'.\r\n",
			iA, iB)
	}

	int8A, ok = iA.(int8)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Int8."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}
	int8B, ok = iB.(int8)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Int8."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}

	if int8A != int8B {
		return RESULT_SIMILAR, nil
	}

	return RESULT_EQUAL, nil
}

// Compares 'int16' Variables.
func CompareInt16(
	iA interface{},
	iB interface{},
	verbose bool) (int, error) {

	var err error
	var msg string
	var ok bool
	var int16A int16
	var int16B int16

	if verbose {
		fmt.Printf("Comparing 'int16' Variables: '%+v' & '%+v'.\r\n",
			iA, iB)
	}

	int16A, ok = iA.(int16)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Int16."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}
	int16B, ok = iB.(int16)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Int16."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}

	if int16A != int16B {
		return RESULT_SIMILAR, nil
	}

	return RESULT_EQUAL, nil
}

// Compares 'int32' Variables.
func CompareInt32(
	iA interface{},
	iB interface{},
	verbose bool) (int, error) {

	var err error
	var msg string
	var ok bool
	var int32A int32
	var int32B int32

	if verbose {
		fmt.Printf("Comparing 'int32' Variables: '%+v' & '%+v'.\r\n",
			iA, iB)
	}

	int32A, ok = iA.(int32)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Int32."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}
	int32B, ok = iB.(int32)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Int32."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}

	if int32A != int32B {
		return RESULT_SIMILAR, nil
	}

	return RESULT_EQUAL, nil
}

// Compares 'int64' Variables.
func CompareInt64(
	iA interface{},
	iB interface{},
	verbose bool) (int, error) {

	var err error
	var msg string
	var ok bool
	var int64A int64
	var int64B int64

	if verbose {
		fmt.Printf("Comparing 'int64' Variables: '%+v' & '%+v'.\r\n",
			iA, iB)
	}

	int64A, ok = iA.(int64)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Int64."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}
	int64B, ok = iB.(int64)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Int64."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}

	if int64A != int64B {
		return RESULT_SIMILAR, nil
	}

	return RESULT_EQUAL, nil
}

// Compares Interfaces.
func CompareInterfaces(
	iA interface{},
	iB interface{},
	config Config,
	verbose bool) (int, error) {

	var err error
	var kindA reflect.Kind
	var kindB reflect.Kind
	var msg string
	var result int
	var typeA reflect.Type
	var typeB reflect.Type

	if verbose {
		fmt.Printf("Comparing Interfaces: '%+v' & '%+v'.\r\n", iA, iB)
	}

	// Compare actual Type (Kind) of Variables.
	typeA = reflect.TypeOf(iA)
	typeB = reflect.TypeOf(iB)
	if typeB == nil {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Reference Object has a Variable with 'nil' Value."
		err = errors.New(msg)
		return ERROR_REFERENCE_OBJECT, err
	}
	if typeA == nil {
		return RESULT_UNKNOWN, nil
	}
	kindA = reflect.TypeOf(iA).Kind()
	kindB = reflect.TypeOf(iB).Kind()
	if kindA != kindB {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Types are different."
		err = errors.New(msg)
		return ERROR_TYPE_MISMATCH, err
	}

	switch kindA {

	case reflect.Map:
		result, err = CompareMaps(iA, iB, config, verbose)

	case reflect.Slice:
		result, err = CompareSlices(iA, iB, config, verbose)

	case reflect.String:
		result, err = CompareStrings(iA, iB, verbose)

	case reflect.Float64:
		result, err = CompareFloat64(iA, iB, verbose)

	case reflect.Bool:
		result, err = CompareBool(iA, iB, verbose)

	case reflect.Float32:
		result, err = CompareFloat32(iA, iB, verbose)

	case reflect.Int:
		result, err = CompareInt(iA, iB, verbose)

	case reflect.Int8:
		result, err = CompareInt8(iA, iB, verbose)

	case reflect.Int16:
		result, err = CompareInt16(iA, iB, verbose)

	case reflect.Int32:
		result, err = CompareInt32(iA, iB, verbose)

	case reflect.Int64:
		result, err = CompareInt64(iA, iB, verbose)

	case reflect.Uint:
		result, err = CompareUint(iA, iB, verbose)

	case reflect.Uint8:
		result, err = CompareUint8(iA, iB, verbose)

	case reflect.Uint16:
		result, err = CompareUint16(iA, iB, verbose)

	case reflect.Uint32:
		result, err = CompareUint32(iA, iB, verbose)

	case reflect.Uint64:
		result, err = CompareUint64(iA, iB, verbose)

	case reflect.Complex64:
		result, err = CompareComplex64(iA, iB, verbose)

	case reflect.Complex128:
		result, err = CompareComplex128(iA, iB, verbose)

	default:
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Unsupported Type: '" + kindA.String() + "."
		err = errors.New(msg)
		result = ERROR_TYPE_UNSUPPORTED
	}

	return result, err
}

// Compares Maps.
func CompareMaps(
	iA interface{},
	iB interface{},
	config Config,
	verbose bool) (int, error) {

	var err error
	var key string
	var keysA []string
	var keysB []string
	var keysAreEqual bool
	var mapA map[string]interface{}
	var mapB map[string]interface{}
	var msg string
	var ok bool
	var result int
	var resultIsDifferent bool
	var resultIsEqual bool
	var resultIsUnknown bool
	var valueA interface{}
	var valueB interface{}

	if verbose {
		fmt.Printf("Comparing Maps: '%+v' & '%+v'.\r\n", iA, iB)
	}

	// 1. Compare Keys.
	keysA, err = GetMapKeyNames(iA, verbose)
	if err != nil {
		if verbose {
			fmt.Println("<ERROR>")
		}
		return ERROR_MAP_TYPE, err
	}
	keysB, err = GetMapKeyNames(iB, verbose)
	if err != nil {
		if verbose {
			fmt.Println("<ERROR>")
		}
		return ERROR_MAP_TYPE, err
	}
	keysAreEqual = MapKeysAreEqual(keysA, keysB, verbose)
	if !keysAreEqual {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Map Keys are different."
		err = errors.New(msg)
		return ERROR_MAP_KEYS_MISMATCH, err
	}

	// 2. Compare Values.
	resultIsDifferent = false
	resultIsEqual = true
	resultIsUnknown = false
	for _, key = range keysA {
		mapA, ok = iA.(map[string]interface{})
		if !ok {
			if verbose {
				fmt.Println("<ERROR>")
			}
			msg = "Can not convert Interface to Map."
			err = errors.New(msg)
			return ERROR_TYPE_CONVERTION, err
		}
		mapB, ok = iB.(map[string]interface{})
		if !ok {
			if verbose {
				fmt.Println("<ERROR>")
			}
			msg = "Can not convert Interface to Map."
			err = errors.New(msg)
			return ERROR_TYPE_CONVERTION, err
		}
		valueA = mapA[key]
		valueB = mapB[key]
		result, err = CompareInterfaces(valueA, valueB, config, verbose)

		// Return on Error.
		if err != nil {
			return result, err
		}

		// If any Comparison's Result is not 'equal',
		// then the whole Comparison's Result is not 'equal'.
		if result != RESULT_EQUAL {
			resultIsEqual = false
			// If any Comparison's Result is 'similar',
			// then the whole Comparison's Result may be 'similar'.
			// If any Comparison's Result is 'unknown',
			// then the whole Comparison's Result is 'unknown'.
			if result == RESULT_SIMILAR {
				resultIsDifferent = true
			} else if result == RESULT_UNKNOWN {
				resultIsUnknown = true
			} else {
				// Not 'equal', not 'similar', not 'unknown', no Error.
				if verbose {
					fmt.Println("<ERROR>")
				}
				return ERROR_INTERNAL, err
			}
		}
	}

	// Summary.
	if resultIsEqual {
		return RESULT_EQUAL, nil
	}
	// 'unknown' is stronger than 'similar'.
	if resultIsUnknown {
		return RESULT_UNKNOWN, nil
	}
	if resultIsDifferent {
		return RESULT_SIMILAR, nil
	}
	// This Line is unreachable.
	if verbose {
		fmt.Println("<ERROR>")
	}
	return ERROR_INTERNAL, nil
}

// Compares Slices.
func CompareSlices(
	iA interface{},
	iB interface{},
	config Config,
	verbose bool) (int, error) {

	var comparIsFull bool
	var err error
	var i int
	var msg string
	var ok bool
	var result int
	var resultIsDifferent bool
	var resultIsEqual bool
	var resultIsUnknown bool
	var sliceA []interface{}
	var sliceB []interface{}
	var sliceElementA interface{}
	var sliceElementB interface{}
	var sliceSizeA int
	var sliceSizeB int

	if verbose {
		fmt.Printf("Comparing Slices: '%+v' & '%+v'.\r\n", iA, iB)
	}

	// Read Configuration.
	if (config & CONFIG_SLICE_COMPAR_FULL) == CONFIG_SLICE_COMPAR_FULL {
		comparIsFull = true
	} else {
		comparIsFull = false
	}

	sliceA, ok = iA.([]interface{})
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Slice."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}
	sliceB, ok = iB.([]interface{})
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Slice."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}

	sliceSizeB = len(sliceB)
	if sliceSizeB == 0 {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Reference Object has empty Slice."
		err = errors.New(msg)
		return ERROR_REFERENCE_OBJECT, err
	}

	sliceSizeA = len(sliceA)
	if sliceSizeA == 0 {
		// When JSON has an empty Slice,
		// we can not compare it with the Reference Object.
		// While it is not an Error to have an empty Slice,
		// we decide that it is a normal Behaviour.
		return RESULT_UNKNOWN, nil
	}

	if !comparIsFull {
		// Compare first Elements of Slices to learn the Structure.
		sliceElementA = sliceA[0]
		sliceElementB = sliceB[0]
		result, err = CompareInterfaces(
			sliceElementA, sliceElementB, config, verbose)

		return result, err
	}

	// Full Comparison.
	if sliceSizeA != sliceSizeB {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not perform a full Slice Comparison. " +
			"Slice Sizes are different."
		err = errors.New(msg)
		return ERROR_SLICE_COMPAR_FULL_FAIL, err
	}

	// Compare all Elements.
	resultIsDifferent = false
	resultIsEqual = true
	resultIsUnknown = false
	for i = 0; i < sliceSizeA; i++ {
		sliceElementA = sliceA[i]
		sliceElementB = sliceB[i]
		result, err = CompareInterfaces(
			sliceElementA, sliceElementB, config, verbose)

		// Return on Error.
		if err != nil {
			return result, err
		}

		// If any Comparison's Result is not 'equal',
		// then the whole Comparison's Result is not 'equal'.
		if result != RESULT_EQUAL {
			resultIsEqual = false
			// If any Comparison's Result is 'similar',
			// then the whole Comparison's Result may be 'similar'.
			// If any Comparison's Result is 'unknown',
			// then the whole Comparison's Result is 'unknown'.
			if result == RESULT_SIMILAR {
				resultIsDifferent = true
			} else if result == RESULT_UNKNOWN {
				resultIsUnknown = true
			} else {
				// Not 'equal', not 'similar', not 'unknown', no Error.
				return ERROR_INTERNAL, err
			}
		}
	}

	// Summary.
	if resultIsEqual {
		return RESULT_EQUAL, nil
	}
	// 'unknown' is stronger than 'similar'.
	if resultIsUnknown {
		return RESULT_UNKNOWN, nil
	}
	if resultIsDifferent {
		return RESULT_SIMILAR, nil
	}
	// This Line is unreachable.
	return ERROR_INTERNAL, nil
}

// Compares Strings.
func CompareStrings(
	iA interface{},
	iB interface{},
	verbose bool) (int, error) {

	var err error
	var msg string
	var ok bool
	var strA string
	var strB string

	if verbose {
		fmt.Printf("Comparing Strings: '%+v' & '%+v'.\r\n", iA, iB)
	}

	strA, ok = iA.(string)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to String."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}
	strB, ok = iB.(string)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to String."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}

	if strA != strB {
		return RESULT_SIMILAR, nil
	}

	return RESULT_EQUAL, nil
}

// Compares 'uint' Variables.
func CompareUint(
	iA interface{},
	iB interface{},
	verbose bool) (int, error) {

	var err error
	var msg string
	var ok bool
	var uintA uint
	var uintB uint

	if verbose {
		fmt.Printf("Comparing 'uint' Variables: '%+v' & '%+v'.\r\n",
			iA, iB)
	}

	uintA, ok = iA.(uint)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Uint."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}
	uintB, ok = iB.(uint)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Uint."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}

	if uintA != uintB {
		return RESULT_SIMILAR, nil
	}

	return RESULT_EQUAL, nil
}

// Compares 'uint8' Variables.
func CompareUint8(
	iA interface{},
	iB interface{},
	verbose bool) (int, error) {

	var err error
	var msg string
	var ok bool
	var uint8A uint8
	var uint8B uint8

	if verbose {
		fmt.Printf("Comparing 'uint8' Variables: '%+v' & '%+v'.\r\n",
			iA, iB)
	}

	uint8A, ok = iA.(uint8)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Uint8."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}
	uint8B, ok = iB.(uint8)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Uint8."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}

	if uint8A != uint8B {
		return RESULT_SIMILAR, nil
	}

	return RESULT_EQUAL, nil
}

// Compares 'uint16' Variables.
func CompareUint16(
	iA interface{},
	iB interface{},
	verbose bool) (int, error) {

	var err error
	var msg string
	var ok bool
	var uint16A uint16
	var uint16B uint16

	if verbose {
		fmt.Printf("Comparing 'uint16' Variables: '%+v' & '%+v'.\r\n",
			iA, iB)
	}

	uint16A, ok = iA.(uint16)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Uint16."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}
	uint16B, ok = iB.(uint16)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Uint16."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}

	if uint16A != uint16B {
		return RESULT_SIMILAR, nil
	}

	return RESULT_EQUAL, nil
}

// Compares 'uint32' Variables.
func CompareUint32(
	iA interface{},
	iB interface{},
	verbose bool) (int, error) {

	var err error
	var msg string
	var ok bool
	var uint32A uint32
	var uint32B uint32

	if verbose {
		fmt.Printf("Comparing 'uint32' Variables: '%+v' & '%+v'.\r\n",
			iA, iB)
	}

	uint32A, ok = iA.(uint32)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Uint32."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}
	uint32B, ok = iB.(uint32)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Uint32."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}

	if uint32A != uint32B {
		return RESULT_SIMILAR, nil
	}

	return RESULT_EQUAL, nil
}

// Compares 'uint64' Variables.
func CompareUint64(
	iA interface{},
	iB interface{},
	verbose bool) (int, error) {

	var err error
	var msg string
	var ok bool
	var uint64A uint64
	var uint64B uint64

	if verbose {
		fmt.Printf("Comparing 'uint64' Variables: '%+v' & '%+v'.\r\n",
			iA, iB)
	}

	uint64A, ok = iA.(uint64)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Uint64."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}
	uint64B, ok = iB.(uint64)
	if !ok {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not convert Interface to Uint64."
		err = errors.New(msg)
		return ERROR_TYPE_CONVERTION, err
	}

	if uint64A != uint64B {
		return RESULT_SIMILAR, nil
	}

	return RESULT_EQUAL, nil
}

// Returns the sorted List of Map's Keys.
func GetMapKeyNames(obj interface{}, verbose bool) ([]string, error) {

	var err error
	var keyNames []string
	var mapKey reflect.Value
	var mapKeys []reflect.Value
	var msg string
	var objType reflect.Type
	var objValue reflect.Value
	var refObj map[string]interface{}

	keyNames = make([]string, 0)
	objType = reflect.TypeOf(obj)

	// Check Object's Type.
	if objType != reflect.TypeOf(refObj) {
		if verbose {
			fmt.Println("<ERROR>")
		}
		msg = "Can not get Key Names. Wrong Type of Object. " +
			"Expected: '" + reflect.TypeOf(refObj).String() + "', but " +
			"got: '" + objType.String() + "'."
		err = errors.New(msg)
		return keyNames, err
	}

	// Get Key Names.
	objValue = reflect.ValueOf(obj)
	mapKeys = objValue.MapKeys()
	for _, mapKey = range mapKeys {
		keyNames = append(keyNames, mapKey.String())
	}

	// Sort Key Names.
	sort.Strings(keyNames)

	return keyNames, err
}

// Compares Two String Slices (Map's String Keys).
func MapKeysAreEqual(a, b []string, verbose bool) bool {

	var i int
	var iMax int

	if verbose {
		fmt.Printf("Comparing Keys: %+v & %+v.\r\n", a, b)
	}

	if len(a) != len(b) {
		return false
	}

	iMax = len(a) - 1
	for i = 0; i <= iMax; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// Outputs Result in a Human-readable Way.
func PrintResult(result int, err error) {
	if err != nil {
		fmt.Println("Error!", err, MSG_DIFFERENT)
	} else {
		if result == RESULT_EQUAL {
			fmt.Println(MSG_EQUAL)
		} else if result == RESULT_SIMILAR {
			fmt.Println(MSG_SIMILAR)
		} else if result == RESULT_UNKNOWN {
			fmt.Println(MSG_UNKNOWN)
		} else {
			fmt.Println(MSG_ERROR_INTERNAL)
		}
	}
}
