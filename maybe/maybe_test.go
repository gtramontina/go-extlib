package maybe_test

import (
	"fmt"
	"github.com/gtramontina/go-collections/internal/assert"
	"github.com/gtramontina/go-collections/maybe"
	"strconv"
	"testing"
)

func TestMaybe(t *testing.T) {
	type sample struct{ value int }

	t.Run(`creates "some" from a non-nullable value`, func(t *testing.T) {
		assert.Equals(t, maybe.Of[string]("value"), maybe.Some("value"))
		assert.Equals(t, maybe.Of[sample](sample{0}), maybe.Some(sample{0}))
		assert.Equals(t, maybe.Of[sample](sample{1}), maybe.Some(sample{1}))
	})

	t.Run(`creates "none" from "nil"`, func(t *testing.T) {
		assert.Equals(t, maybe.Of[int](nil), maybe.None[int]())
	})

	t.Run("as string", func(t *testing.T) {
		assert.Eq(t, maybe.Of[int](nil).String(), "None()")
		assert.Eq(t, maybe.Of[int](0).String(), "Some[int](0)")
		assert.Eq(t, maybe.Of[string]("0").String(), "Some[string](0)")
		assert.Eq(t, maybe.Of[sample](sample{0}).String(), "Some[maybe_test.sample]({value:0})")
	})

	t.Run("is comparable", func(t *testing.T) {
		assert.True(t, maybe.Some(10).Equals(maybe.Some(10)))
		assert.False(t, maybe.Some(10).Equals(maybe.Some(20)))
		assert.True(t, maybe.Some("10").Equals(maybe.Some("10")))
		assert.False(t, maybe.Some("10").Equals(maybe.Some("20")))
		assert.True(t, maybe.None[int]().Equals(maybe.None[int]()))
		assert.False(t, maybe.Some(10).Equals(maybe.None[int]()))
		assert.False(t, maybe.None[string]().Equals(maybe.Some("10")))
	})

	t.Run("pattern matches", func(t *testing.T) {
		whenSome := func(it string) string { return fmt.Sprintf("some %s", it) }
		whenNone := func() string { return "none" }
		assert.Eq(t, maybe.Match(maybe.Some("value"), whenSome, whenNone), "some value")
		assert.Eq(t, maybe.Match(maybe.None[string](), whenSome, whenNone), "none")
	})

	t.Run("maps to another type", func(t *testing.T) {
		assert.Equals(t, maybe.Map(maybe.Some(10), func(it int) string { return fmt.Sprintf("value: %d", it) }), maybe.Some("value: 10"))
		assert.Equals(t, maybe.Map(maybe.Some(10), func(it int) sample { return sample{it} }), maybe.Some(sample{10}))
		assert.Equals(t, maybe.Map(maybe.Some("10"), func(it string) int { out, _ := strconv.ParseInt(it, 10, 0); return int(out) }), maybe.Some(10))
		assert.Equals(t, maybe.Map(maybe.Some(10), func(it int) interface{} { return nil }), maybe.None[interface{}]())
	})

	t.Run(`"none" always maps to "none"`, func(t *testing.T) {
		assert.Equals(t, maybe.Map(maybe.None[int](), func(it int) string { return fmt.Sprintf("value: %d", it) }), maybe.None[string]())
	})

	t.Run("flat maps to another type", func(t *testing.T) {
		assert.Equals(t, maybe.FlatMap(maybe.Some(10), func(it int) maybe.Maybe[string] { return maybe.Some(fmt.Sprintf("value: %d", it)) }), maybe.Some("value: 10"))
		assert.Equals(t, maybe.FlatMap(maybe.Some(10), func(it int) maybe.Maybe[sample] { return maybe.Some(sample{it}) }), maybe.Some(sample{10}))
		assert.Equals(t, maybe.FlatMap(maybe.Some("10"), func(it string) maybe.Maybe[int] { out, _ := strconv.ParseInt(it, 10, 0); return maybe.Some(int(out)) }), maybe.Some(10))
		assert.Equals(t, maybe.FlatMap(maybe.Some(10), func(it int) maybe.Maybe[interface{}] { return maybe.Of[interface{}](nil) }), maybe.None[interface{}]())
	})

	t.Run(`"none" always flatMaps to "none"`, func(t *testing.T) {
		assert.Equals(t, maybe.FlatMap(maybe.None[int](), func(it int) maybe.Maybe[string] { return maybe.Some(fmt.Sprintf("value: %d", it)) }), maybe.None[string]())
	})

	t.Run(`can unwrap a "some"`, func(t *testing.T) {
		assert.Eq(t, maybe.Some(10).Unwrap(), 10)
		assert.Eq(t, maybe.Some("10").Unwrap(), "10")
	})

	t.Run(`panics when unwrapping "none"`, func(t *testing.T) {
		assert.Panic(t, func() { maybe.None[int]().Unwrap() }, "nothing to unwrap from None()")
		assert.Panic(t, func() { maybe.Map(maybe.Some(10), func(it int) interface{} { return nil }).Unwrap() }, "nothing to unwrap from None()")
	})
}
