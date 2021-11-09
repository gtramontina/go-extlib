package collections_test

import (
	"fmt"
	"testing"

	"github.com/gtramontina/go-extlib/collections"
	"github.com/gtramontina/go-extlib/internal/assert"
)

func TestFoldLeft(t *testing.T) {
	t.Run("is left associative", func(t *testing.T) {
		f := func(a, b int) int { return a - b }
		assert.Eq(t, collections.FoldLeft([]int{}, f, 9), 9)            // (9)
		assert.Eq(t, collections.FoldLeft([]int{1}, f, 9), 8)           // (9 - 1)
		assert.Eq(t, collections.FoldLeft([]int{1, 2}, f, 9), 6)        // ((9 - 1) - 2)
		assert.Eq(t, collections.FoldLeft([]int{1, 2, 3}, f, 9), 3)     // (((9 - 1) - 2) - 3)
		assert.Eq(t, collections.FoldLeft([]int{1, 2, 3, 4}, f, 9), -1) // ((((9 - 1) - 2) - 3) - 4)
		assert.Eq(t, collections.FoldLeft([]int{1, 2, 3, 4}, f, 9), f(f(f(f(9, 1), 2), 3), 4))
	})

	t.Run("can fold into a different type", func(t *testing.T) {
		f := func(a string, b int) string { return fmt.Sprintf("(%s + %d)", a, b) }
		assert.Eq(t, collections.FoldLeft([]int{1, 2, 3}, f, "0"), "(((0 + 1) + 2) + 3)")
	})
}

func TestFoldRight(t *testing.T) {
	t.Run("is right associative", func(t *testing.T) {
		f := func(a, b int) int { return a - b }
		assert.Eq(t, collections.FoldRight([]int{}, f, 9), 9)           // (9)
		assert.Eq(t, collections.FoldRight([]int{1}, f, 9), -8)         // (1 - 9)
		assert.Eq(t, collections.FoldRight([]int{1, 2}, f, 9), 8)       // (1 - (2 - 9))
		assert.Eq(t, collections.FoldRight([]int{1, 2, 3}, f, 9), -7)   // (1 - (2 - (3 - 9)))
		assert.Eq(t, collections.FoldRight([]int{1, 2, 3, 4}, f, 9), 7) // (1 - (2 - (3 - (4 - 9))))
		assert.Eq(t, collections.FoldRight([]int{1, 2, 3, 4}, f, 9), f(1, f(2, f(3, f(4, 9)))))
	})

	t.Run("can fold into a different type", func(t *testing.T) {
		f := func(a int, b string) string { return fmt.Sprintf("(%d + %s)", a, b) }
		assert.Eq(t, collections.FoldRight([]int{1, 2, 3}, f, "0"), "(1 + (2 + (3 + 0)))")
	})
}
