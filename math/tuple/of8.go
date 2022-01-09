package tuple

// OfEight is an 8-tuple, also known as an octa, octet, octad or octuplet.
//
// See also: OfOne, OfTwo, OfThree, OfFour, OfFifth, OfSix, OfSeven, OfNine,
// OfTen.
type OfEight[A, B, C, D, E, F, G, H any] struct{ tuple }

// Of8 creates a OfSeven 8-tuple.
func Of8[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, h H) OfEight[A, B, C, D, E, F, G, H] {
	return OfEight[A, B, C, D, E, F, G, H]{tuple{[]any{a, b, c, d, e, f, g, h}}}
}

// Get1 returns the first element of this tuple.
func (m OfEight[A, B, C, D, E, F, G, H]) Get1() A { return m.elements[0].(A) }

// Get2 returns the second element of this tuple.
func (m OfEight[A, B, C, D, E, F, G, H]) Get2() B { return m.elements[1].(B) }

// Get3 returns the third element of this tuple.
func (m OfEight[A, B, C, D, E, F, G, H]) Get3() C { return m.elements[2].(C) }

// Get4 returns the fourth element of this tuple.
func (m OfEight[A, B, C, D, E, F, G, H]) Get4() D { return m.elements[3].(D) }

// Get5 returns the fifth element of this tuple.
func (m OfEight[A, B, C, D, E, F, G, H]) Get5() E { return m.elements[4].(E) }

// Get6 returns the sixth element of this tuple.
func (m OfEight[A, B, C, D, E, F, G, H]) Get6() F { return m.elements[5].(F) }

// Get7 returns the seventh element of this tuple.
func (m OfEight[A, B, C, D, E, F, G, H]) Get7() G { return m.elements[6].(G) }

// Get8 returns the eighth element of this tuple.
func (m OfEight[A, B, C, D, E, F, G, H]) Get8() H { return m.elements[7].(H) }
