package iterator_test

import (
	"testing"

	"github.com/gtramontina/go-extlib/iterator"
	"github.com/gtramontina/go-extlib/testing/assert"
)

func TestSplit(t *testing.T) {
	t.Run("empty iterator", func(t *testing.T) {
		left, right := iterator.Split(iterator.From[int](), func(int) bool { return true })
		assert.False(t, left.HasNext())
		assert.False(t, right.HasNext())
	})

	t.Run("yields an iterator with all items and an empty iterator if the predicate is true for all items", func(t *testing.T) {
		left, right := iterator.Split(iterator.From('a'), func(r rune) bool { return true })
		assert.True(t, left.HasNext())
		assert.DeepEqual(t, left.Next(), 'a')
		assert.False(t, left.HasNext())
		assert.False(t, right.HasNext())
	})

	t.Run("yields an empty iterator and an iterator with all items if the predicate is false for all items", func(t *testing.T) {
		left, right := iterator.Split(iterator.From('a'), func(r rune) bool { return false })
		assert.False(t, left.HasNext())
		assert.True(t, right.HasNext())
		assert.DeepEqual(t, right.Next(), 'a')
		assert.False(t, right.HasNext())
	})

	t.Run("yields an iterator with the items that match the predicate and an iterator with the items that don't", func(t *testing.T) {
		evens, odds := iterator.Split(iterator.From[int](1, 2, 3, 4, 5), func(i int) bool { return i%2 == 0 })
		assert.True(t, evens.HasNext())
		assert.DeepEqual(t, evens.Next(), 2)
		assert.True(t, evens.HasNext())
		assert.DeepEqual(t, evens.Next(), 4)
		assert.False(t, evens.HasNext())
		assert.True(t, odds.HasNext())
		assert.DeepEqual(t, odds.Next(), 1)
		assert.True(t, odds.HasNext())
		assert.DeepEqual(t, odds.Next(), 3)
		assert.True(t, odds.HasNext())
		assert.DeepEqual(t, odds.Next(), 5)
		assert.False(t, odds.HasNext())
	})
}
