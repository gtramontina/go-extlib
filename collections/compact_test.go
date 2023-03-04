package collections_test

import (
	"testing"

	"github.com/gtramontina/go-extlib/collections"
	"github.com/gtramontina/go-extlib/testing/assert"
)

func ptr[T any](s T) *T { return &s }

func TestCompact(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		compacted := collections.Compact([]*string{})
		assert.DeepEqual(t, compacted, []*string{})
	})

	t.Run("slice with one item and this item is nil", func(t *testing.T) {
		compacted := collections.Compact([]*string{nil})
		assert.DeepEqual(t, compacted, []*string{})
	})

	t.Run("slice with one item and this item is not nil", func(t *testing.T) {
		compacted := collections.Compact([]*int{ptr(1)})
		assert.DeepEqual(t, compacted, []*int{ptr(1)})
	})

	t.Run("slice with many items and some of them are nil", func(t *testing.T) {
		compacted := collections.Compact([]*string{nil, ptr("a"), nil, ptr("b")})
		assert.DeepEqual(t, compacted, []*string{ptr("a"), ptr("b")})
	})
}
