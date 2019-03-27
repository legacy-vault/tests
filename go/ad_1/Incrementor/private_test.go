// private_test.go.

package incrementor

// Tests of Incrementor's private Methods.

import (
	"testing"
)

func Test_getNumber(t *testing.T) {

	var inc Incrementor
	var resultExpected int
	var resultReceived int

	// Prepare the Test.
	inc.number = 123
	resultExpected = inc.number

	// Perform the Test.
	resultReceived = inc.getNumber()
	if resultReceived != resultExpected {
		t.Error("Error in 'getNumber' Method")
		t.FailNow()
	}
}

func Test_init(t *testing.T) {

	var inc Incrementor

	// Prepare the Test.
	inc.number = 555
	inc.numberMin = 111
	inc.numberMax = 999

	// Perform the Test.
	inc.init()
	if inc.numberMin != MinValueDefault {
		t.Error("Error in 'init' Method")
		t.FailNow()
	}
	if inc.numberMax != MaxValueDefault {
		t.Error("Error in 'init' Method")
		t.FailNow()
	}
	if inc.number != inc.numberMin {
		t.Error("Error in 'init' Method")
		t.FailNow()
	}
}

func Test_resetNumber(t *testing.T) {

	var inc Incrementor
	var resultExpected int
	var resultReceived int

	// Prepare the Test.
	inc.number = 123
	inc.numberMin = 101
	resultExpected = inc.numberMin

	// Perform the Test.
	inc.resetNumber()
	resultReceived = inc.number
	if resultReceived != resultExpected {
		t.Error("Error in 'resetNumber' Method")
		t.FailNow()
	}
}

func Test_setMaximumValue(t *testing.T) {

	var inc Incrementor
	var resultExpected int
	var resultReceived int

	// Prepare the Test.
	inc.numberMax = 100
	resultExpected = 200

	// Perform the Test.
	inc.setMaximumValue(resultExpected)
	resultReceived = inc.numberMax
	if resultReceived != resultExpected {
		t.Error("Error in 'setMaximumValue' Method")
		t.FailNow()
	}
}

func Test_setMinimumValue(t *testing.T) {

	var inc Incrementor
	var resultExpected int
	var resultReceived int

	// Prepare the Test.
	inc.numberMin = 200
	resultExpected = 100

	// Perform the Test.
	inc.setMinimumValue(resultExpected)
	resultReceived = inc.numberMin
	if resultReceived != resultExpected {
		t.Error("Error in 'setMinimumValue' Method")
		t.FailNow()
	}
}
