package result_test

import (
	"errors"
	"fmt"
	"github.com/gtramontina/go-collections/internal/assert"
	"github.com/gtramontina/go-collections/maybe"
	"github.com/gtramontina/go-collections/result"
	"testing"
)

func TestResult(t *testing.T) {
	type sample struct{ value int }

	t.Run("when type checking", func(t *testing.T) {
		t.Run("Ok is always Ok", func(t *testing.T) {
			assert.True(t, result.Ok[int](1).IsOk())
			assert.True(t, result.Ok[string]("value").IsOk())
			assert.True(t, result.Ok[sample](sample{1}).IsOk())
		})

		t.Run("Ok is never Err", func(t *testing.T) {
			assert.False(t, result.Ok[int](1).IsErr())
			assert.False(t, result.Ok[string]("value").IsErr())
			assert.False(t, result.Ok[sample](sample{1}).IsErr())
		})

		t.Run("Err is always Err", func(t *testing.T) {
			assert.True(t, result.Err[int](errors.New("error message")).IsErr())
			assert.True(t, result.Err[string](errors.New("error message")).IsErr())
			assert.True(t, result.Err[sample](errors.New("error message")).IsErr())
		})

		t.Run("Err is never Ok", func(t *testing.T) {
			assert.False(t, result.Err[int](errors.New("error message")).IsOk())
			assert.False(t, result.Err[string](errors.New("error message")).IsOk())
			assert.False(t, result.Err[sample](errors.New("error message")).IsOk())
		})
	})

	t.Run("when rendering as string", func(t *testing.T) {
		t.Run("Ok informs the type and value", func(t *testing.T) {
			assert.Eq(t, result.Ok[int](1).String(), "Ok[int](1)")
			assert.Eq(t, result.Ok[string]("value").String(), "Ok[string](value)")
			assert.Eq(t, result.Ok[sample](sample{1}).String(), "Ok[result_test.sample]({value:1})")
		})

		t.Run("Err informs the type and value", func(t *testing.T) {
			assert.Eq(t, result.Err[int](errors.New("error message")).String(), "Err(error message)")
			assert.Eq(t, result.Err[string](errors.New("error message")).String(), "Err(error message)")
			assert.Eq(t, result.Err[sample](errors.New("error message")).String(), "Err(error message)")
		})
	})

	t.Run("when comparing", func(t *testing.T) {
		t.Run("Ok against Ok with == distinguishes by type and value", func(t *testing.T) {
			assert.True(t, result.Ok[int](1) == result.Ok[int](1))
			assert.False(t, result.Ok[int](1) == result.Ok[int](2))
			assert.True(t, result.Ok[string]("value") == result.Ok[string]("value"))
			assert.False(t, result.Ok[string]("value") == result.Ok[string]("other value"))
			assert.True(t, result.Ok[sample](sample{1}) == result.Ok[sample](sample{1}))
			assert.False(t, result.Ok[sample](sample{1}) == result.Ok[sample](sample{2}))
		})

		t.Run("Ok against Ok with .Equals distinguishes by type and value", func(t *testing.T) {
			assert.True(t, result.Ok[int](1).Equals(result.Ok[int](1)))
			assert.False(t, result.Ok[int](1).Equals(result.Ok[int](2)))
			assert.True(t, result.Ok[string]("value").Equals(result.Ok[string]("value")))
			assert.False(t, result.Ok[string]("value").Equals(result.Ok[string]("other value")))
			assert.True(t, result.Ok[sample](sample{1}).Equals(result.Ok[sample](sample{1})))
			assert.False(t, result.Ok[sample](sample{1}).Equals(result.Ok[sample](sample{2})))
		})

		t.Run("Err against Err with == requires the error to be the same", func(t *testing.T) {
			err := errors.New("error message")
			differentErr := errors.New("different error message")
			assert.True(t, result.Err[int](err) == result.Err[int](err))
			assert.False(t, result.Err[int](err) == result.Err[int](differentErr))
			assert.True(t, result.Err[string](err) == result.Err[string](err))
			assert.False(t, result.Err[string](err) == result.Err[string](differentErr))
			assert.True(t, result.Err[sample](err) == result.Err[sample](err))
			assert.False(t, result.Err[sample](err) == result.Err[sample](differentErr))
			assert.True(t, result.Err[int](differentErr) == result.Err[int](differentErr))
		})

		t.Run("Err against Err with .Equals distinguishes by type only", func(t *testing.T) {
			assert.True(t, result.Err[int](errors.New("error message")).Equals(result.Err[int](errors.New("error message"))))
			assert.False(t, result.Err[int](errors.New("error message")).Equals(result.Err[int](errors.New("different error message"))))
			assert.True(t, result.Err[string](errors.New("error message")).Equals(result.Err[string](errors.New("error message"))))
			assert.False(t, result.Err[string](errors.New("error message")).Equals(result.Err[string](errors.New("different error message"))))
			assert.True(t, result.Err[sample](errors.New("error message")).Equals(result.Err[sample](errors.New("error message"))))
			assert.False(t, result.Err[sample](errors.New("error message")).Equals(result.Err[sample](errors.New("different error message"))))
			assert.True(t, result.Err[int](errors.New("different error message")).Equals(result.Err[int](errors.New("different error message"))))
		})

		t.Run("Err against Ok is always false with ==", func(t *testing.T) {
			assert.False(t, result.Err[int](errors.New("error message")) == result.Ok[int](1))
			assert.False(t, result.Err[string](errors.New("error message")) == result.Ok[string]("value"))
			assert.False(t, result.Ok[string]("value") == result.Err[string](errors.New("error message")))
			assert.False(t, result.Ok[sample](sample{1}) == result.Err[sample](errors.New("error message")))
		})

		t.Run("Err against Ok is always false with .Equals", func(t *testing.T) {
			assert.False(t, result.Err[int](errors.New("error message")).Equals(result.Ok[int](1)))
			assert.False(t, result.Ok[int](1).Equals(result.Err[int](errors.New("error message"))))
		})
	})

	t.Run("when getting ok", func(t *testing.T) {
		t.Run("Ok becomes Some", func(t *testing.T) {
			assert.Equals(t, result.Ok[int](1).Ok(), maybe.Some(1))
			assert.Equals(t, result.Ok[string]("value").Ok(), maybe.Some("value"))
			assert.Equals(t, result.Ok[sample](sample{1}).Ok(), maybe.Some(sample{1}))
		})

		t.Run("Err becomes None", func(t *testing.T) {
			assert.Equals(t, result.Err[int](errors.New("error message")).Ok(), maybe.None[int]())
			assert.Equals(t, result.Err[string](errors.New("error message")).Ok(), maybe.None[string]())
			assert.Equals(t, result.Err[sample](errors.New("error message")).Ok(), maybe.None[sample]())
		})
	})

	t.Run("when unwrapping ok", func(t *testing.T) {
		t.Run("Ok returns the underlying value", func(t *testing.T) {
			assert.Eq(t, result.Ok[int](1).Unwrap(), 1)
			assert.Eq(t, result.Ok[string]("value").Unwrap(), "value")
			assert.Eq(t, result.Ok[sample](sample{1}).Unwrap(), sample{1})
		})

		t.Run("Err panics with the error", func(t *testing.T) {
			assert.Panic(t, func() { result.Err[int](errors.New("error message int")).Unwrap() }, errors.New("error message int"))
			assert.Panic(t, func() { result.Err[string](errors.New("error message string")).Unwrap() }, errors.New("error message string"))
			assert.Panic(t, func() { result.Err[sample](errors.New("error message sample")).Unwrap() }, errors.New("error message sample"))
		})
	})

	t.Run("when unwrapping err", func(t *testing.T) {
		t.Run("Ok panics with the underlying value", func(t *testing.T) {
			assert.Panic(t, func() { result.Ok[int](1).UnwrapErr() }, 1)
			assert.Panic(t, func() { result.Ok[string]("value").UnwrapErr() }, "value")
			assert.Panic(t, func() { result.Ok[sample](sample{1}).UnwrapErr() }, sample{1})
		})

		t.Run("Err returns the underlying error", func(t *testing.T) {
			assert.DeepEqual(t, result.Err[int](errors.New("error message int")).UnwrapErr(), errors.New("error message int"))
			assert.DeepEqual(t, result.Err[string](errors.New("error message string")).UnwrapErr(), errors.New("error message string"))
			assert.DeepEqual(t, result.Err[sample](errors.New("error message sample")).UnwrapErr(), errors.New("error message sample"))
		})
	})

	t.Run("when unwrapping with default value", func(t *testing.T) {
		t.Run("Ok returns the underlying value", func(t *testing.T) {
			assert.Eq(t, result.Ok[int](1).UnwrapOr(-1), 1)
			assert.Eq(t, result.Ok[string]("value").UnwrapOr("empty"), "value")
		})

		t.Run("Err returns the given default value", func(t *testing.T) {
			assert.Eq(t, result.Err[int](errors.New("error message")).UnwrapOr(-1), -1)
			assert.Eq(t, result.Err[string](errors.New("error message")).UnwrapOr("empty"), "empty")
		})
	})

	t.Run("when unwrapping with default function", func(t *testing.T) {
		t.Run("Ok returns the underlying value", func(t *testing.T) {
			assert.Eq(t, result.Ok[int](1).UnwrapOrElse(func() int { return -1 }), 1)
			assert.Eq(t, result.Ok[string]("value").UnwrapOrElse(func() string { return "empty" }), "value")
		})

		t.Run("Err returns the result of evaluating the given default function", func(t *testing.T) {
			assert.Eq(t, result.Err[int](errors.New("error message")).UnwrapOrElse(func() int { return -1 }), -1)
			assert.Eq(t, result.Err[string](errors.New("error message")).UnwrapOrElse(func() string { return "empty" }), "empty")
		})
	})

	t.Run("when pattern-matching", func(t *testing.T) {
		whenOk := func(it int) string { return fmt.Sprintf("ok: %d", it) }
		whenErr := func(it error) string { return "err: " + it.Error() }

		t.Run("Ok evaluates the first function", func(t *testing.T) {
			assert.Eq(t, result.Match(result.Ok[int](1), whenOk, whenErr), "ok: 1")
		})

		t.Run("Err evaluates the second function", func(t *testing.T) {
			assert.Eq(t, result.Match(result.Err[int](errors.New("error message")), whenOk, whenErr), "err: error message")
		})
	})

	t.Run("when mapping", func(t *testing.T) {
		t.Run("Ok wraps in Ok the result of applying the mapper function on the value", func(t *testing.T) {
			assert.Equals(t, result.Map(result.Ok[int](1), func(it int) int { return it + 1 }), result.Ok[int](2))
			assert.Equals(t, result.Map(result.Ok[int](1), func(it int) string { return fmt.Sprintf("%d", it) }), result.Ok[string]("1"))
		})

		t.Run("Err wraps in Err the result of applying the mapper function on the value", func(t *testing.T) {
			assert.Equals(t, result.Map(result.Err[int](errors.New("error message")), func(it int) int { return it + 1 }), result.Err[int](errors.New("error message")))
			assert.Equals(t, result.Map(result.Err[int](errors.New("error message")), func(it int) string { return fmt.Sprintf("%d", it) }), result.Err[string](errors.New("error message")))
		})
	})

	t.Run("when mapping err", func(t *testing.T) {
		t.Run("Err wraps in Err the result of applying the mapper function on the value", func(t *testing.T) {
			assert.Equals(t, result.MapErr(result.Err[int](errors.New("error message")), func(it error) error { return fmt.Errorf("error#1: %v", it) }), result.Err[int](errors.New("error#1: error message")))
			assert.Equals(t, result.MapErr(result.Err[string](errors.New("error message")), func(it error) error { return fmt.Errorf("error#2: %v", it) }), result.Err[string](errors.New("error#2: error message")))
		})

		t.Run("Ok remains Ok", func(t *testing.T) {
			assert.Equals(t, result.MapErr(result.Ok[int](1), func(it error) error { return fmt.Errorf("error: %w", it) }), result.Ok[int](1))
			assert.Equals(t, result.MapErr(result.Ok[string]("value"), func(it error) error { return fmt.Errorf("error: %w", it) }), result.Ok[string]("value"))
		})
	})
}
