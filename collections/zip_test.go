package collections_test

import (
	"testing"

	"github.com/gtramontina/go-extlib/collections"
	"github.com/gtramontina/go-extlib/testing/assert"
	"github.com/gtramontina/go-extlib/tuple"
)

func TestZip(t *testing.T) {
	t.Run("empty collections yields empty collection of tuples", func(t *testing.T) {
		assert.DeepEqual(t, collections.Zip([]int{}, []int{}), []tuple.OfTwo[int, int]{})
	})

	t.Run("collections of same size yields collection of tuples of the same size", func(t *testing.T) {
		assert.DeepEqual(t, collections.Zip([]int{1}, []int{2}), []tuple.OfTwo[int, int]{
			tuple.Of2(1, 2),
		})

		assert.DeepEqual(t, collections.Zip([]int{1, 2}, []int{3, 4}), []tuple.OfTwo[int, int]{
			tuple.Of2(1, 3),
			tuple.Of2(2, 4),
		})

		assert.DeepEqual(t, collections.Zip([]int{1, 2, 3}, []int{4, 5, 6}), []tuple.OfTwo[int, int]{
			tuple.Of2(1, 4),
			tuple.Of2(2, 5),
			tuple.Of2(3, 6),
		})
	})

	t.Run("it is limited by the smaller collection", func(t *testing.T) {
		assert.DeepEqual(t, collections.Zip([]int{1, 2, 3}, []int{4, 5}), []tuple.OfTwo[int, int]{
			tuple.Of2(1, 4),
			tuple.Of2(2, 5),
		})

		assert.DeepEqual(t, collections.Zip([]int{1, 2}, []int{4, 5, 6}), []tuple.OfTwo[int, int]{
			tuple.Of2(1, 4),
			tuple.Of2(2, 5),
		})
	})

	t.Run("works with different types", func(t *testing.T) {
		assert.DeepEqual(t, collections.Zip([]int{1, 2, 3}, []string{"a", "b", "c"}), []tuple.OfTwo[int, string]{
			tuple.Of2(1, "a"),
			tuple.Of2(2, "b"),
			tuple.Of2(3, "c"),
		})

		assert.DeepEqual(t, collections.Zip([]string{"a", "b", "c"}, []int{1, 2, 3}), []tuple.OfTwo[string, int]{
			tuple.Of2("a", 1),
			tuple.Of2("b", 2),
			tuple.Of2("c", 3),
		})
	})
}
