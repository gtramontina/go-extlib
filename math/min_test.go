package math_test

import (
	m "math"
	"testing"

	"github.com/gtramontina/go-extlib/math"
	"github.com/gtramontina/go-extlib/testing/assert"
)

func TestMin(t *testing.T) {
	Inf := m.Inf
	NaN := m.NaN
	IsNaN := m.IsNaN

	assert.Eq(t, math.Min[int](0, 0), 0)
	assert.Eq(t, math.Min[int](0, 1), 0)
	assert.Eq(t, math.Min[int](1, 0), 0)
	assert.Eq(t, math.Min[int](1, 1), 1)
	assert.Eq(t, math.Min[int](+0, +0), +0)
	assert.Eq(t, math.Min[int](+0, -0), -0)
	assert.Eq(t, math.Min[int](-0, +0), -0)
	assert.Eq(t, math.Min[int](-0, -0), -0)
	assert.Eq(t, math.Min[int](1, int(Inf(-1))), int(Inf(-1)))
	assert.Eq(t, math.Min[int](int(Inf(-1)), 1), int(Inf(-1)))

	assert.Eq(t, math.Min[float64](0, 0), 0)
	assert.Eq(t, math.Min[float64](0, 1), 0)
	assert.Eq(t, math.Min[float64](1, 0), 0)
	assert.Eq(t, math.Min[float64](1, 1), 1)
	assert.Eq(t, math.Min[float64](+0, +0), +0)
	assert.Eq(t, math.Min[float64](+0, -0), -0)
	assert.Eq(t, math.Min[float64](-0, +0), -0)
	assert.Eq(t, math.Min[float64](-0, -0), -0)
	assert.Eq(t, math.Min[float64](1, Inf(-1)), Inf(-1))
	assert.Eq(t, math.Min[float64](Inf(-1), 1), Inf(-1))
	assert.True(t, IsNaN(math.Min[float64](100, NaN())))
	assert.True(t, IsNaN(math.Min[float64](NaN(), 100)))
}
