package iterator_test

import (
	"testing"

	"github.com/gtramontina/go-extlib/iterator"
	"github.com/gtramontina/go-extlib/testing/assert"
)

func TestCycle(t *testing.T) {
	t.Run("cycle over an empty slice", func(t *testing.T) {
		iter := iterator.Cycle[int]()
		assert.DeepEqual(t, iter.HasNext(), false)
		assert.PanicsWith(t, func() { iter.Next() }, iterator.ErrIteratorEmpty)
	})

	t.Run("cycle over a single element slice", func(t *testing.T) {
		iter := iterator.Cycle[int](1)
		assert.DeepEqual(t, iter.HasNext(), true)
		assert.DeepEqual(t, iter.Next(), 1)
		assert.DeepEqual(t, iter.HasNext(), true)
		assert.DeepEqual(t, iter.Next(), 1)
		assert.DeepEqual(t, iter.HasNext(), true)
		assert.DeepEqual(t, iter.Next(), 1)
		assert.DeepEqual(t, iter.HasNext(), true)
		assert.DeepEqual(t, iter.Next(), 1)
		assert.DeepEqual(t, iter.HasNext(), true)
		assert.DeepEqual(t, iter.Next(), 1)
	})

	t.Run("cycle over a multiple element slice", func(t *testing.T) {
		iter := iterator.Cycle[int](1, 2, 3)
		assert.DeepEqual(t, iter.HasNext(), true)
		assert.DeepEqual(t, iter.Next(), 1)
		assert.DeepEqual(t, iter.HasNext(), true)
		assert.DeepEqual(t, iter.Next(), 2)
		assert.DeepEqual(t, iter.HasNext(), true)
		assert.DeepEqual(t, iter.Next(), 3)
		assert.DeepEqual(t, iter.HasNext(), true)
		assert.DeepEqual(t, iter.Next(), 1)
		assert.DeepEqual(t, iter.HasNext(), true)
		assert.DeepEqual(t, iter.Next(), 2)
		assert.DeepEqual(t, iter.HasNext(), true)
		assert.DeepEqual(t, iter.Next(), 3)
		assert.DeepEqual(t, iter.HasNext(), true)
		assert.DeepEqual(t, iter.Next(), 1)
		assert.DeepEqual(t, iter.HasNext(), true)
		assert.DeepEqual(t, iter.Next(), 2)
		assert.DeepEqual(t, iter.HasNext(), true)
		assert.DeepEqual(t, iter.Next(), 3)
		assert.DeepEqual(t, iter.HasNext(), true)
		assert.DeepEqual(t, iter.Next(), 1)
		assert.DeepEqual(t, iter.HasNext(), true)
	})
}
