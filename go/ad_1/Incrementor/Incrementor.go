// Incrementor.go.

package incrementor

// Incrementor Type ("Class").

// Incrementor Type ("Class") allows storing and incrementing (using the Step
// of 1) an integer Value. The upper Limit for Increment can be set manually,
// the lower Limit is Zero and cannot be changed manually. By default, uses
// the maximum integer Value (of the current Machine's Architecture) as the
// upper Limit.

// Incrementor internal Number Value's default Limits.
const (
	MaxValueDefault = MaxInt
	MinValueDefault = 0
)

// Errors.
const (
	ErrfMaxValue = "Upper Limit Value Error: %v"
)

type Incrementor struct {

	// Internal Number.
	number int

	// Internal Number Limits.
	numberMin int
	numberMax int
}
