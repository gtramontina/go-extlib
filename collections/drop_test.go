package collections_test

import (
	"testing"

	"github.com/gtramontina/go-extlib/collections"
	"github.com/gtramontina/go-extlib/testing/assert"
)

func TestDrop(t *testing.T) {
	t.Run("drop 0 from an empty slice", func(t *testing.T) {
		dropped := collections.Drop([]string{}, 0)
		assert.DeepEqual(t, dropped, []string{})
	})

	t.Run("drops from empty slice", func(t *testing.T) {
		taken := collections.Drop([]string{}, 1)
		assert.DeepEqual(t, taken, []string{})
	})

	t.Run("drop 0 from slice of size 1", func(t *testing.T) {
		taken := collections.Drop([]string{"a"}, 0)
		assert.DeepEqual(t, taken, []string{"a"})
	})

	t.Run("drop 1 from slice of size 2", func(t *testing.T) {
		taken := collections.Drop([]string{"a", "b"}, 1)
		assert.DeepEqual(t, taken, []string{"b"})
	})

	t.Run("drop many from a larger slice", func(t *testing.T) {
		taken := collections.Drop([]int{1, 2, 3, 4, 5, 6, 7, 8}, 5)
		assert.DeepEqual(t, taken, []int{6, 7, 8})
	})

	t.Run("drop more than the slice size", func(t *testing.T) {
		taken := collections.Drop([]int{1, 2, 3, 4, 5, 6, 7, 8}, 10)
		assert.DeepEqual(t, taken, []int{})
	})
}
