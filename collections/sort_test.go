package collections_test

import (
	"math"
	"sort"
	"testing"

	"github.com/gtramontina/go-extlib/collections"
	"github.com/gtramontina/go-extlib/testing/assert"
)

func TestSort(t *testing.T) {
	t.Run("ordered types ascending", func(t *testing.T) {
		assert.DeepEqual(t, collections.Sort([]int{}), []int{})
		assert.DeepEqual(t, collections.Sort([]int{0}), []int{0})
		assert.DeepEqual(t, collections.Sort([]int{0, 1}), []int{0, 1})
		assert.DeepEqual(t, collections.Sort([]int{1, 0}), []int{0, 1})
		assert.DeepEqual(t, collections.Sort([]int{1, 2, 0}), []int{0, 1, 2})
		assert.DeepEqual(t, collections.Sort([]int{2, 10, 1}), []int{1, 2, 10})

		assert.DeepEqual(t, collections.Sort([]string{}), []string{})
		assert.DeepEqual(t, collections.Sort([]string{"a"}), []string{"a"})
		assert.DeepEqual(t, collections.Sort([]string{"a", "b"}), []string{"a", "b"})
		assert.DeepEqual(t, collections.Sort([]string{"b", "a"}), []string{"a", "b"})
		assert.DeepEqual(t, collections.Sort([]string{"b", "c", "a"}), []string{"a", "b", "c"})
		assert.DeepEqual(t, collections.Sort([]string{"2", "10", "1"}), []string{"1", "10", "2"})

		assert.DeepEqual(t, collections.Sort([]float64{}), []float64{})
		assert.DeepEqual(t, collections.Sort([]float64{0}), []float64{0})
		assert.DeepEqual(t, collections.Sort([]float64{0, 1.1}), []float64{0, 1.1})
		assert.DeepEqual(t, collections.Sort([]float64{1.1, 0}), []float64{0, 1.1})
		assert.DeepEqual(t, collections.Sort([]float64{1.1, 2.2, 0}), []float64{0, 1.1, 2.2})
		assert.DeepEqual(t, collections.Sort([]float64{2.2, 10.1, 1.1}), []float64{1.1, 2.2, 10.1})

		t.Run("does not mutate the original slice", func(t *testing.T) {
			shuffled := []int{18, 2, 3, 16, 20, 15}
			_ = collections.Sort(shuffled)
			assert.DeepEqual(t, shuffled, []int{18, 2, 3, 16, 20, 15})
		})

		t.Run("is equivalent to sort.Ints", func(t *testing.T) {
			shuffled := []int{18, 2, 3, 16, 20, 15, 19, 13, 14, 10, 9, 6, 17, 5, 11, 12, 0, 4, 7, 8, 1}
			sorted := collections.Sort(shuffled)
			sort.Ints(shuffled) // this changes the slice in place
			assert.DeepEqual(t, sorted, shuffled)
		})

		t.Run("is equivalent to sort.Strings", func(t *testing.T) {
			shuffled := []string{"e", "q", "E", "w", "1", "H", "t", "D", "l", "P", "I", "T", "7", "4", "R", "K", "L", "W", "2", "m", "F", "A", "B", "r", "s", "9", "o", "Y", "3", "V", "J", "C", "a", "f", "O", "5", "c", "d", "v", "x", "U", "X", "u", "8", "z", "M", "0", "Q", "j", "G", "6", "N", "k", "i", "Z", "p", "n", "y", "h", "S", "g", "b"}
			sorted := collections.Sort(shuffled)
			sort.Strings(shuffled) // this changes the slice in place
			assert.DeepEqual(t, sorted, shuffled)
		})

		t.Run("is equivalent to sort.Float64s", func(t *testing.T) {
			shuffled := []float64{18.8, 2.2, 3.3, 16.6, 20.0, 15.5, 19.9, 13.3, 14.4, 10.0, 9.9, 6.6, 17.7, 5.5, 11.1, 12.2, 0.0, 4.4, 7.7, 8.8, math.NaN(), 1.1}
			sorted := collections.Sort(shuffled)
			sort.Float64s(shuffled) // this changes the slice in place
			assert.True(t, math.IsNaN(sorted[0]))
			assert.True(t, math.IsNaN(shuffled[0]))
			assert.DeepEqual(t, sorted[1:], shuffled[1:])
		})
	})

	t.Run("ordered types descending", func(t *testing.T) {
		assert.DeepEqual(t, collections.SortDescending([]int{}), []int{})
		assert.DeepEqual(t, collections.SortDescending([]int{0}), []int{0})
		assert.DeepEqual(t, collections.SortDescending([]int{0, 1}), []int{1, 0})
		assert.DeepEqual(t, collections.SortDescending([]int{1, 0}), []int{1, 0})
		assert.DeepEqual(t, collections.SortDescending([]int{1, 2, 0}), []int{2, 1, 0})
		assert.DeepEqual(t, collections.SortDescending([]int{2, 10, 1}), []int{10, 2, 1})

		assert.DeepEqual(t, collections.SortDescending([]string{}), []string{})
		assert.DeepEqual(t, collections.SortDescending([]string{"a"}), []string{"a"})
		assert.DeepEqual(t, collections.SortDescending([]string{"a", "b"}), []string{"b", "a"})
		assert.DeepEqual(t, collections.SortDescending([]string{"b", "a"}), []string{"b", "a"})
		assert.DeepEqual(t, collections.SortDescending([]string{"b", "c", "a"}), []string{"c", "b", "a"})
		assert.DeepEqual(t, collections.SortDescending([]string{"2", "10", "1"}), []string{"2", "10", "1"})

		assert.DeepEqual(t, collections.SortDescending([]float64{}), []float64{})
		assert.DeepEqual(t, collections.SortDescending([]float64{0}), []float64{0})
		assert.DeepEqual(t, collections.SortDescending([]float64{0, 1.1}), []float64{1.1, 0})
		assert.DeepEqual(t, collections.SortDescending([]float64{1.1, 0}), []float64{1.1, 0})
		assert.DeepEqual(t, collections.SortDescending([]float64{1.1, 2.2, 0}), []float64{2.2, 1.1, 0})
		assert.DeepEqual(t, collections.SortDescending([]float64{2.2, 10.1, 1.1}), []float64{10.1, 2.2, 1.1})

		sorted := collections.SortDescending([]float64{9.9, 6.6, 17.7, 5.5, 11.1, 12.2, 0.0, 4.4, 7.7, 8.8, math.NaN(), 1.1})
		assert.True(t, math.IsNaN(sorted[len(sorted)-1]))
		assert.DeepEqual(t, sorted[:len(sorted)-1], []float64{17.7, 12.2, 11.1, 9.9, 8.8, 7.7, 6.6, 5.5, 4.4, 1.1, 0.0})

		t.Run("does not mutate the original slice", func(t *testing.T) {
			shuffled := []int{18, 2, 3, 16, 20, 15}
			_ = collections.SortDescending(shuffled)
			assert.DeepEqual(t, shuffled, []int{18, 2, 3, 16, 20, 15})
		})
	})

	t.Run("non ordered types", func(t *testing.T) {
		type person struct {
			name string
			age  int
		}

		byName := func(i, j person) bool { return i.name < j.name }
		byAge := func(i, j person) bool { return i.age < j.age }

		assert.DeepEqual(t, collections.SortBy([]person{}, byName), []person{})
		assert.DeepEqual(t, collections.SortBy([]person{}, byAge), []person{})

		assert.DeepEqual(t, collections.SortBy([]person{
			{name: "Jane", age: 10},
		}, byName), []person{
			{name: "Jane", age: 10},
		})
		assert.DeepEqual(t, collections.SortBy([]person{
			{name: "Jane", age: 10},
		}, byAge), []person{
			{name: "Jane", age: 10},
		})

		assert.DeepEqual(t, collections.SortBy([]person{
			{name: "Jane", age: 10},
			{name: "Carl", age: 12},
		}, byName), []person{
			{name: "Carl", age: 12},
			{name: "Jane", age: 10},
		})
		assert.DeepEqual(t, collections.SortBy([]person{
			{name: "Jane", age: 10},
			{name: "Carl", age: 12},
		}, byAge), []person{
			{name: "Jane", age: 10},
			{name: "Carl", age: 12},
		})

		assert.DeepEqual(t, collections.SortBy([]person{
			{name: "Jane", age: 10},
			{name: "Linda", age: 8},
			{name: "Carl", age: 12},
		}, byName), []person{
			{name: "Carl", age: 12},
			{name: "Jane", age: 10},
			{name: "Linda", age: 8},
		})
		assert.DeepEqual(t, collections.SortBy([]person{
			{name: "Jane", age: 10},
			{name: "Linda", age: 8},
			{name: "Carl", age: 12},
		}, byAge), []person{
			{name: "Linda", age: 8},
			{name: "Jane", age: 10},
			{name: "Carl", age: 12},
		})

		t.Run("does not mutate the original slice", func(t *testing.T) {
			shuffled := []int{18, 2, 3, 16, 20, 15}
			_ = collections.SortBy(shuffled, func(i, j int) bool { return i < j })
			assert.DeepEqual(t, shuffled, []int{18, 2, 3, 16, 20, 15})
		})
	})
}
