package collections_test

import (
	"testing"

	"github.com/gtramontina/go-extlib/collections"
	"github.com/gtramontina/go-extlib/testing/assert"
)

func TestIsEmpty(t *testing.T) {
	assert.True(t, collections.IsEmpty([]int{}))
	assert.False(t, collections.IsEmpty([]int{0}))

	assert.True(t, collections.IsEmpty([]string{}))
	assert.False(t, collections.IsEmpty([]string{""}))

	assert.True(t, collections.IsEmpty([]struct{}{}))
	assert.False(t, collections.IsEmpty([]struct{}{{}}))
}
