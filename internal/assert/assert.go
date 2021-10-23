package assert

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// True checks whether the given truth is actually true and fails with the given
// (optional) messages otherwise.
func True(t *testing.T, truth bool, message ...string) {
	t.Helper()
	assert(t, truth, func() string { return strings.Join(message, "\n") })
}

// False checks whether the given truth is not true and fails with the given
// (optional) messages otherwise.
func False(t *testing.T, truth bool, message ...string) {
	t.Helper()
	refute(t, truth, func() string { return strings.Join(message, "\n") })
}

type equaler[Type any] interface{ Equals(Type) bool }

// Equals checks whether the actual value is equal to the expected value by
// comparing them using Equals method. This assertion requires types to
// implement an equaler interface: `Equals(Type) bool`. It fails when actual is
// not equal to expected.
func Equals[Type equaler[Type]](t *testing.T, actual Type, expected Type) {
	t.Helper()
	assert(t, actual.Equals(expected), func() string {
		return strings.Join([]string{
			red("Assertion failed: expected values to be equal."), "",
			bold(blue("Actual:")), fmt.Sprintf("%+v", actual), "",
			bold(blue("Expected:")), fmt.Sprintf("%+v", expected), "",
		}, "\n")
	})
}

// Eq checks whether the actual value is equal to the expected value by
// comparing them using ==. It fails when actual is not equal to expected.
func Eq[Type comparable](t *testing.T, actual Type, expected Type) {
	t.Helper()
	assert(t, actual == expected, func() string {
		return strings.Join([]string{
			red("Assertion failed: expected values to be eq (==)."), "",
			bold(blue("Actual:")), fmt.Sprintf("%+v", actual), "",
			bold(blue("Expected:")), fmt.Sprintf("%+v", expected), "",
		}, "\n")
	})
}

// DeepEqual checks whether the actual value is equal to the expected value by
// comparing them using reflect.DeepEqual. It fails when actual is not equal to
// expected.
func DeepEqual[Type any](t *testing.T, actual Type, expected Type) {
	t.Helper()
	assert(t, reflect.DeepEqual(actual, expected), func() string {
		return strings.Join([]string{
			red("Assertion failed: expected values to be deep equal."), "",
			bold(blue("Actual:")), fmt.Sprintf("[%s] %+v", reflect.TypeOf(actual), actual), "",
			bold(blue("Expected:")), fmt.Sprintf("[%s] %+v", reflect.TypeOf(actual), expected), "",
		}, "\n")
	})
}

// NotEquals checks whether the actual value is not equal to the expected value
// by comparing them using Equals method. This assertion requires types to
// implement an equaler interface: `Equals(Type) bool`. It fails when actual is
// equal to expected.
func NotEquals[Type equaler[Type]](t *testing.T, left Type, right Type) {
	t.Helper()
	refute(t, left.Equals(right), func() string {
		return strings.Join([]string{
			red("Assertion failed: expected values to be different."), "",
			bold(blue("Left:")), fmt.Sprintf("%+v", left), "",
			bold(blue("Right:")), fmt.Sprintf("%+v", right), "",
		}, "\n")
	})
}

// NoError checks whether err is nil. It fails when it is not nil.
func NoError(t *testing.T, err error) {
	t.Helper()
	assert(t, err == nil, func() string {
		return strings.Join([]string{
			red("Assertion failed: expected no error."), "",
			bold(blue("Error: ")), err.Error(), "",
		}, "\n")
	})
}

// ---

func assert(t *testing.T, truth bool, lazyMessage func() string) {
	t.Helper()
	if !truth {
		message := lazyMessage()
		if len(message) == 0 {
			message = "Assertion failed!"
		}

		t.Error(message)
		t.FailNow()
	}
}

func refute(t *testing.T, truth bool, lazyMessage func() string) {
	t.Helper()
	assert(t, !truth, lazyMessage)
}

// ---

func red(text string) string   { return reset("\033[31m" + text) }
func blue(text string) string  { return reset("\033[34m" + text) }
func bold(text string) string  { return reset("\033[1m" + text) }
func reset(text string) string { return text + "\033[0m" }
