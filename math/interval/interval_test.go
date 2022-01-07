package interval_test

import (
	"testing"

	"github.com/gtramontina/go-extlib/internal/assert"
	"github.com/gtramontina/go-extlib/math/interval"
)

func TestInterval(t *testing.T) {
	t.Run("open", func(t *testing.T) {
		t.Run("generates empty intervals", func(t *testing.T) {
			assert.DeepEqual(t, interval.Open(1, 0).List(1), []int{})
			assert.DeepEqual(t, interval.Open(0, 0).List(1), []int{})
			assert.DeepEqual(t, interval.Open[uint](1, 0).List(1), []uint{})
			assert.DeepEqual(t, interval.Open[float32](1, 0).List(1), []float32{})
		})

		t.Run("generates bounded intervals", func(t *testing.T) {
			assert.DeepEqual(t, interval.Open(0, 1).List(1), []int{})
			assert.DeepEqual(t, interval.Open(0, 2).List(1), []int{1})
			assert.DeepEqual(t, interval.Open(1, 5).List(1), []int{2, 3, 4})
			assert.DeepEqual(t, interval.Open[uint](1, 5).List(1), []uint{2, 3, 4})
			assert.DeepEqual(t, interval.Open[float32](1, 5).List(0.5), []float32{1.5, 2.0, 2.5, 3.0, 3.5, 4.0, 4.5})
		})

		t.Run("represents itself as string", func(t *testing.T) {
			assert.Eq(t, interval.Open(0, 0).String(), "Interval(0,0)")
			assert.Eq(t, interval.Open(0, 1).String(), "Interval(0,1)")
			assert.Eq(t, interval.Open(1, 5).String(), "Interval(1,5)")
			assert.Eq(t, interval.Open[uint](1, 5).String(), "Interval(1,5)")
			assert.Eq(t, interval.Open[float32](1.1, 5.5).String(), "Interval(1.1,5.5)")
		})
	})

	t.Run("left closed right open", func(t *testing.T) {
		t.Run("generates empty intervals", func(t *testing.T) {
			assert.DeepEqual(t, interval.LeftClosedRightOpen(1, 0).List(1), []int{})
			assert.DeepEqual(t, interval.LeftClosedRightOpen(0, 0).List(1), []int{})
			assert.DeepEqual(t, interval.LeftClosedRightOpen[uint](1, 0).List(1), []uint{})
			assert.DeepEqual(t, interval.LeftClosedRightOpen[float32](1, 0).List(1), []float32{})
		})

		t.Run("generates bounded intervals", func(t *testing.T) {
			assert.DeepEqual(t, interval.LeftClosedRightOpen(0, 1).List(1), []int{0})
			assert.DeepEqual(t, interval.LeftClosedRightOpen(0, 2).List(1), []int{0, 1})
			assert.DeepEqual(t, interval.LeftClosedRightOpen(1, 5).List(1), []int{1, 2, 3, 4})
			assert.DeepEqual(t, interval.LeftClosedRightOpen[uint](1, 5).List(1), []uint{1, 2, 3, 4})
			assert.DeepEqual(t, interval.LeftClosedRightOpen[float32](1, 5).List(0.5), []float32{1, 1.5, 2.0, 2.5, 3.0, 3.5, 4.0, 4.5})
		})

		t.Run("represents itself as string", func(t *testing.T) {
			assert.Eq(t, interval.LeftClosedRightOpen(0, 0).String(), "Interval[0,0)")
			assert.Eq(t, interval.LeftClosedRightOpen(0, 1).String(), "Interval[0,1)")
			assert.Eq(t, interval.LeftClosedRightOpen(1, 5).String(), "Interval[1,5)")
			assert.Eq(t, interval.LeftClosedRightOpen[uint](1, 5).String(), "Interval[1,5)")
			assert.Eq(t, interval.LeftClosedRightOpen[float32](1.1, 5.5).String(), "Interval[1.1,5.5)")
		})
	})

	t.Run("left open right closed", func(t *testing.T) {
		t.Run("generates empty intervals", func(t *testing.T) {
			assert.DeepEqual(t, interval.LeftOpenRightClosed(1, 0).List(1), []int{})
			assert.DeepEqual(t, interval.LeftOpenRightClosed(0, 0).List(1), []int{})
			assert.DeepEqual(t, interval.LeftOpenRightClosed[uint](1, 0).List(1), []uint{})
			assert.DeepEqual(t, interval.LeftOpenRightClosed[float32](1, 0).List(1), []float32{})
		})

		t.Run("generates bounded intervals", func(t *testing.T) {
			assert.DeepEqual(t, interval.LeftOpenRightClosed(0, 1).List(1), []int{1})
			assert.DeepEqual(t, interval.LeftOpenRightClosed(0, 2).List(1), []int{1, 2})
			assert.DeepEqual(t, interval.LeftOpenRightClosed(1, 5).List(1), []int{2, 3, 4, 5})
			assert.DeepEqual(t, interval.LeftOpenRightClosed[uint](1, 5).List(1), []uint{2, 3, 4, 5})
			assert.DeepEqual(t, interval.LeftOpenRightClosed[float32](1, 5).List(0.5), []float32{1.5, 2.0, 2.5, 3.0, 3.5, 4.0, 4.5, 5.0})
		})

		t.Run("represents itself as string", func(t *testing.T) {
			assert.Eq(t, interval.LeftOpenRightClosed(0, 0).String(), "Interval(0,0]")
			assert.Eq(t, interval.LeftOpenRightClosed(0, 1).String(), "Interval(0,1]")
			assert.Eq(t, interval.LeftOpenRightClosed(1, 5).String(), "Interval(1,5]")
			assert.Eq(t, interval.LeftOpenRightClosed[uint](1, 5).String(), "Interval(1,5]")
			assert.Eq(t, interval.LeftOpenRightClosed[float32](1.1, 5.5).String(), "Interval(1.1,5.5]")
		})
	})

	t.Run("closed", func(t *testing.T) {
		t.Run("generates degenerate intervals", func(t *testing.T) {
			assert.DeepEqual(t, interval.Closed(0, 0).List(1), []int{0})
			assert.DeepEqual(t, interval.Closed(1, 1).List(1), []int{1})
			assert.DeepEqual(t, interval.Closed[uint](1, 1).List(1), []uint{1})
			assert.DeepEqual(t, interval.Closed[float32](1.1, 1.1).List(1), []float32{1.1})
		})

		t.Run("generates empty intervals", func(t *testing.T) {
			assert.DeepEqual(t, interval.Closed(1, 0).List(1), []int{})
			assert.DeepEqual(t, interval.Closed[uint](1, 0).List(1), []uint{})
			assert.DeepEqual(t, interval.Closed[float32](1, 0).List(1), []float32{})
		})

		t.Run("generates bounded intervals", func(t *testing.T) {
			assert.DeepEqual(t, interval.Closed(0, 1).List(1), []int{0, 1})
			assert.DeepEqual(t, interval.Closed(0, 2).List(1), []int{0, 1, 2})
			assert.DeepEqual(t, interval.Closed(1, 5).List(1), []int{1, 2, 3, 4, 5})
			assert.DeepEqual(t, interval.Closed[uint](1, 5).List(1), []uint{1, 2, 3, 4, 5})
			assert.DeepEqual(t, interval.Closed[float32](1, 5).List(0.5), []float32{1.0, 1.5, 2.0, 2.5, 3.0, 3.5, 4.0, 4.5, 5.0})
		})

		t.Run("represents itself as string", func(t *testing.T) {
			assert.Eq(t, interval.Closed(0, 0).String(), "Interval[0,0]")
			assert.Eq(t, interval.Closed(0, 1).String(), "Interval[0,1]")
			assert.Eq(t, interval.Closed(1, 5).String(), "Interval[1,5]")
			assert.Eq(t, interval.Closed[uint](1, 5).String(), "Interval[1,5]")
			assert.Eq(t, interval.Closed[float32](1.1, 5.5).String(), "Interval[1.1,5.5]")
		})
	})
}
