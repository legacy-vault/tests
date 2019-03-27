// public.go.

package incrementor

import "fmt"

// Public Methods of Incrementor.

// Returns the internal Number.
func (this Incrementor) GetNumber() int {

	return this.getNumber()
}

// Increments the internal Number.
func (this *Incrementor) IncrementNumber() {

	if this.number == this.numberMax {
		this.resetNumber()
		return
	}

	this.number++
}

// Sets the maximum allowed Value for the internal Number.
func (this *Incrementor) SetMaximumValue(
	maxValue int,
) error {

	var err error

	// New Limit Check.
	if maxValue < 0 {
		err = fmt.Errorf(ErrfMaxValue, maxValue)
		return err
	}

	// Set the upper Limit.
	this.setMaximumValue(maxValue)

	// Compare current Value with new upper Limit.
	if this.number > this.numberMax {
		this.resetNumber()
	}
	return nil
}

// Creates a new initialized Incrementor.
func New() Incrementor {

	var inc Incrementor

	inc.init()
	return inc
}
