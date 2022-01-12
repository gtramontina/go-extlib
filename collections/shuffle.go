package collections

import (
	"math/rand"
	"time"
)

// Shuffle shuffles the given slice using time.Now() as the random seed. It
// returns a new slice containing the shuffled items.
func Shuffle[Type any](input []Type) []Type {
	return RandShuffle(rand.New(rand.NewSource(time.Now().UnixNano())), input)
}

// RandShuffle shuffles the given slice using the given randomizer. It returns a
// new slice containing the shuffled items.
func RandShuffle[Type any](r *rand.Rand, input []Type) []Type {
	output := make([]Type, len(input))
	copy(output, input)
	r.Shuffle(len(input), func(i, j int) { output[i], output[j] = output[j], output[i] })
	return output
}
