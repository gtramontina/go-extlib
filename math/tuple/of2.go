package tuple

// OfTwo is a 2-tuple, also known as a couple, double, ordered pair, two-ple,
// twin, dual, duad, dyad or twosome.
//
// See also: OfOne, OfThree, OfFour, OfFive, OfSix, OfSeven, OfEight, OfNine,
// OfTen.
type OfTwo[A, B any] struct{ tuple }

// Of2 creates a OfTwo 2-tuple.
func Of2[A, B any](a A, b B) OfTwo[A, B] {
	return OfTwo[A, B]{tuple{[]any{a, b}}}
}

// Get1 returns the first element of this tuple.
func (m OfTwo[A, B]) Get1() A { return m.elements[0].(A) }

// Get2 returns the second element of this tuple.
func (m OfTwo[A, B]) Get2() B { return m.elements[1].(B) }
