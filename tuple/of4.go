package tuple

// OfFour is a 4-tuple, also known as a quad, tetrad, quartet or quadruplet
//
// See also: OfOne, OfTwo, OfThree, OfFive, OfSix, OfSeven, OfEight, OfNine,
// OfTen.
type OfFour[A, B, C, D any] struct{ tuple }

// Of4 creates a OfFour 4-tuple.
func Of4[A, B, C, D any](a A, b B, c C, d D) OfFour[A, B, C, D] {
	return OfFour[A, B, C, D]{tuple{[]any{a, b, c, d}}}
}

// Get1 returns the first element of this tuple.
func (m OfFour[A, B, C, D]) Get1() A { return m.elements[0].(A) }

// Get2 returns the second element of this tuple.
func (m OfFour[A, B, C, D]) Get2() B { return m.elements[1].(B) }

// Get3 returns the third element of this tuple.
func (m OfFour[A, B, C, D]) Get3() C { return m.elements[2].(C) }

// Get4 returns the fourth element of this tuple.
func (m OfFour[A, B, C, D]) Get4() D { return m.elements[3].(D) }
