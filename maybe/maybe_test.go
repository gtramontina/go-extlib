package maybe_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/gtramontina/go-extlib/maybe"
	"github.com/gtramontina/go-extlib/testing/assert"
)

func TestMaybe(t *testing.T) {
	mustParseInt := func(in string) int {
		out, _ := strconv.ParseInt(in, 10, 0)

		return int(out)
	}

	type sample struct{ value int }

	t.Run("when type checking", func(t *testing.T) {
		t.Run("Some is always Some", func(t *testing.T) {
			assert.True(t, maybe.Some(1).IsSome())
			assert.True(t, maybe.Some("value").IsSome())
			assert.True(t, maybe.Some(sample{1}).IsSome())
		})

		t.Run("Some is never None", func(t *testing.T) {
			assert.False(t, maybe.Some(1).IsNone())
			assert.False(t, maybe.Some("value").IsNone())
			assert.False(t, maybe.Some(sample{1}).IsNone())
		})

		t.Run("None is always None", func(t *testing.T) {
			assert.True(t, maybe.None[int]().IsNone())
			assert.True(t, maybe.None[string]().IsNone())
			assert.True(t, maybe.None[sample]().IsNone())
		})

		t.Run("None is never Some", func(t *testing.T) {
			assert.False(t, maybe.None[int]().IsSome())
			assert.False(t, maybe.None[string]().IsSome())
			assert.False(t, maybe.None[sample]().IsSome())
		})
	})

	t.Run("when creating from unknown input", func(t *testing.T) {
		t.Run("results in Some if non-nullable value is given", func(t *testing.T) {
			assert.Equals(t, maybe.Of[int](1), maybe.Some(1))
			assert.Equals(t, maybe.Of[string]("value"), maybe.Some("value"))
			assert.Equals(t, maybe.Of[sample](sample{1}), maybe.Some(sample{1}))
		})

		t.Run("results in None if null value is given", func(t *testing.T) {
			assert.Equals(t, maybe.Of[int](nil), maybe.None[int]())
		})
	})

	t.Run("when rendering as string", func(t *testing.T) {
		t.Run("Some informs the type and value", func(t *testing.T) {
			assert.Eq(t, maybe.Some(1).String(), "Some[int](1)")
			assert.Eq(t, maybe.Some("value").String(), "Some[string](value)")
			assert.Eq(t, maybe.Some(sample{1}).String(), "Some[maybe_test.sample]({value:1})")
		})

		t.Run("None is always empty", func(t *testing.T) {
			assert.Eq(t, maybe.Of[int](nil).String(), "None()")
		})
	})

	t.Run("when comparing", func(t *testing.T) {
		t.Run("Some against Some with == distinguishes by type and value", func(t *testing.T) {
			assert.True(t, maybe.Some(1) == maybe.Some(1))
			assert.False(t, maybe.Some(1) == maybe.Some(2))
			assert.True(t, maybe.Some("value") == maybe.Some("value"))
			assert.False(t, maybe.Some("value") == maybe.Some("other value"))
			assert.True(t, maybe.Some(sample{1}) == maybe.Some(sample{1}))
			assert.False(t, maybe.Some(sample{1}) == maybe.Some(sample{2}))
		})

		t.Run("Some against Some with .Equals distinguishes by type and value", func(t *testing.T) {
			assert.True(t, maybe.Some(1).Equals(maybe.Some(1)))
			assert.False(t, maybe.Some(1).Equals(maybe.Some(2)))
			assert.True(t, maybe.Some("value").Equals(maybe.Some("value")))
			assert.False(t, maybe.Some("value").Equals(maybe.Some("other value")))
			assert.True(t, maybe.Some(sample{1}).Equals(maybe.Some(sample{1})))
			assert.False(t, maybe.Some(sample{1}).Equals(maybe.Some(sample{2})))
		})

		t.Run("None against None with == distinguishes by type", func(t *testing.T) {
			assert.True(t, maybe.None[int]() == maybe.None[int]())
			assert.True(t, maybe.None[string]() == maybe.None[string]())
			assert.True(t, maybe.None[sample]() == maybe.None[sample]())
		})

		t.Run("None against None with .Equals distinguishes by type", func(t *testing.T) {
			assert.True(t, maybe.None[int]().Equals(maybe.None[int]()))
			assert.True(t, maybe.None[string]().Equals(maybe.None[string]()))
			assert.True(t, maybe.None[sample]().Equals(maybe.None[sample]()))
		})

		t.Run("None against Some is always false with ==", func(t *testing.T) {
			assert.False(t, maybe.None[int]() == maybe.Some(1))
			assert.False(t, maybe.Some(1) == maybe.None[int]())
		})

		t.Run("None against Some is always false with .Equals", func(t *testing.T) {
			assert.False(t, maybe.None[int]().Equals(maybe.Some(1)))
			assert.False(t, maybe.Some(1).Equals(maybe.None[int]()))
		})
	})

	t.Run("when pattern-matching", func(t *testing.T) {
		whenSome := func(it string) string { return fmt.Sprintf("some %s", it) }
		whenNone := func() string { return "none" }

		t.Run("Some evaluates the first function", func(t *testing.T) {
			assert.Eq(t, maybe.Match(maybe.Some("value"), whenSome, whenNone), "some value")
		})

		t.Run("None evaluates the second function", func(t *testing.T) {
			assert.Eq(t, maybe.Match(maybe.None[string](), whenSome, whenNone), "none")
		})
	})

	t.Run("when mapping", func(t *testing.T) {
		t.Run("Some remains Some but of the mapped type", func(t *testing.T) {
			assert.Equals(t, maybe.Map(maybe.Some(1), func(it int) int { return it * 2 }), maybe.Some(2))
			assert.Equals(t, maybe.Map(maybe.Some(1), func(it int) string { return fmt.Sprintf("value: %d", it) }), maybe.Some("value: 1"))
			assert.Equals(t, maybe.Map(maybe.Some(1), func(it int) sample { return sample{it} }), maybe.Some(sample{1}))
			assert.Equals(t, maybe.Map(maybe.Some("1"), func(it string) int { return mustParseInt(it) }), maybe.Some(1))
			assert.Equals(t, maybe.Map(maybe.Some(1), func(it int) maybe.Maybe[int] { return maybe.Some(it * 2) }), maybe.Some(maybe.Some(2)))
		})

		t.Run("Some becomes None of the mapped type if the result is null", func(t *testing.T) {
			assert.Equals(t, maybe.Map(maybe.Some(1), func(it int) any { return nil }), maybe.None[any]())
			// assert.Equals(t, maybe.Map(maybe.Some(1), func(it int) *sample { return nil }), maybe.None[*sample]()) // FIXME: pointers
		})

		t.Run("None always remains None but of the mapped type", func(t *testing.T) {
			assert.Equals(t, maybe.Map(maybe.None[int](), func(it int) int { return 1 }), maybe.None[int]())
			assert.Equals(t, maybe.Map(maybe.None[int](), func(it int) string { return "string" }), maybe.None[string]())
		})
	})

	t.Run("when flat-mapping", func(t *testing.T) {
		t.Run("behaves like mapping if result is not a Maybe", func(t *testing.T) {
			t.Run("Some remains Some but of the mapped type", func(t *testing.T) {
				assert.Equals(t, maybe.FlatMap[int, int](maybe.Some(1), func(it int) any { return it * 2 }), maybe.Some(2))
				assert.Equals(t, maybe.FlatMap[int, string](maybe.Some(1), func(it int) any { return fmt.Sprintf("value: %d", it) }), maybe.Some("value: 1"))
				assert.Equals(t, maybe.FlatMap[int, sample](maybe.Some(1), func(it int) any { return sample{it} }), maybe.Some(sample{1}))
				assert.Equals(t, maybe.FlatMap[string, int](maybe.Some("1"), func(it string) any { return mustParseInt(it) }), maybe.Some(1))
			})

			t.Run("Some becomes None of the mapped type if the result is null", func(t *testing.T) {
				assert.Equals(t, maybe.FlatMap[int, any](maybe.Some(1), func(it int) any { return nil }), maybe.None[any]())
			})

			t.Run("None always remains None but of the mapped type", func(t *testing.T) {
				assert.Equals(t, maybe.FlatMap[int, int](maybe.None[int](), func(it int) any { return it * 2 }), maybe.None[int]())
				assert.Equals(t, maybe.FlatMap[int, *sample](maybe.None[int](), func(it int) any { return nil }), maybe.None[*sample]())
			})
		})

		t.Run("flattens if result is a Maybe", func(t *testing.T) {
			t.Run("Some remains Some but of the mapped type", func(t *testing.T) {
				assert.Equals(t, maybe.FlatMap[int, int](maybe.Some(1), func(it int) any { return maybe.Some(it * 2) }), maybe.Some(2))
				assert.Equals(t, maybe.FlatMap[int, string](maybe.Some(1), func(it int) any { return maybe.Some(fmt.Sprintf("value: %d", it)) }), maybe.Some("value: 1"))
				assert.Equals(t, maybe.FlatMap[int, sample](maybe.Some(1), func(it int) any { return maybe.Some(sample{it}) }), maybe.Some(sample{1}))
				assert.Equals(t, maybe.FlatMap[string, int](maybe.Some("1"), func(it string) any { return maybe.Some(mustParseInt(it)) }), maybe.Some(1))
			})

			t.Run("Some becomes None of the mapped type if the result is null", func(t *testing.T) {
				assert.Equals(t, maybe.FlatMap[int, any](maybe.Some(1), func(it int) any { return maybe.Of[any](nil) }), maybe.None[any]())
			})

			t.Run("None always remains None but of the mapped type", func(t *testing.T) {
				assert.Equals(t, maybe.FlatMap[int, string](maybe.None[int](), func(it int) any { return maybe.Some(fmt.Sprintf("value: %d", it)) }), maybe.None[string]())
				assert.Equals(t, maybe.FlatMap[int, *sample](maybe.None[int](), func(it int) any { return maybe.Of[*sample](nil) }), maybe.None[*sample]())
			})
		})
	})

	t.Run("when unwrapping", func(t *testing.T) {
		t.Run("Some returns the underlying value", func(t *testing.T) {
			assert.Eq(t, maybe.Some(1).Unwrap(), 1)
			assert.Eq(t, maybe.Some("value").Unwrap(), "value")
			assert.Eq(t, maybe.Some(sample{1}).Unwrap(), sample{1})
		})

		t.Run("None panics", func(t *testing.T) {
			assert.PanicsWith(t, func() { maybe.None[int]().Unwrap() }, "nothing to unwrap from None()")
			assert.PanicsWith(t, func() { maybe.None[string]().Unwrap() }, "nothing to unwrap from None()")
			assert.PanicsWith(t, func() { maybe.None[sample]().Unwrap() }, "nothing to unwrap from None()")
		})
	})

	t.Run("when unwrapping with default value", func(t *testing.T) {
		t.Run("Some returns the underlying value", func(t *testing.T) {
			assert.Eq(t, maybe.Some(1).UnwrapOr(-1), 1)
			assert.Eq(t, maybe.Some("value").UnwrapOr("empty"), "value")
		})

		t.Run("None returns the given default value", func(t *testing.T) {
			assert.Eq(t, maybe.None[int]().UnwrapOr(-1), -1)
			assert.Eq(t, maybe.None[string]().UnwrapOr("empty"), "empty")
			assert.Eq(t, maybe.None[sample]().UnwrapOr(sample{-1}), sample{-1})
		})
	})

	t.Run("when unwrapping with default function", func(t *testing.T) {
		t.Run("Some returns the underlying value", func(t *testing.T) {
			assert.Eq(t, maybe.Some(1).UnwrapOrElse(func() int { return -1 }), 1)
			assert.Eq(t, maybe.Some("value").UnwrapOrElse(func() string { return "empty" }), "value")
			assert.Eq(t, maybe.Some(sample{1}).UnwrapOrElse(func() sample { return sample{-1} }), sample{1})
		})

		t.Run("None returns the result of evaluating the given default function", func(t *testing.T) {
			assert.Eq(t, maybe.None[int]().UnwrapOrElse(func() int { return -1 }), -1)
			assert.Eq(t, maybe.None[string]().UnwrapOrElse(func() string { return "empty" }), "empty")
			assert.Eq(t, maybe.None[sample]().UnwrapOrElse(func() sample { return sample{-1} }), sample{-1})
		})
	})

	t.Run("holds monadic properties", func(t *testing.T) {
		t.Run("left identity", func(t *testing.T) {
			v := 1
			m := maybe.Some(v)
			f := func(it int) any { return maybe.Some(it * 2) }

			a := maybe.FlatMap[int, int](m, f)
			b := f(v).(maybe.Maybe[int])
			assert.True(t, a.Equals(b))
		})

		t.Run("right identity", func(t *testing.T) {
			m := maybe.Some(1)
			a := maybe.FlatMap[int, int](m, func(it int) any { return maybe.Some(it) })

			assert.True(t, a.Equals(m))
		})

		t.Run("associativity", func(t *testing.T) {
			m := maybe.Some(1)
			f := func(it int) any { return maybe.Some(it * 2) }
			g := func(it int) any { return maybe.Some(it + 4) }

			a := maybe.FlatMap[int, int](maybe.FlatMap[int, int](m, f), g)
			b := maybe.FlatMap[int, int](m, func(it int) any { return maybe.FlatMap[int, int](f(it).(maybe.Maybe[int]), g) })
			assert.True(t, a.Equals(b))
		})
	})
}
