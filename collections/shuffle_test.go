package collections_test

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/gtramontina/go-extlib/collections"
	"github.com/gtramontina/go-extlib/testing/assert"
)

func TestShuffle(t *testing.T) {
	t.Run("with given random", func(t *testing.T) {
		r := rand.New(rand.NewSource(-1))
		assert.DeepEqual(t, collections.RandShuffle(r, []int{}), []int{})
		assert.DeepEqual(t, collections.RandShuffle(r, []int{1}), []int{1})
		assert.DeepEqual(t, collections.RandShuffle(r, []int{1, 2}), []int{2, 1})
		assert.DeepEqual(t, collections.RandShuffle(r, []int{1, 2, 3}), []int{2, 3, 1})
		assert.DeepEqual(t, collections.RandShuffle(r, []int{1, 2, 3, 4}), []int{4, 1, 2, 3})

		t.Run("does not mutate the original slice", func(t *testing.T) {
			original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
			_ = collections.RandShuffle(r, original)
			assert.DeepEqual(t, original, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
		})
	})

	t.Run("with default random", func(t *testing.T) {
		assert.DeepEqual(t, collections.Shuffle([]int{}), []int{})
		assert.DeepEqual(t, collections.Shuffle([]int{1}), []int{1})

		missed := 0
		numberOfRuns := 1000
		acceptableMissRate := 1e-10
		testCollection := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for i := 0; i < numberOfRuns; i++ {
			if reflect.DeepEqual(collections.Shuffle(testCollection), collections.Shuffle(testCollection)) {
				missed++
			}
		}
		missRate := float64(missed) / float64(numberOfRuns)

		assert.True(t, missRate <= acceptableMissRate, fmt.Sprintf("expected miss rate to <= %f, but got %f", acceptableMissRate, missRate))

		t.Run("does not mutate the original slice", func(t *testing.T) {
			original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
			_ = collections.Shuffle(original)
			assert.DeepEqual(t, original, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
		})
	})
}
