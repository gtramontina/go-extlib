package iterator_test

import (
	"testing"

	"github.com/gtramontina/go-extlib/iterator"
	"github.com/gtramontina/go-extlib/testing/assert"
)

func TestSliceIterator(t *testing.T) {
	t.Run("empty slice does not HasNext", func(t *testing.T) {
		iter := iterator.FromSlice([]int{})
		assert.False(t, iter.HasNext())
	})

	t.Run("slice of one HasNext until Next is taken", func(t *testing.T) {
		iter := iterator.FromSlice([]string{"first"})
		assert.True(t, iter.HasNext())
		assert.DeepEqual(t, iter.Next(), "first")
		assert.False(t, iter.HasNext())
	})

	t.Run("every call to Next tracks whether HasNext", func(t *testing.T) {
		iter := iterator.FromSlice([]string{"first", "second", "third"})

		assert.True(t, iter.HasNext())
		assert.DeepEqual(t, iter.Next(), "first")

		assert.True(t, iter.HasNext())
		assert.DeepEqual(t, iter.Next(), "second")

		assert.True(t, iter.HasNext())
		assert.DeepEqual(t, iter.Next(), "third")

		assert.False(t, iter.HasNext())
	})

	t.Run("panics when Next is called after HasNext is false", func(t *testing.T) {
		iter := iterator.FromSlice([]string{"first", "second", "third"})
		iter.Next()
		iter.Next()
		iter.Next()
		assert.PanicsWith(t, func() { iter.Next() }, iterator.ErrIteratorEmpty)
	})
}
