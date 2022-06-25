package tuple

// OfTen is an 10-tuple, also known as a decad or decade.
//
// See also: OfOne, OfTwo, OfThree, OfFour, OfFifth, OfSix, OfSeven, OfEight,
// OfNine.
type OfTen[A, B, C, D, E, F, G, H, I, J any] struct{ tuple }

// Of10 creates a OfTen 10-tuple.
func Of10[A, B, C, D, E, F, G, H, I, J any](
	a A, b B, c C, d D, e E, f F, g G, h H, i I, j J,
) OfTen[A, B, C, D, E, F, G, H, I, J] {
	return OfTen[A, B, C, D, E, F, G, H, I, J]{tuple{[]any{a, b, c, d, e, f, g, h, i, j}}}
}

// Get1 returns the first element of this tuple.
func (m OfTen[A, B, C, D, E, F, G, H, I, J]) Get1() A { return m.elements[0].(A) }

// Get2 returns the second element of this tuple.
func (m OfTen[A, B, C, D, E, F, G, H, I, J]) Get2() B { return m.elements[1].(B) }

// Get3 returns the third element of this tuple.
func (m OfTen[A, B, C, D, E, F, G, H, I, J]) Get3() C { return m.elements[2].(C) }

// Get4 returns the fourth element of this tuple.
func (m OfTen[A, B, C, D, E, F, G, H, I, J]) Get4() D { return m.elements[3].(D) }

// Get5 returns the fifth element of this tuple.
func (m OfTen[A, B, C, D, E, F, G, H, I, J]) Get5() E { return m.elements[4].(E) }

// Get6 returns the sixth element of this tuple.
func (m OfTen[A, B, C, D, E, F, G, H, I, J]) Get6() F { return m.elements[5].(F) }

// Get7 returns the seventh element of this tuple.
func (m OfTen[A, B, C, D, E, F, G, H, I, J]) Get7() G { return m.elements[6].(G) }

// Get8 returns the eighth element of this tuple.
func (m OfTen[A, B, C, D, E, F, G, H, I, J]) Get8() H { return m.elements[7].(H) }

// Get9 returns the ninth element of this tuple.
func (m OfTen[A, B, C, D, E, F, G, H, I, J]) Get9() I { return m.elements[8].(I) }

// Get10 returns the tenth element of this tuple.
func (m OfTen[A, B, C, D, E, F, G, H, I, J]) Get10() J { return m.elements[9].(J) }
