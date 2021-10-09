package collections_test

import (
	"fmt"
	"github.com/gtramontina/go-collections/collections"
	"github.com/gtramontina/go-collections/internal/assert"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	t.Run("empty collection yields empty collection", func(t *testing.T) {
		dummy := func(it int) int { return it }
		assert.DeepEqual(t, collections.Map([]int{}, dummy), []int{})
	})

	t.Run("maps all items", func(t *testing.T) {
		identity := func(it int) int { return it }
		assert.DeepEqual(t, collections.Map([]int{1}, identity), []int{1})
		assert.DeepEqual(t, collections.Map([]int{1, 2}, identity), []int{1, 2})

		add1 := func(it int) int { return it + 1 }
		assert.DeepEqual(t, collections.Map([]int{1, 2}, add1), []int{2, 3})
	})

	t.Run("maps to different types", func(t *testing.T) {
		toString := func(it int) string { return fmt.Sprintf("%d", it) }
		assert.DeepEqual(t, collections.Map([]int{1, 2}, toString), []string{"1", "2"})

		type data struct{ today time.Time }
		type view struct{ today string }
		toView := func(it data) view { return view{it.today.Format(time.Stamp)} }
		assert.DeepEqual(t, collections.Map([]data{
			{today: time.Date(1985, time.December, 1, 0, 0, 0, 0, time.UTC)},
			{today: time.Date(2021, time.November, 20, 15, 10, 5, 0, time.UTC)},
		}, toView), []view{
			{today: "Dec  1 00:00:00"},
			{today: "Nov 20 15:10:05"},
		})
	})
}
