package collections_test

import (
	"github.com/gtramontina/go-extlib/collections"
	"github.com/gtramontina/go-extlib/internal/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	t.Run("empty collection yields an empty collection", func(t *testing.T) {
		dummy := func(int) bool { return true }
		assert.DeepEqual(t, collections.Filter([]int{}, dummy), []int{})
	})

	t.Run("filters all out", func(t *testing.T) {
		allOut := func(int) bool { return false }
		assert.DeepEqual(t, collections.Filter([]int{1}, allOut), []int{})
		assert.DeepEqual(t, collections.Filter([]int{1, 2}, allOut), []int{})
	})

	t.Run("filters all in", func(t *testing.T) {
		allIn := func(int) bool { return true }
		assert.DeepEqual(t, collections.Filter([]int{1}, allIn), []int{1})
		assert.DeepEqual(t, collections.Filter([]int{1, 2}, allIn), []int{1, 2})
	})

	t.Run("conditionally filters elements", func(t *testing.T) {
		even := func(it int) bool { return it%2 == 0 }
		odd := func(it int) bool { return !even(it) }
		assert.DeepEqual(t, collections.Filter([]int{1, 2, 3, 4}, even), []int{2, 4})
		assert.DeepEqual(t, collections.Filter([]int{1, 2, 3, 4}, odd), []int{1, 3})
	})

	t.Run("handles different types", func(t *testing.T) {
		notEmpty := func(it string) bool { return len(it) > 0 }
		assert.DeepEqual(t, collections.Filter([]string{"a", "", "c"}, notEmpty), []string{"a", "c"})

		type person struct {
			name string
			age  uint
		}
		legalAge := func(it person) bool { return it.age >= 18 }
		assert.DeepEqual(t, collections.Filter([]person{
			{"Jane", 21},
			{"Jack", 15},
			{"John", 18},
		}, legalAge), []person{
			{"Jane", 21},
			{"John", 18},
		})
	})
}
