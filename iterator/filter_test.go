package iterator_test

import (
	"testing"

	"github.com/gtramontina/go-extlib/iterator"
	"github.com/gtramontina/go-extlib/testing/assert"
)

func TestFilter(t *testing.T) {
	t.Run("an empty iterator is empty", func(t *testing.T) {
		iter := iterator.Filter(iterator.From[int](), func(int) bool { return true })
		assert.False(t, iter.HasNext())
	})

	t.Run("yields an empty iterator if the predicate is false for all items", func(t *testing.T) {
		iter := iterator.Filter(iterator.From('a'), func(r rune) bool { return false })
		assert.False(t, iter.HasNext())
	})

	t.Run("yields an iterator with the same items if the predicate is true for all items", func(t *testing.T) {
		iter := iterator.Filter(iterator.From('a'), func(r rune) bool { return true })
		assert.True(t, iter.HasNext())
		assert.DeepEqual(t, iter.Next(), 'a')
		assert.False(t, iter.HasNext())
	})

	t.Run("yields only items that match the predicate", func(t *testing.T) {
		iter := iterator.Filter(iterator.From[int](1, 2, 3, 4, 5), func(i int) bool { return i%2 == 0 })
		assert.True(t, iter.HasNext())
		assert.DeepEqual(t, iter.Next(), 2)
		assert.True(t, iter.HasNext())
		assert.DeepEqual(t, iter.Next(), 4)
		assert.False(t, iter.HasNext())
	})
}
