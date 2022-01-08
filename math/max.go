package math

import (
	"math"

	"github.com/gtramontina/go-extlib/math/internal"
)

// Max returns the largest of x or y.
func Max[Real internal.Real](x, y Real) Real {
	return Real(math.Max(float64(x), float64(y)))
}
