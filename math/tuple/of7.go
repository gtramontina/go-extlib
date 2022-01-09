package tuple

// OfSeven is a 7-tuple, also known as a heptuple or heptad.
//
// See also: OfOne, OfTwo, OfThree, OfFour, OfFifth, OfSix, OfEight, OfNine,
// OfTen.
type OfSeven[A, B, C, D, E, F, G any] struct{ tuple }

// Of7 creates a OfSeven 7-tuple.
func Of7[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, g G) OfSeven[A, B, C, D, E, F, G] {
	return OfSeven[A, B, C, D, E, F, G]{tuple{[]any{a, b, c, d, e, f, g}}}
}

// Get1 returns the first element of this tuple.
func (m OfSeven[A, B, C, D, E, F, G]) Get1() A { return m.elements[0].(A) }

// Get2 returns the second element of this tuple.
func (m OfSeven[A, B, C, D, E, F, G]) Get2() B { return m.elements[1].(B) }

// Get3 returns the third element of this tuple.
func (m OfSeven[A, B, C, D, E, F, G]) Get3() C { return m.elements[2].(C) }

// Get4 returns the fourth element of this tuple.
func (m OfSeven[A, B, C, D, E, F, G]) Get4() D { return m.elements[3].(D) }

// Get5 returns the fifth element of this tuple.
func (m OfSeven[A, B, C, D, E, F, G]) Get5() E { return m.elements[4].(E) }

// Get6 returns the sixth element of this tuple.
func (m OfSeven[A, B, C, D, E, F, G]) Get6() F { return m.elements[5].(F) }

// Get7 returns the seventh element of this tuple.
func (m OfSeven[A, B, C, D, E, F, G]) Get7() G { return m.elements[6].(G) }
