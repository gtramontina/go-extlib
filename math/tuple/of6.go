package tuple

// OfSix is a 6-tuple, also known as a hextuple or hexad.
//
// See also: OfOne, OfTwo, OfThree, OfFour, OfFifth, OfSeven, OfEight, OfNine,
// OfTen.
type OfSix[A, B, C, D, E, F any] struct{ tuple }

// Of6 creates a OfSix 6-tuple.
func Of6[A, B, C, D, E, F any](a A, b B, c C, d D, e E, f F) OfSix[A, B, C, D, E, F] {
	return OfSix[A, B, C, D, E, F]{tuple{[]any{a, b, c, d, e, f}}}
}

// Get1 returns the first element of this tuple.
func (m OfSix[A, B, C, D, E, F]) Get1() A { return m.elements[0].(A) }

// Get2 returns the second element of this tuple.
func (m OfSix[A, B, C, D, E, F]) Get2() B { return m.elements[1].(B) }

// Get3 returns the third element of this tuple.
func (m OfSix[A, B, C, D, E, F]) Get3() C { return m.elements[2].(C) }

// Get4 returns the fourth element of this tuple.
func (m OfSix[A, B, C, D, E, F]) Get4() D { return m.elements[3].(D) }

// Get5 returns the fifth element of this tuple.
func (m OfSix[A, B, C, D, E, F]) Get5() E { return m.elements[4].(E) }

// Get6 returns the sixth element of this tuple.
func (m OfSix[A, B, C, D, E, F]) Get6() F { return m.elements[5].(F) }
