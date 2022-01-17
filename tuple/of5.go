package tuple

// OfFive is a 5-tuple, also known as a pentuple, quint or pentad.
//
// See also: OfOne, OfTwo, OfThree, OfFour, OfSix, OfSeven, OfEight, OfNine,
// OfTen.
type OfFive[A, B, C, D, E any] struct{ tuple }

// Of5 creates a OfFive 5-tuple.
func Of5[A, B, C, D, E any](a A, b B, c C, d D, e E) OfFive[A, B, C, D, E] {
	return OfFive[A, B, C, D, E]{tuple{[]any{a, b, c, d, e}}}
}

// Get1 returns the first element of this tuple.
func (m OfFive[A, B, C, D, E]) Get1() A { return m.elements[0].(A) }

// Get2 returns the second element of this tuple.
func (m OfFive[A, B, C, D, E]) Get2() B { return m.elements[1].(B) }

// Get3 returns the third element of this tuple.
func (m OfFive[A, B, C, D, E]) Get3() C { return m.elements[2].(C) }

// Get4 returns the fourth element of this tuple.
func (m OfFive[A, B, C, D, E]) Get4() D { return m.elements[3].(D) }

// Get5 returns the fifth element of this tuple.
func (m OfFive[A, B, C, D, E]) Get5() E { return m.elements[4].(E) }
