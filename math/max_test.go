package math_test

import (
	"testing"

	m "math"

	"github.com/gtramontina/go-extlib/internal/assert"
	"github.com/gtramontina/go-extlib/math"
)

func TestMax(t *testing.T) {
	var Inf = m.Inf
	var NaN = m.NaN
	var IsNaN = m.IsNaN

	assert.Eq(t, math.Max[int](0, 0), 0)
	assert.Eq(t, math.Max[int](0, 1), 1)
	assert.Eq(t, math.Max[int](1, 0), 1)
	assert.Eq(t, math.Max[int](1, 1), 1)
	assert.Eq(t, math.Max[int](+0, +0), +0)
	assert.Eq(t, math.Max[int](+0, -0), +0)
	assert.Eq(t, math.Max[int](-0, +0), +0)
	assert.Eq(t, math.Max[int](-0, -0), -0)

	assert.Eq(t, math.Max[float64](0, 0), 0)
	assert.Eq(t, math.Max[float64](0, 1), 1)
	assert.Eq(t, math.Max[float64](1, 0), 1)
	assert.Eq(t, math.Max[float64](1, 1), 1)
	assert.Eq(t, math.Max[float64](+0, +0), +0)
	assert.Eq(t, math.Max[float64](+0, -0), +0)
	assert.Eq(t, math.Max[float64](-0, +0), +0)
	assert.Eq(t, math.Max[float64](-0, -0), -0)
	assert.Eq(t, math.Max[float64](1, Inf(+1)), Inf(+1))
	assert.Eq(t, math.Max[float64](Inf(+1), 1), Inf(+1))
	assert.True(t, IsNaN(math.Max[float64](100, NaN())))
	assert.True(t, IsNaN(math.Max[float64](NaN(), 100)))
}
