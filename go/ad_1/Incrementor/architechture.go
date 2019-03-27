// architechture.go.

package incrementor

// Machine's Architecture Parameters.

import (
	"math/bits"
)

// Built-in Type Limits.
const (
	MaxInt int = (1<<bits.UintSize)/2 - 1
	MinInt int = (1 << bits.UintSize) / -2
)
