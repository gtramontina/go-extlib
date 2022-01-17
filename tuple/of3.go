package tuple

// OfThree is a 3-tuple, also known as a treble, triplet, triad, ordered triple
// or threesome
//
// See also: OfOne, OfTwo, OfFour, OfFive, OfSix, OfSeven, OfEight, OfNine,
// OfTen.
type OfThree[A, B, C any] struct{ tuple }

// Of3 creates a OfThree 3-tuple.
func Of3[A, B, C any](a A, b B, c C) OfThree[A, B, C] {
	return OfThree[A, B, C]{tuple{[]any{a, b, c}}}
}

// Get1 returns the first element of this tuple.
func (m OfThree[A, B, C]) Get1() A { return m.elements[0].(A) }

// Get2 returns the second element of this tuple.
func (m OfThree[A, B, C]) Get2() B { return m.elements[1].(B) }

// Get3 returns the third element of this tuple.
func (m OfThree[A, B, C]) Get3() C { return m.elements[2].(C) }
