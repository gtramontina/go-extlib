package math

import (
	"math"

	"github.com/gtramontina/go-extlib/math/internal"
)

// Min returns the smallest of x or y.
func Min[Real internal.Real](x, y Real) Real {
	return Real(math.Min(float64(x), float64(y)))
}
