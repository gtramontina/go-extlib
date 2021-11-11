package collections_test

import (
	"testing"

	"github.com/gtramontina/go-extlib/collections"
	"github.com/gtramontina/go-extlib/internal/assert"
)

func TestReverse(t *testing.T) {
	type sample struct{ value int }

	t.Run("reverses the order of the elements of the given collection", func(t *testing.T) {
		assert.DeepEqual(t, collections.Reverse([]int{}), []int{})
		assert.DeepEqual(t, collections.Reverse([]int{1}), []int{1})
		assert.DeepEqual(t, collections.Reverse([]int{1, 2}), []int{2, 1})
		assert.DeepEqual(t, collections.Reverse([]int{1, 2, 3}), []int{3, 2, 1})

		assert.DeepEqual(t, collections.Reverse([]string{}), []string{})
		assert.DeepEqual(t, collections.Reverse([]string{"a"}), []string{"a"})
		assert.DeepEqual(t, collections.Reverse([]string{"a", "b"}), []string{"b", "a"})
		assert.DeepEqual(t, collections.Reverse([]string{"a", "b", "c"}), []string{"c", "b", "a"})

		assert.DeepEqual(t, collections.Reverse([]sample{}), []sample{})
		assert.DeepEqual(t, collections.Reverse([]sample{{1}}), []sample{{1}})
		assert.DeepEqual(t, collections.Reverse([]sample{{1}, {2}}), []sample{{2}, {1}})
		assert.DeepEqual(t, collections.Reverse([]sample{{1}, {2}, {3}}), []sample{{3}, {2}, {1}})

		assert.DeepEqual(t, collections.Reverse(collections.Reverse([]int{1, 2, 3})), []int{1, 2, 3})
	})
}
