package set_test

import (
	"testing"

	"github.com/gtramontina/go-extlib/internal/assert"
	"github.com/gtramontina/go-extlib/set"
)

func TestSet(t *testing.T) {
	t.Run("empty sets never contain any members", func(t *testing.T) {
		assert.False(t, set.New[int]().Contains(0))
		assert.False(t, set.New[int]().Contains(1))
	})

	t.Run("is comparable to other sets", func(t *testing.T) {
		assert.True(t, set.New[int]().Equals(set.New[int]()))
		assert.False(t, set.New(0).Equals(set.New[int]()))
		assert.True(t, set.New(0).Equals(set.New(0)))
		assert.False(t, set.New(0).Equals(set.New(1)))
		assert.False(t, set.New(0).Equals(set.New(0, 1)))
		assert.False(t, set.New(0, 1).Equals(set.New(0)))
		assert.True(t, set.New(0, 1).Equals(set.New(1, 0)))
	})

	t.Run("filled sets may contain members", func(t *testing.T) {
		assert.True(t, set.New(0).Contains(0))
		assert.False(t, set.New(0).Contains(1))
		assert.True(t, set.New(0, 1).Contains(0))
		assert.True(t, set.New(0, 1).Contains(1))
		assert.False(t, set.New(0, 1).Contains(2))
	})

	t.Run("filled sets may be supersets of other sets", func(t *testing.T) {
		assert.False(t, set.New[int]().SuperSetOf(set.New(0)))
		assert.True(t, set.New(0).SuperSetOf(set.New(0)))
		assert.False(t, set.New(0).SuperSetOf(set.New(1)))
		assert.True(t, set.New(0, 1).SuperSetOf(set.New(0)))
		assert.True(t, set.New(0, 1).SuperSetOf(set.New(1)))
		assert.False(t, set.New(0, 1).SuperSetOf(set.New(2)))
		assert.True(t, set.New(0, 1, 2).SuperSetOf(set.New(0, 1)))
		assert.True(t, set.New(0, 1, 2).SuperSetOf(set.New(0, 2)))
		assert.True(t, set.New(0, 1, 2).SuperSetOf(set.New(1, 2)))
	})

	t.Run("keeps track of its cardinality", func(t *testing.T) {
		assert.Eq(t, set.New[int]().Cardinality(), 0)
		assert.Eq(t, set.New(0).Cardinality(), 1)
		assert.Eq(t, set.New(0, 1).Cardinality(), 2)
	})

	t.Run("ignores duplicate members", func(t *testing.T) {
		assert.Equals(t, set.New(0, 0), set.New(0))
		assert.Eq(t, set.New(0, 0).Cardinality(), 1)
		assert.Equals(t, set.New(0, 1, 1), set.New(0, 1))
		assert.Eq(t, set.New(0, 1, 1).Cardinality(), 2)
	})

	t.Run("can add new members", func(t *testing.T) {
		assert.Equals(t, set.New[int]().Add(0), set.New(0))
		assert.Equals(t, set.New[int]().Add(0).Add(1), set.New(0, 1))
		assert.Equals(t, set.New[int]().Add(0).Add(1).Add(2), set.New(0, 1, 2))

		t.Run("does not mutate the sets", func(t *testing.T) {
			empty := set.New[int]()
			oneMember := empty.Add(0)
			twoMembers := oneMember.Add(1)
			assert.NotEquals(t, empty, oneMember)
			assert.NotEquals(t, empty, twoMembers)
			assert.NotEquals(t, oneMember, twoMembers)
		})
	})

	t.Run("can remove members", func(t *testing.T) {
		assert.Equals(t, set.New(0).Remove(1), set.New(0))
		assert.Equals(t, set.New(0).Remove(0), set.New[int]())
		assert.Equals(t, set.New(0, 1).Remove(0), set.New(1))

		t.Run("does not mutate the sets", func(t *testing.T) {
			twoMembers := set.New(0, 1)
			oneMember := twoMembers.Remove(1)
			empty := oneMember.Remove(0)
			assert.NotEquals(t, empty, oneMember)
			assert.NotEquals(t, empty, twoMembers)
			assert.NotEquals(t, oneMember, twoMembers)
		})
	})

	t.Run("can perform union of sets", func(t *testing.T) {
		assert.Equals(t, set.New[int]().Union(set.New[int]()), set.New[int]())
		assert.Equals(t, set.New(0).Union(set.New[int]()), set.New(0))
		assert.Equals(t, set.New[int]().Union(set.New(0)), set.New(0))
		assert.Equals(t, set.New(0).Union(set.New(0)), set.New(0))
		assert.Equals(t, set.New(0).Union(set.New(1)), set.New(0, 1))
		assert.Equals(t, set.New(0, 1).Union(set.New(1, 2)), set.New(0, 1, 2))
		assert.Equals(t, set.New(1, 2).Union(set.New(0, 1)), set.New(0, 1, 2))

		t.Run("does not mutate the sets", func(t *testing.T) {
			setA := set.New(0, 1)
			setB := set.New(1, 2)
			_ = setA.Union(setB)
			assert.Equals(t, setA, set.New(0, 1))
			assert.Equals(t, setB, set.New(1, 2))
		})
	})

	t.Run("can get intersection of sets", func(t *testing.T) {
		assert.Equals(t, set.New[int]().Intersection(set.New[int]()), set.New[int]())
		assert.Equals(t, set.New(0).Intersection(set.New[int]()), set.New[int]())
		assert.Equals(t, set.New[int]().Intersection(set.New(0)), set.New[int]())
		assert.Equals(t, set.New(0).Intersection(set.New(0)), set.New(0))
		assert.Equals(t, set.New(0).Intersection(set.New(1)), set.New[int]())
		assert.Equals(t, set.New(0, 1).Intersection(set.New(1, 2)), set.New(1))
		assert.Equals(t, set.New(1, 2).Intersection(set.New(0, 1)), set.New(1))
		assert.Equals(t, set.New(0, 1, 2).Intersection(set.New(1, 2, 3)), set.New(1, 2))

		t.Run("does not mutate the sets", func(t *testing.T) {
			setA := set.New(0, 1)
			setB := set.New(1, 2)
			_ = setA.Intersection(setB)
			assert.Equals(t, setA, set.New(0, 1))
			assert.Equals(t, setB, set.New(1, 2))
		})
	})

	t.Run("can get difference of sets", func(t *testing.T) {
		assert.Equals(t, set.New[int]().Difference(set.New[int]()), set.New[int]())
		assert.Equals(t, set.New(0).Difference(set.New[int]()), set.New(0))
		assert.Equals(t, set.New[int]().Difference(set.New(0)), set.New[int]())
		assert.Equals(t, set.New(0).Difference(set.New(0)), set.New[int]())
		assert.Equals(t, set.New(0).Difference(set.New(1)), set.New(0))
		assert.Equals(t, set.New(0, 1).Difference(set.New(1, 2)), set.New(0))
		assert.Equals(t, set.New(1, 2).Difference(set.New(0, 1)), set.New(2))
		assert.Equals(t, set.New(0, 1, 2).Difference(set.New(1, 2, 3)), set.New(0))

		t.Run("does not mutate the sets", func(t *testing.T) {
			setA := set.New(0, 1)
			setB := set.New(1, 2)
			_ = setA.Difference(setB)
			assert.Equals(t, setA, set.New(0, 1))
			assert.Equals(t, setB, set.New(1, 2))
		})
	})

	t.Run("can get symmetric difference of sets", func(t *testing.T) {
		assert.Equals(t, set.New[int]().SymmetricDifference(set.New[int]()), set.New[int]())
		assert.Equals(t, set.New(0).SymmetricDifference(set.New[int]()), set.New(0))
		assert.Equals(t, set.New[int]().SymmetricDifference(set.New(0)), set.New(0))
		assert.Equals(t, set.New(0).SymmetricDifference(set.New(0)), set.New[int]())
		assert.Equals(t, set.New(0).SymmetricDifference(set.New(1)), set.New(0, 1))
		assert.Equals(t, set.New(0, 1).SymmetricDifference(set.New(1, 2)), set.New(0, 2))
		assert.Equals(t, set.New(1, 2).SymmetricDifference(set.New(0, 1)), set.New(0, 2))
		assert.Equals(t, set.New(0, 1, 2).SymmetricDifference(set.New(1, 2, 3)), set.New(0, 3))

		t.Run("does not mutate the sets", func(t *testing.T) {
			setA := set.New(0, 1)
			setB := set.New(1, 2)
			_ = setA.SymmetricDifference(setB)
			assert.Equals(t, setA, set.New(0, 1))
			assert.Equals(t, setB, set.New(1, 2))
		})
	})

	t.Run("allows filtering sets", func(t *testing.T) {
		passthrough := func(int) bool { return true }
		even := func(it int) bool { return it%2 == 0 }
		odd := func(it int) bool { return !even(it) }

		assert.Equals(t, set.New[int]().Filter(passthrough), set.New[int]())
		assert.Equals(t, set.New(0).Filter(passthrough), set.New(0))
		assert.Equals(t, set.New(1).Filter(passthrough), set.New(1))
		assert.Equals(t, set.New(1, 2, 3, 4).Filter(even), set.New(2, 4))
		assert.Equals(t, set.New(1, 2, 3, 4).Filter(odd), set.New(1, 3))

		t.Run("does not mutate the sets", func(t *testing.T) {
			setA := set.New(0, 1)
			_ = setA.Filter(passthrough)
			assert.Equals(t, setA, set.New(0, 1))
		})
	})

	type person struct{ name string }
	t.Run("handles different types", func(t *testing.T) {
		assert.Equals(t, set.New("a"), set.New("a"))
		assert.NotEquals(t, set.New("a"), set.New("b"))
		assert.Equals(t, set.New(true), set.New(true))
		assert.NotEquals(t, set.New(true), set.New(false))
		assert.Equals(t, set.New(person{"Jane Doe"}), set.New(person{"Jane Doe"}))
		assert.NotEquals(t, set.New(person{"Jane Doe"}), set.New(person{"John Doe"}))
	})

	t.Run("renders itself as string", func(t *testing.T) {
		assert.Eq(t, set.New[int]().String(), "Set(int){}")
		assert.Eq(t, set.New(0).String(), "Set(int){0}")
		assert.Eq(t, set.New(1).String(), "Set(int){1}")
		assert.Eq(t, set.New(0, 1).String(), "Set(int){0, 1}")
		assert.Eq(t, set.New(1, 0).String(), "Set(int){0, 1}")
		assert.Eq(t, set.New("0", "1").String(), `Set(string){0, 1}`)
		assert.Eq(t, set.New("B", "A").String(), `Set(string){A, B}`)
		assert.Eq(t, set.New(person{"Jane"}, person{"John"}).String(), `Set(set_test.person){{name:Jane}, {name:John}}`)
	})
}
