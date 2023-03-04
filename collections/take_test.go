package collections_test

import (
	"testing"

	"github.com/gtramontina/go-extlib/collections"
	"github.com/gtramontina/go-extlib/testing/assert"
)

func TestTake(t *testing.T) {
	t.Run("take 0 from empty slice", func(t *testing.T) {
		taken := collections.Take([]string{}, 0)
		assert.DeepEqual(t, taken, []string{})
	})

	t.Run("takes from empty slice", func(t *testing.T) {
		taken := collections.Take([]string{}, 1)
		assert.DeepEqual(t, taken, []string{})
	})

	t.Run("take 0 from slice of size 1", func(t *testing.T) {
		taken := collections.Take([]string{"a"}, 0)
		assert.DeepEqual(t, taken, []string{})
	})

	t.Run("take 1 from slice of size 2", func(t *testing.T) {
		taken := collections.Take([]string{"a", "b"}, 1)
		assert.DeepEqual(t, taken, []string{"a"})
	})

	t.Run("take many from a larger slice", func(t *testing.T) {
		taken := collections.Take([]int{1, 2, 3, 4, 5, 6, 7, 8}, 5)
		assert.DeepEqual(t, taken, []int{1, 2, 3, 4, 5})
	})

	t.Run("take more than the slice size", func(t *testing.T) {
		taken := collections.Take([]int{1, 2, 3, 4, 5, 6, 7, 8}, 10)
		assert.DeepEqual(t, taken, []int{1, 2, 3, 4, 5, 6, 7, 8})
	})
}
