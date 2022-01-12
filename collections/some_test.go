package collections_test

import (
	"testing"

	"github.com/gtramontina/go-extlib/collections"
	"github.com/gtramontina/go-extlib/testing/assert"
)

func TestSome(t *testing.T) {
	t.Run("empty collection is always false", func(t *testing.T) {
		dummy := func(int) bool { return true }
		assert.False(t, collections.Some([]int{}, dummy))
	})

	t.Run("runs the predicate on all elements", func(t *testing.T) {
		hasEven := func(it int) bool { return it%2 == 0 }
		assert.True(t, collections.Some([]int{0}, hasEven))
		assert.False(t, collections.Some([]int{1}, hasEven))
		assert.False(t, collections.Some([]int{1, 3}, hasEven))
		assert.True(t, collections.Some([]int{1, 0}, hasEven))
		assert.True(t, collections.Some([]int{2, 0}, hasEven))
		assert.True(t, collections.Some([]int{11, 19, 16}, hasEven))

		hasOdd := func(it int) bool { return it%2 != 0 }
		assert.True(t, collections.Some([]int{1}, hasOdd))
		assert.False(t, collections.Some([]int{0}, hasOdd))
		assert.False(t, collections.Some([]int{2, 0}, hasOdd))
		assert.True(t, collections.Some([]int{2, 1}, hasOdd))
		assert.True(t, collections.Some([]int{1, 0}, hasOdd))
		assert.True(t, collections.Some([]int{12, 16, 15}, hasOdd))
	})

	t.Run("handles different types", func(t *testing.T) {
		hasEmpty := func(it string) bool { return len(it) == 0 }
		assert.False(t, collections.Some([]string{"a"}, hasEmpty))
		assert.True(t, collections.Some([]string{""}, hasEmpty))
		assert.False(t, collections.Some([]string{"a", "b", "c"}, hasEmpty))
		assert.True(t, collections.Some([]string{"", "b", "c"}, hasEmpty))
		assert.True(t, collections.Some([]string{"a", "", "c"}, hasEmpty))
		assert.True(t, collections.Some([]string{"a", "b", ""}, hasEmpty))
	})
}
