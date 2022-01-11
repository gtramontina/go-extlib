package collections_test

import (
	"testing"

	"github.com/gtramontina/go-extlib/collections"
	"github.com/gtramontina/go-extlib/testing/assert"
)

func TestContains(t *testing.T) {
	type sample struct{ value int }

	t.Run("finds whether a collection contains a given value", func(t *testing.T) {
		assert.False(t, collections.Contains([]int{}, 1))
		assert.True(t, collections.Contains([]int{1}, 1))
		assert.False(t, collections.Contains([]int{0}, 1))
		assert.True(t, collections.Contains([]int{0, 1}, 1))
		assert.False(t, collections.Contains([]int{0, 1}, 2))

		assert.False(t, collections.Contains([]string{}, "a"))
		assert.True(t, collections.Contains([]string{"b"}, "b"))
		assert.False(t, collections.Contains([]string{"a"}, "b"))
		assert.True(t, collections.Contains([]string{"a", "b"}, "b"))
		assert.False(t, collections.Contains([]string{"a", "b"}, "c"))

		assert.False(t, collections.Contains([]sample{}, sample{1}))
		assert.True(t, collections.Contains([]sample{{1}}, sample{1}))
		assert.False(t, collections.Contains([]sample{{0}}, sample{1}))
		assert.True(t, collections.Contains([]sample{{0}, {1}}, sample{1}))
		assert.False(t, collections.Contains([]sample{{0}, {1}}, sample{2}))
	})
}
