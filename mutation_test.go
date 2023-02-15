//go:build mutation

package main_test

import (
	"testing"

	"github.com/gtramontina/ooze"
)

func TestMutation(t *testing.T) {
	ooze.Release(
		t,
		ooze.ForceColors(),
		ooze.WithTestCommand("make test.failfast.mutation"),
		ooze.WithMinimumThreshold(0.5),
	)
}
