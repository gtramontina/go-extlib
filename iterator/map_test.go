package iterator_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gtramontina/go-extlib/iterator"
	"github.com/gtramontina/go-extlib/testing/assert"
)

func Empty[T any]() iterator.Iterator[T] {
	return iterator.FromSlice([]T{})
}

func TestMap(t *testing.T) {
	t.Run("empty collection yields empty collection", func(t *testing.T) {
		dummy := func(it int) int { return it }
		assert.DeepEqual(t, iterator.Map(Empty[int](), dummy).Collect(), []int{})
	})

	t.Run("maps all items", func(t *testing.T) {
		identity := func(it int) int { return it }
		assert.DeepEqual(t, iterator.Map(iterator.FromSlice([]int{1}), identity).Collect(), []int{1})
		assert.DeepEqual(t, iterator.Map(iterator.FromSlice([]int{1, 2}), identity).Collect(), []int{1, 2})

		add1 := func(it int) int { return it + 1 }
		assert.DeepEqual(t, iterator.Map(iterator.FromSlice([]int{1, 2}), add1).Collect(), []int{2, 3})
	})

	t.Run("maps to different types", func(t *testing.T) {
		toString := func(it int) string { return fmt.Sprintf("%d", it) }
		assert.DeepEqual(t, iterator.Map(iterator.FromSlice([]int{1, 2}), toString).Collect(), []string{"1", "2"})

		type data struct{ today time.Time }
		type view struct{ today string }
		toView := func(it data) view { return view{it.today.Format(time.Stamp)} }
		assert.DeepEqual(t, iterator.Map(iterator.FromSlice([]data{
			{today: time.Date(1985, time.December, 1, 0, 0, 0, 0, time.UTC)},
			{today: time.Date(2021, time.November, 20, 15, 10, 5, 0, time.UTC)},
		}), toView).Collect(), []view{
			{today: "Dec  1 00:00:00"},
			{today: "Nov 20 15:10:05"},
		})
	})
}
