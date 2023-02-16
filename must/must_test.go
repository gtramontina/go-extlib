package must_test

import (
	"errors"
	"testing"

	"github.com/gtramontina/go-extlib/must"
	"github.com/gtramontina/go-extlib/testing/assert"
)

func TestReturn(t *testing.T) {
	t.Run("returns the result when error is nil", func(t *testing.T) {
		assert.DeepEqual(t, 1, must.Return(1, nil))
		assert.DeepEqual(t, "hi", must.Return("hi", nil))
		assert.DeepEqual(t, []string{}, must.Return([]string{}, nil))

		testFn := func() (string, error) { return "hello", nil }
		assert.DeepEqual(t, "hello", must.Return(testFn()))
	})

	t.Run("panics when error is not nil", func(t *testing.T) {
		assert.PanicsWith(t, func() {
			must.Return(1, errors.New("boom"))
		}, errors.New("boom"))

		assert.PanicsWith(t, func() {
			must.Return("oops", errors.New("crash"))
		}, errors.New("crash"))

		testFn := func() (string, error) { return "", errors.New("borked") }
		assert.PanicsWith(t, func() {
			must.Return(testFn())
		}, errors.New("borked"))
	})
}
