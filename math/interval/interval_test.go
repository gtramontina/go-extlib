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
		})

		t.Run("generates bounded intervals", func(t *testing.T) {
			assert.DeepEqual(t, interval.Open(0, 1).List(1), []int{})
			assert.DeepEqual(t, interval.Open(0, 2).List(1), []int{1})
			assert.DeepEqual(t, interval.Open(1, 5).List(1), []int{2, 3, 4})
		})

		t.Run("represents itself as string", func(t *testing.T) {
			assert.Eq(t, interval.Open(0, 0).String(), "Interval(0,0)")
			assert.Eq(t, interval.Open(0, 1).String(), "Interval(0,1)")
			assert.Eq(t, interval.Open(1, 5).String(), "Interval(1,5)")
		})
	})

	t.Run("left closed right open", func(t *testing.T) {
		t.Run("generates empty intervals", func(t *testing.T) {
			assert.DeepEqual(t, interval.LeftClosedRightOpen(1, 0).List(1), []int{})
			assert.DeepEqual(t, interval.LeftClosedRightOpen(0, 0).List(1), []int{})
		})

		t.Run("generates bounded intervals", func(t *testing.T) {
			assert.DeepEqual(t, interval.LeftClosedRightOpen(0, 1).List(1), []int{0})
			assert.DeepEqual(t, interval.LeftClosedRightOpen(0, 2).List(1), []int{0, 1})
			assert.DeepEqual(t, interval.LeftClosedRightOpen(1, 5).List(1), []int{1, 2, 3, 4})
		})

		t.Run("represents itself as string", func(t *testing.T) {
			assert.Eq(t, interval.LeftClosedRightOpen(0, 0).String(), "Interval[0,0)")
			assert.Eq(t, interval.LeftClosedRightOpen(0, 1).String(), "Interval[0,1)")
			assert.Eq(t, interval.LeftClosedRightOpen(1, 5).String(), "Interval[1,5)")
		})
	})

	t.Run("left open right closed", func(t *testing.T) {
		t.Run("generates empty intervals", func(t *testing.T) {
			assert.DeepEqual(t, interval.LeftOpenRightClosed(1, 0).List(1), []int{})
			assert.DeepEqual(t, interval.LeftOpenRightClosed(0, 0).List(1), []int{})
		})

		t.Run("generates bounded intervals", func(t *testing.T) {
			assert.DeepEqual(t, interval.LeftOpenRightClosed(0, 1).List(1), []int{1})
			assert.DeepEqual(t, interval.LeftOpenRightClosed(0, 2).List(1), []int{1, 2})
			assert.DeepEqual(t, interval.LeftOpenRightClosed(1, 5).List(1), []int{2, 3, 4, 5})
		})

		t.Run("represents itself as string", func(t *testing.T) {
			assert.Eq(t, interval.LeftOpenRightClosed(0, 0).String(), "Interval(0,0]")
			assert.Eq(t, interval.LeftOpenRightClosed(0, 1).String(), "Interval(0,1]")
			assert.Eq(t, interval.LeftOpenRightClosed(1, 5).String(), "Interval(1,5]")
		})
	})

	t.Run("closed", func(t *testing.T) {
		t.Run("generates empty intervals", func(t *testing.T) {
			assert.DeepEqual(t, interval.Closed(1, 0).List(1), []int{})
		})

		t.Run("generates degenerate intervals", func(t *testing.T) {
			assert.DeepEqual(t, interval.Closed(0, 0).List(1), []int{0})
			assert.DeepEqual(t, interval.Closed(1, 1).List(1), []int{1})
		})

		t.Run("generates bounded intervals", func(t *testing.T) {
			assert.DeepEqual(t, interval.Closed(0, 1).List(1), []int{0, 1})
			assert.DeepEqual(t, interval.Closed(0, 2).List(1), []int{0, 1, 2})
			assert.DeepEqual(t, interval.Closed(1, 5).List(1), []int{1, 2, 3, 4, 5})
		})

		t.Run("represents itself as string", func(t *testing.T) {
			assert.Eq(t, interval.Closed(0, 0).String(), "Interval[0,0]")
			assert.Eq(t, interval.Closed(0, 1).String(), "Interval[0,1]")
			assert.Eq(t, interval.Closed(1, 5).String(), "Interval[1,5]")
		})
	})
}
