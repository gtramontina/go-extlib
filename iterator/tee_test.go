package iterator_test

import (
	"testing"

	"github.com/gtramontina/go-extlib/iterator"
	"github.com/gtramontina/go-extlib/testing/assert"
)

func TestTee(t *testing.T) {
	t.Run("with an empty iterator", func(t *testing.T) {
		iter1, iter2 := iterator.Tee(iterator.FromSlice([]int{}))
		assert.DeepEqual(t, iter1.HasNext(), false)
		assert.DeepEqual(t, iter2.HasNext(), false)
	})

	t.Run("with an iterator with a single element", func(t *testing.T) {
		iter1, iter2 := iterator.Tee(iterator.FromSlice([]int{1}))
		assert.DeepEqual(t, iter1.HasNext(), true)
		assert.DeepEqual(t, iter1.Next(), 1)

		assert.DeepEqual(t, iter2.HasNext(), true)
		assert.DeepEqual(t, iter2.Next(), 1)
	})

	t.Run("with an iterator with multiple elements", func(t *testing.T) {
		iter1, iter2 := iterator.Tee(iterator.FromSlice([]int{1, 2, 3}))
		assert.DeepEqual(t, iter1.HasNext(), true)
		assert.DeepEqual(t, iter1.Next(), 1)

		assert.DeepEqual(t, iter2.HasNext(), true)
		assert.DeepEqual(t, iter2.Next(), 1)

		assert.DeepEqual(t, iter1.HasNext(), true)
		assert.DeepEqual(t, iter1.Next(), 2)

		assert.DeepEqual(t, iter2.HasNext(), true)
		assert.DeepEqual(t, iter2.Next(), 2)

		assert.DeepEqual(t, iter1.HasNext(), true)
		assert.DeepEqual(t, iter1.Next(), 3)

		assert.DeepEqual(t, iter2.HasNext(), true)
		assert.DeepEqual(t, iter2.Next(), 3)
	})

	t.Run("collect one iterator does not affect the other", func(t *testing.T) {
		iter1, iter2 := iterator.Tee(iterator.FromSlice([]int{1, 2, 3}))
		assert.DeepEqual(t, iter1.Collect(), []int{1, 2, 3})
		assert.DeepEqual(t, iter2.Collect(), []int{1, 2, 3})
	})

	t.Run("collect from partially iterated iterator", func(t *testing.T) {
		iter1, iter2 := iterator.Tee(iterator.FromSlice([]int{1, 2, 3}))
		assert.DeepEqual(t, iter1.HasNext(), true)
		assert.DeepEqual(t, iter1.Next(), 1)

		assert.DeepEqual(t, iter1.HasNext(), true)
		assert.DeepEqual(t, iter1.Next(), 2)

		assert.DeepEqual(t, iter2.HasNext(), true)
		assert.DeepEqual(t, iter2.Next(), 1)

		assert.DeepEqual(t, iter1.Collect(), []int{3})
		assert.DeepEqual(t, iter2.Collect(), []int{2, 3})
	})

	t.Run("nested tees do not affect each other", func(t *testing.T) {
		toBeTeedAgain, iter2 := iterator.Tee(iterator.FromSlice([]int{1, 2, 3}))
		iter3, iter4 := iterator.Tee(toBeTeedAgain)

		assert.DeepEqual(t, iter2.HasNext(), true)
		assert.DeepEqual(t, iter2.Next(), 1)

		assert.DeepEqual(t, iter3.HasNext(), true)
		assert.DeepEqual(t, iter3.Next(), 1)

		assert.DeepEqual(t, iter4.HasNext(), true)
		assert.DeepEqual(t, iter4.Next(), 1)

		assert.DeepEqual(t, iter2.HasNext(), true)
		assert.DeepEqual(t, iter2.Next(), 2)

		assert.DeepEqual(t, iter3.HasNext(), true)
		assert.DeepEqual(t, iter3.Next(), 2)

		assert.DeepEqual(t, iter4.HasNext(), true)
		assert.DeepEqual(t, iter4.Next(), 2)

		assert.DeepEqual(t, iter2.HasNext(), true)
		assert.DeepEqual(t, iter2.Next(), 3)

		assert.DeepEqual(t, iter3.HasNext(), true)
		assert.DeepEqual(t, iter3.Next(), 3)

		assert.DeepEqual(t, iter4.HasNext(), true)
		assert.DeepEqual(t, iter4.Next(), 3)

		assert.DeepEqual(t, iter2.HasNext(), false)
		assert.DeepEqual(t, iter3.HasNext(), false)
		assert.DeepEqual(t, iter4.HasNext(), false)
	})
}
