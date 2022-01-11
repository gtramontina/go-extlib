package either_test

import (
	"fmt"
	"testing"

	"github.com/gtramontina/go-extlib/either"
	"github.com/gtramontina/go-extlib/maybe"
	"github.com/gtramontina/go-extlib/testing/assert"
)

func TestEither(t *testing.T) {
	type sample struct{ value int }

	t.Run("when type checking", func(t *testing.T) {
		t.Run("Left is always Left", func(t *testing.T) {
			assert.True(t, either.Left[int, int](1).IsLeft())
			assert.True(t, either.Left[string, string]("value").IsLeft())
			assert.True(t, either.Left[sample, sample](sample{1}).IsLeft())
		})

		t.Run("Left is never Right", func(t *testing.T) {
			assert.False(t, either.Left[int, int](1).IsRight())
			assert.False(t, either.Left[string, string]("value").IsRight())
			assert.False(t, either.Left[sample, sample](sample{1}).IsRight())
		})

		t.Run("Right is always Right", func(t *testing.T) {
			assert.True(t, either.Right[int, int](1).IsRight())
			assert.True(t, either.Right[string, string]("value").IsRight())
			assert.True(t, either.Right[sample, sample](sample{1}).IsRight())
		})

		t.Run("Right is never Left", func(t *testing.T) {
			assert.False(t, either.Right[int, int](1).IsLeft())
			assert.False(t, either.Right[string, string]("value").IsLeft())
			assert.False(t, either.Right[sample, sample](sample{1}).IsLeft())
		})
	})

	t.Run("when rendering as string", func(t *testing.T) {
		t.Run("Left informs the type and value", func(t *testing.T) {
			assert.Eq(t, either.Left[int, int](1).String(), "Left[int](1)")
			assert.Eq(t, either.Left[string, string]("value").String(), "Left[string](value)")
			assert.Eq(t, either.Left[sample, sample](sample{1}).String(), "Left[either_test.sample]({value:1})")
		})

		t.Run("Right informs the type and value", func(t *testing.T) {
			assert.Eq(t, either.Right[int, int](1).String(), "Right[int](1)")
			assert.Eq(t, either.Right[string, string]("value").String(), "Right[string](value)")
			assert.Eq(t, either.Right[sample, sample](sample{1}).String(), "Right[either_test.sample]({value:1})")
		})
	})

	t.Run("when comparing", func(t *testing.T) {
		t.Run("Left against Left with == distinguishes by type and value", func(t *testing.T) {
			assert.True(t, either.Left[int, int](1) == either.Left[int, int](1))
			assert.False(t, either.Left[int, int](1) == either.Left[int, int](2))
			assert.True(t, either.Left[string, string]("value") == either.Left[string, string]("value"))
			assert.False(t, either.Left[string, string]("value") == either.Left[string, string]("other value"))
			assert.True(t, either.Left[sample, sample](sample{1}) == either.Left[sample, sample](sample{1}))
			assert.False(t, either.Left[sample, sample](sample{1}) == either.Left[sample, sample](sample{2}))
		})

		t.Run("Left against Left with .Equals distinguishes by type and value", func(t *testing.T) {
			assert.True(t, either.Left[int, int](1).Equals(either.Left[int, int](1)))
			assert.False(t, either.Left[int, int](1).Equals(either.Left[int, int](2)))
			assert.True(t, either.Left[string, string]("value").Equals(either.Left[string, string]("value")))
			assert.False(t, either.Left[string, string]("value").Equals(either.Left[string, string]("other value")))
			assert.True(t, either.Left[sample, sample](sample{1}).Equals(either.Left[sample, sample](sample{1})))
			assert.False(t, either.Left[sample, sample](sample{1}).Equals(either.Left[sample, sample](sample{2})))
		})

		t.Run("Right against Right with == distinguishes by type", func(t *testing.T) {
			assert.True(t, either.Right[int, int](1) == either.Right[int, int](1))
			assert.False(t, either.Right[int, int](1) == either.Right[int, int](2))
			assert.True(t, either.Right[string, string]("value") == either.Right[string, string]("value"))
			assert.False(t, either.Right[string, string]("value") == either.Right[string, string]("other value"))
			assert.True(t, either.Right[sample, sample](sample{1}) == either.Right[sample, sample](sample{1}))
			assert.False(t, either.Right[sample, sample](sample{1}) == either.Right[sample, sample](sample{2}))
		})

		t.Run("Right against Right with .Equals distinguishes by type", func(t *testing.T) {
			assert.True(t, either.Right[int, int](1).Equals(either.Right[int, int](1)))
			assert.False(t, either.Right[int, int](1).Equals(either.Right[int, int](2)))
			assert.True(t, either.Right[string, string]("value").Equals(either.Right[string, string]("value")))
			assert.False(t, either.Right[string, string]("value").Equals(either.Right[string, string]("other value")))
			assert.True(t, either.Right[sample, sample](sample{1}).Equals(either.Right[sample, sample](sample{1})))
			assert.False(t, either.Right[sample, sample](sample{1}).Equals(either.Right[sample, sample](sample{2})))
		})

		t.Run("Right against Left is always false with ==", func(t *testing.T) {
			assert.False(t, either.Right[int, int](1) == either.Left[int, int](1))
			assert.False(t, either.Right[int, string]("1") == either.Left[int, string](1))
			assert.False(t, either.Left[int, int](1) == either.Right[int, int](1))
			assert.False(t, either.Left[string, int]("1") == either.Right[string, int](1))
		})

		t.Run("Right against Left is always false with .Equals", func(t *testing.T) {
			assert.False(t, either.Right[int, int](1).Equals(either.Left[int, int](1)))
			assert.False(t, either.Left[int, int](1).Equals(either.Right[int, int](1)))
		})
	})

	t.Run("when getting left", func(t *testing.T) {
		t.Run("Left becomes Some", func(t *testing.T) {
			assert.Equals(t, either.Left[int, int](1).Left(), maybe.Some(1))
			assert.Equals(t, either.Left[int, string](1).Left(), maybe.Some(1))
			assert.Equals(t, either.Left[string, string]("value").Left(), maybe.Some("value"))
			assert.Equals(t, either.Left[string, int]("value").Left(), maybe.Some("value"))
			assert.Equals(t, either.Left[sample, sample](sample{1}).Left(), maybe.Some(sample{1}))
		})

		t.Run("Right becomes None", func(t *testing.T) {
			assert.Equals(t, either.Right[int, int](1).Left(), maybe.None[int]())
			assert.Equals(t, either.Right[int, string]("value").Left(), maybe.None[int]())
			assert.Equals(t, either.Right[string, string]("value").Left(), maybe.None[string]())
			assert.Equals(t, either.Right[string, int](1).Left(), maybe.None[string]())
			assert.Equals(t, either.Right[sample, sample](sample{1}).Left(), maybe.None[sample]())
		})
	})

	t.Run("when getting left with default value", func(t *testing.T) {
		t.Run("Left returns the underlying value", func(t *testing.T) {
			assert.Eq(t, either.Left[int, int](1).LeftOr(-1), 1)
			assert.Eq(t, either.Left[string, int]("value").LeftOr("empty"), "value")
		})

		t.Run("Right returns the given default value", func(t *testing.T) {
			assert.Eq(t, either.Right[int, int](1).LeftOr(-1), -1)
			assert.Eq(t, either.Right[string, int](1).LeftOr("empty"), "empty")
		})
	})

	t.Run("when getting left with default function", func(t *testing.T) {
		t.Run("Left returns the underlying value", func(t *testing.T) {
			assert.Eq(t, either.Left[int, int](1).LeftOrElse(func() int { return -1 }), 1)
			assert.Eq(t, either.Left[string, int]("value").LeftOrElse(func() string { return "empty" }), "value")
		})

		t.Run("Right returns the result of evaluating the given default function", func(t *testing.T) {
			assert.Eq(t, either.Right[int, int](1).LeftOrElse(func() int { return -1 }), -1)
			assert.Eq(t, either.Right[string, int](1).LeftOrElse(func() string { return "empty" }), "empty")
		})
	})

	t.Run("when getting right", func(t *testing.T) {
		t.Run("Left becomes None", func(t *testing.T) {
			assert.Equals(t, either.Left[int, int](1).Right(), maybe.None[int]())
			assert.Equals(t, either.Left[int, string](1).Right(), maybe.None[string]())
			assert.Equals(t, either.Left[string, string]("value").Right(), maybe.None[string]())
			assert.Equals(t, either.Left[string, int]("value").Right(), maybe.None[int]())
			assert.Equals(t, either.Left[sample, sample](sample{1}).Right(), maybe.None[sample]())
		})

		t.Run("Right becomes Some", func(t *testing.T) {
			assert.Equals(t, either.Right[int, int](1).Right(), maybe.Some(1))
			assert.Equals(t, either.Right[int, string]("value").Right(), maybe.Some("value"))
			assert.Equals(t, either.Right[string, string]("value").Right(), maybe.Some("value"))
			assert.Equals(t, either.Right[string, int](1).Right(), maybe.Some(1))
			assert.Equals(t, either.Right[sample, sample](sample{1}).Right(), maybe.Some(sample{1}))
		})
	})

	t.Run("when getting right with default value", func(t *testing.T) {
		t.Run("Left returns the given default value", func(t *testing.T) {
			assert.Eq(t, either.Left[int, int](1).RightOr(-1), -1)
			assert.Eq(t, either.Left[int, string](1).RightOr("empty"), "empty")
		})

		t.Run("Right returns the underlying value", func(t *testing.T) {
			assert.Eq(t, either.Right[int, int](1).RightOr(-1), 1)
			assert.Eq(t, either.Right[int, string]("value").RightOr("empty"), "value")
		})
	})

	t.Run("when getting right with default function", func(t *testing.T) {
		t.Run("Left returns the result of evaluating the given default function", func(t *testing.T) {
			assert.Eq(t, either.Left[int, int](1).RightOrElse(func() int { return -1 }), -1)
			assert.Eq(t, either.Left[int, string](1).RightOrElse(func() string { return "empty" }), "empty")
		})

		t.Run("Right returns the underlying value", func(t *testing.T) {
			assert.Eq(t, either.Right[int, int](1).RightOrElse(func() int { return -1 }), 1)
			assert.Eq(t, either.Right[int, string]("value").RightOrElse(func() string { return "empty" }), "value")
		})
	})

	t.Run("when unwrapping left", func(t *testing.T) {
		t.Run("Left returns the underlying value", func(t *testing.T) {
			assert.Eq(t, either.Left[int, int](1).UnwrapLeft(), 1)
			assert.Eq(t, either.Left[string, int]("value").UnwrapLeft(), "value")
			assert.Eq(t, either.Left[sample, sample](sample{1}).UnwrapLeft(), sample{1})
		})

		t.Run("Right panics", func(t *testing.T) {
			assert.Panic(t, func() { either.Right[int, int](1).UnwrapLeft() }, "nothing to unwrap left from Right")
			assert.Panic(t, func() { either.Right[string, string]("value").UnwrapLeft() }, "nothing to unwrap left from Right")
			assert.Panic(t, func() { either.Right[sample, sample](sample{1}).UnwrapLeft() }, "nothing to unwrap left from Right")
		})
	})

	t.Run("when unwrapping right", func(t *testing.T) {
		t.Run("Left panics", func(t *testing.T) {
			assert.Panic(t, func() { either.Left[int, int](1).UnwrapRight() }, "nothing to unwrap right from Left")
			assert.Panic(t, func() { either.Left[string, string]("value").UnwrapRight() }, "nothing to unwrap right from Left")
			assert.Panic(t, func() { either.Left[sample, sample](sample{1}).UnwrapRight() }, "nothing to unwrap right from Left")
		})

		t.Run("Right returns the underlying value", func(t *testing.T) {
			assert.Eq(t, either.Right[int, int](1).UnwrapRight(), 1)
			assert.Eq(t, either.Right[int, string]("value").UnwrapRight(), "value")
			assert.Eq(t, either.Right[sample, sample](sample{1}).UnwrapRight(), sample{1})
		})
	})

	t.Run("when flipping", func(t *testing.T) {
		t.Run("Left becomes Right", func(t *testing.T) {
			assert.Equals(t, either.Left[int, int](1).Flip(), either.Right[int, int](1))
			assert.Equals(t, either.Left[int, string](1).Flip(), either.Right[string, int](1))
			assert.Equals(t, either.Left[string, string]("value").Flip(), either.Right[string, string]("value"))
			assert.Equals(t, either.Left[string, int]("value").Flip(), either.Right[int, string]("value"))
			assert.Equals(t, either.Left[sample, sample](sample{1}).Flip(), either.Right[sample, sample](sample{1}))
		})

		t.Run("Right becomes Left", func(t *testing.T) {
			assert.Equals(t, either.Right[int, int](1).Flip(), either.Left[int, int](1))
			assert.Equals(t, either.Right[int, string]("value").Flip(), either.Left[string, int]("value"))
			assert.Equals(t, either.Right[string, string]("value").Flip(), either.Left[string, string]("value"))
			assert.Equals(t, either.Right[string, int](1).Flip(), either.Left[int, string](1))
			assert.Equals(t, either.Right[sample, sample](sample{1}).Flip(), either.Left[sample, sample](sample{1}))
		})
	})

	t.Run("when pattern-matching", func(t *testing.T) {
		whenLeft := func(it int) string { return fmt.Sprintf("left: %d", it) }
		whenRight := func(it string) string { return "right: " + it }

		t.Run("Left evaluates the first function", func(t *testing.T) {
			assert.Eq(t, either.Match(either.Left[int, string](1), whenLeft, whenRight), "left: 1")
		})

		t.Run("Right evaluates the second function", func(t *testing.T) {
			assert.Eq(t, either.Match(either.Right[int, string]("1"), whenLeft, whenRight), "right: 1")
		})
	})

	t.Run("when mapping", func(t *testing.T) {
		t.Run("Left wraps in Left the result of applying the mapper function on the value", func(t *testing.T) {
			assert.Equals(t, either.Map(either.Left[int, int](1), func(it int) int { return it + 1 }), either.Left[int, int](2))
			assert.Equals(t, either.Map(either.Left[int, int](1), func(it int) string { return fmt.Sprintf("%d", it) }), either.Left[string, string]("1"))
		})

		t.Run("Right wraps in Right the result of applying the mapper function on the value", func(t *testing.T) {
			assert.Equals(t, either.Map(either.Right[int, int](1), func(it int) int { return it + 1 }), either.Right[int, int](2))
			assert.Equals(t, either.Map(either.Right[int, int](1), func(it int) string { return fmt.Sprintf("%d", it) }), either.Right[string, string]("1"))
		})
	})

	t.Run("when mapping left", func(t *testing.T) {
		t.Run("Left wraps in Left the result of applying the mapper function on the value", func(t *testing.T) {
			assert.Equals(t, either.MapLeft(either.Left[int, int](1), func(it int) int { return it + 1 }), either.Left[int, int](2))
			assert.Equals(t, either.MapLeft(either.Left[int, int](1), func(it int) string { return fmt.Sprintf("%d", it) }), either.Left[string, int]("1"))
		})

		t.Run("Right remains Right obeying the mapper function output type", func(t *testing.T) {
			assert.Equals(t, either.MapLeft(either.Right[int, int](1), func(it int) int { return it + 1 }), either.Right[int, int](1))
			assert.Equals(t, either.MapLeft(either.Right[int, int](1), func(it int) string { return fmt.Sprintf("%d", it) }), either.Right[string, int](1))
		})
	})

	t.Run("when mapping right", func(t *testing.T) {
		t.Run("Right wraps in Right the result of applying the mapper function on the value", func(t *testing.T) {
			assert.Equals(t, either.MapRight(either.Right[int, int](1), func(it int) int { return it + 1 }), either.Right[int, int](2))
			assert.Equals(t, either.MapRight(either.Right[int, int](1), func(it int) string { return fmt.Sprintf("%d", it) }), either.Right[int, string]("1"))
		})

		t.Run("Left remains Left obeying the mapper function output type", func(t *testing.T) {
			assert.Equals(t, either.MapRight(either.Left[int, int](1), func(it int) int { return it + 1 }), either.Left[int, int](1))
			assert.Equals(t, either.MapRight(either.Left[int, int](1), func(it int) string { return fmt.Sprintf("%d", it) }), either.Left[int, string](1))
		})
	})
}
