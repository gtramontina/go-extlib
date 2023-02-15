package collections_test

import (
	"testing"

	"github.com/gtramontina/go-extlib/collections"
	"github.com/gtramontina/go-extlib/testing/assert"
)

func TestChunk(t *testing.T) {
	t.Run("panics when the number of chunks given is less than 1", func(t *testing.T) {
		assert.PanicsWith(t, func() { collections.Chunk([]int{0}, -1) }, "chunk size must be greater than 1")
		assert.PanicsWith(t, func() { collections.Chunk([]int{0}, 0) }, "chunk size must be greater than 1")
	})

	t.Run("when empty collection, returns an empty collection of chunks", func(t *testing.T) {
		assert.DeepEqual(t, collections.Chunk([]int{}, 1), [][]int{})
	})

	t.Run("when the chunk size equal to the number of elements returns a collection with a single chunk", func(t *testing.T) {
		assert.DeepEqual(t, collections.Chunk([]int{0}, 1), [][]int{{0}})
		assert.DeepEqual(t, collections.Chunk([]int{0, 1}, 2), [][]int{{0, 1}})
		assert.DeepEqual(t, collections.Chunk([]int{0, 1, 2}, 3), [][]int{{0, 1, 2}})
	})

	t.Run("chunks evenly", func(t *testing.T) {
		assert.DeepEqual(t, collections.Chunk([]int{0, 1}, 1), [][]int{{0}, {1}})
		assert.DeepEqual(t, collections.Chunk([]int{0, 1, 2}, 1), [][]int{{0}, {1}, {2}})
		assert.DeepEqual(t, collections.Chunk([]int{0, 1}, 2), [][]int{{0, 1}})
		assert.DeepEqual(t, collections.Chunk([]int{0, 1, 2, 3}, 2), [][]int{{0, 1}, {2, 3}})
		assert.DeepEqual(t, collections.Chunk([]int{0, 1, 2, 3, 4, 5}, 2), [][]int{{0, 1}, {2, 3}, {4, 5}})
		assert.DeepEqual(t, collections.Chunk([]int{0, 1, 2, 3, 4, 5}, 3), [][]int{{0, 1, 2}, {3, 4, 5}})
	})

	t.Run("chunks unevenly", func(t *testing.T) {
		assert.DeepEqual(t, collections.Chunk([]int{0, 1, 2}, 2), [][]int{{0, 1}, {2}})
		assert.DeepEqual(t, collections.Chunk([]int{0, 1, 2, 3, 4}, 2), [][]int{{0, 1}, {2, 3}, {4}})
		assert.DeepEqual(t, collections.Chunk([]int{0, 1, 2, 3}, 3), [][]int{{0, 1, 2}, {3}})
		assert.DeepEqual(t, collections.Chunk([]int{0, 1, 2, 3, 4, 5, 6}, 3), [][]int{{0, 1, 2}, {3, 4, 5}, {6}})
	})
}
