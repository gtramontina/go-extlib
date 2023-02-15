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
		ooze.Parallel(),
		ooze.WithTestCommand("make test.failfast"),
		ooze.WithMinimumThreshold(0.5),
	)
}
