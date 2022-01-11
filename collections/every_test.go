package collections_test

import (
	"testing"

	"github.com/gtramontina/go-extlib/collections"
	"github.com/gtramontina/go-extlib/testing/assert"
)

func TestEvery(t *testing.T) {
	t.Run("empty collection is always true", func(t *testing.T) {
		dummy := func(int) bool { return true }
		assert.True(t, collections.Every([]int{}, dummy))
	})

	t.Run("runs the predicate on all elements", func(t *testing.T) {
		allEven := func(it int) bool { return it%2 == 0 }
		assert.True(t, collections.Every([]int{0}, allEven))
		assert.False(t, collections.Every([]int{0, 1}, allEven))
		assert.False(t, collections.Every([]int{1, 0}, allEven))
		assert.True(t, collections.Every([]int{0, 2}, allEven))
		assert.True(t, collections.Every([]int{2, 0}, allEven))
		assert.True(t, collections.Every([]int{12, 16, 14}, allEven))

		allOdd := func(it int) bool { return it%2 != 0 }
		assert.True(t, collections.Every([]int{1}, allOdd))
		assert.False(t, collections.Every([]int{1, 2}, allOdd))
		assert.False(t, collections.Every([]int{2, 1}, allOdd))
		assert.True(t, collections.Every([]int{1, 3}, allOdd))
		assert.True(t, collections.Every([]int{3, 1}, allOdd))
		assert.True(t, collections.Every([]int{11, 19, 15}, allOdd))
	})

	t.Run("handles different types", func(t *testing.T) {
		notEmpty := func(it string) bool { return len(it) > 0 }
		assert.True(t, collections.Every([]string{"a"}, notEmpty))
		assert.False(t, collections.Every([]string{""}, notEmpty))
		assert.True(t, collections.Every([]string{"a", "b", "c"}, notEmpty))
		assert.False(t, collections.Every([]string{"", "b", "c"}, notEmpty))
		assert.False(t, collections.Every([]string{"a", "", "c"}, notEmpty))
		assert.False(t, collections.Every([]string{"a", "b", ""}, notEmpty))
	})
}
