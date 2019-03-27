// private.go.

package incrementor

// Private Methods of Incrementor.

// Gets the internal Number.
func (this Incrementor) getNumber() int {

	return this.number
}

// Initializes the Incrementor:
//	* Sets default Limits;
//	* Sets default internal Number.
func (this *Incrementor) init() {

	this.setMinimumValue(MinValueDefault)
	this.setMaximumValue(MaxValueDefault)
	this.resetNumber()
}

// Resets the internal Number to its minimal Value.
func (this *Incrementor) resetNumber() {

	this.number = this.numberMin
}

// Sets the maximum allowed Value for the internal Number.
func (this *Incrementor) setMaximumValue(
	maxValue int,
) {

	this.numberMax = maxValue
}

// Sets the minimum allowed Value for the internal Number.
func (this *Incrementor) setMinimumValue(
	minValue int,
) {

	this.numberMin = minValue
}
