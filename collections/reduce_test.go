package collections_test

import (
	"testing"

	"github.com/gtramontina/go-extlib/collections"
	"github.com/gtramontina/go-extlib/internal/assert"
)

func TestReduceLeft(t *testing.T) {
	t.Run("panics when the given collection is empty", func(t *testing.T) {
		assert.Panic(t, func() { collections.ReduceLeft([]int{}, func(acc, val int) int { return acc + val }) }, "cannot ReduceLeft an empty slice")
		assert.Panic(t, func() { collections.ReduceLeft([]string{}, func(acc, val string) string { return acc + val }) }, "cannot ReduceLeft an empty slice")
	})

	t.Run("is left associative", func(t *testing.T) {
		f := func(a, b int) int { return a - b }
		assert.Eq(t, collections.ReduceLeft([]int{1}, f), 1)           // (1)
		assert.Eq(t, collections.ReduceLeft([]int{1, 2}, f), -1)       // (1 - 2)
		assert.Eq(t, collections.ReduceLeft([]int{1, 2, 3}, f), -4)    // ((1 - 2) - 3)
		assert.Eq(t, collections.ReduceLeft([]int{1, 2, 3, 4}, f), -8) // (((1 - 2) - 3) - 4)
		assert.Eq(t, collections.ReduceLeft([]int{1, 2, 3, 4}, f), f(f(f(1, 2), 3), 4))
	})
}

func TestReduceRight(t *testing.T) {
	t.Run("panics when the given collection is empty", func(t *testing.T) {
		assert.Panic(t, func() { collections.ReduceRight([]int{}, func(acc, val int) int { return acc + val }) }, "cannot ReduceRight an empty slice")
		assert.Panic(t, func() { collections.ReduceRight([]string{}, func(acc, val string) string { return acc + val }) }, "cannot ReduceRight an empty slice")
	})

	t.Run("is right associative", func(t *testing.T) {
		f := func(a, b int) int { return a - b }
		assert.Eq(t, collections.ReduceRight([]int{1}, f), 1)           // (1)
		assert.Eq(t, collections.ReduceRight([]int{1, 2}, f), -1)       // (1 - 2)
		assert.Eq(t, collections.ReduceRight([]int{1, 2, 3}, f), 2)     // (1 - (2 - 3))
		assert.Eq(t, collections.ReduceRight([]int{1, 2, 3, 4}, f), -2) // (1 - (2 - (3 - 4)))
		assert.Eq(t, collections.ReduceRight([]int{1, 2, 3, 4}, f), f(1, f(2, f(3, 4))))
	})
}
