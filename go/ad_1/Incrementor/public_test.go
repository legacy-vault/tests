// public_test.go.

package incrementor

// Tests of Incrementor's public Methods.

import (
	"fmt"
	"math/bits"
	"testing"
)

// This Test depends on the Machine's Architecture and is, thus, manual.
func Test_ArchitectureLimits(t *testing.T) {

	// Manual Test.
	fmt.Println("Bits in an unsigned Integer:", bits.UintSize)
	fmt.Println("MinInt:", MinInt)
	fmt.Println("MaxInt:", MaxInt)
}

func Test_GetNumber(t *testing.T) {

	var inc Incrementor
	var result int

	// Prepare the Test.
	inc.number = 100

	// Perform the Test.
	result = inc.GetNumber()
	if result != inc.number {
		t.Error("Error in 'GetNumber' Method")
		t.FailNow()
	}
}

func Test_IncrementNumber(t *testing.T) {

	var inc Incrementor
	var result int

	// Test 1: Normal Increment.
	inc.number = 10
	inc.numberMax = 20
	inc.IncrementNumber()
	result = inc.number
	if result != (10 + 1) {
		t.Error("Error in 'IncrementNumber' Method")
		t.FailNow()
	}

	// Test 2: Increment with Overflow.
	inc.number = 20
	inc.numberMax = 20
	inc.numberMin = 10
	inc.IncrementNumber()
	result = inc.number
	if result != inc.numberMin {
		t.Error("Error in 'IncrementNumber' Method")
		t.FailNow()
	}
}

func Test_SetMaximumValue(t *testing.T) {

	var err error
	var inc Incrementor

	// Test 1: Negative upper Limit.
	err = inc.SetMaximumValue(-1)
	if err == nil {
		t.Error("Error in 'SetMaximumValue' Method")
		t.FailNow()
	}

	// Test 2: Normal Behaviour.
	inc.number = 10
	inc.numberMax = 10
	err = inc.SetMaximumValue(20)
	if err != nil {
		t.Error("Error in 'SetMaximumValue' Method")
		t.FailNow()
	}
	if inc.numberMax != 20 {
		t.Error("Error in 'SetMaximumValue' Method")
		t.FailNow()
	}
	if inc.number != 10 {
		t.Error("Error in 'SetMaximumValue' Method")
		t.FailNow()
	}

	// Test 3: Upper Limit breaks the current Value.
	inc.number = 20
	inc.numberMin = 5
	inc.numberMax = 20
	err = inc.SetMaximumValue(10)
	if err != nil {
		t.Error("Error in 'SetMaximumValue' Method")
		t.FailNow()
	}
	if inc.numberMax != 10 {
		t.Error("Error in 'SetMaximumValue' Method")
		t.FailNow()
	}
	if inc.number != inc.numberMin {
		t.Error("Error in 'SetMaximumValue' Method")
		t.FailNow()
	}
}

func Test_New(t *testing.T) {

	var inc Incrementor

	inc = New()
	if inc.numberMin != MinValueDefault {
		t.Error("Error in 'New' Method")
		t.FailNow()
	}
	if inc.numberMax != MaxValueDefault {
		t.Error("Error in 'New' Method")
		t.FailNow()
	}
	if inc.number != inc.numberMin {
		t.Error("Error in 'New' Method")
		t.FailNow()
	}
}
