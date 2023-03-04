package collections_test

import (
	"testing"

	"github.com/gtramontina/go-extlib/collections"
	"github.com/gtramontina/go-extlib/hashmap"
	"github.com/gtramontina/go-extlib/testing/assert"
)

func TestGroupBy(t *testing.T) {
	t.Run("empty slice returns an empty hashmap", func(t *testing.T) {
		initialLetter := func(s string) rune { return rune(s[0]) }
		assert.DeepEqual(t, collections.GroupBy([]string{}, initialLetter), hashmap.New[rune, []string]())
	})

	t.Run("slice with one entry returns hashmap of groups by the given function", func(t *testing.T) {
		initialLetter := func(s string) rune { return rune(s[0]) }
		assert.DeepEqual(t, collections.GroupBy([]string{"a"}, initialLetter), hashmap.New[rune, []string]().
			Put('a', []string{"a"}),
		)
	})

	t.Run("slice with multiple entries returns hashmap of groups by the given function", func(t *testing.T) {
		initialLetter := func(s string) rune { return rune(s[0]) }
		assert.DeepEqual(t, collections.GroupBy([]string{"a", "b", "c"}, initialLetter), hashmap.New[rune, []string]().
			Put('a', []string{"a"}).
			Put('b', []string{"b"}).
			Put('c', []string{"c"}),
		)

		assert.DeepEqual(t, collections.GroupBy([]string{"hello", "world", "hi", "there"}, initialLetter), hashmap.New[rune, []string]().
			Put('h', []string{"hello", "hi"}).
			Put('w', []string{"world"}).
			Put('t', []string{"there"}),
		)
	})

	t.Run("with more complex types", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		ageGroup := func(p Person) int { return p.Age / 10 }
		assert.DeepEqual(t, collections.GroupBy([]Person{
			{Name: "John", Age: 20},
			{Name: "Jane", Age: 25},
			{Name: "Joe", Age: 32},
			{Name: "Jill", Age: 46},
		}, ageGroup), hashmap.New[int, []Person]().
			Put(2, []Person{{Name: "John", Age: 20}, {Name: "Jane", Age: 25}}).
			Put(3, []Person{{Name: "Joe", Age: 32}}).
			Put(4, []Person{{Name: "Jill", Age: 46}}),
		)
	})
}
