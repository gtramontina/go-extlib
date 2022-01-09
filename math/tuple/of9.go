package tuple

// OfNine is an 9-tuple, also known as a nonad or ennead.
//
// See also: OfOne, OfTwo, OfThree, OfFour, OfFifth, OfSix, OfSeven, OfEight,
// OfTen.
type OfNine[A, B, C, D, E, F, G, H, I any] struct{ tuple }

// Of9 creates a OfNine 9-tuple.
func Of9[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I) OfNine[A, B, C, D, E, F, G, H, I] {
	return OfNine[A, B, C, D, E, F, G, H, I]{tuple{[]any{a, b, c, d, e, f, g, h, i}}}
}

// Get1 returns the first element of this tuple.
func (m OfNine[A, B, C, D, E, F, G, H, I]) Get1() A { return m.elements[0].(A) }

// Get2 returns the second element of this tuple.
func (m OfNine[A, B, C, D, E, F, G, H, I]) Get2() B { return m.elements[1].(B) }

// Get3 returns the third element of this tuple.
func (m OfNine[A, B, C, D, E, F, G, H, I]) Get3() C { return m.elements[2].(C) }

// Get4 returns the fourth element of this tuple.
func (m OfNine[A, B, C, D, E, F, G, H, I]) Get4() D { return m.elements[3].(D) }

// Get5 returns the fifth element of this tuple.
func (m OfNine[A, B, C, D, E, F, G, H, I]) Get5() E { return m.elements[4].(E) }

// Get6 returns the sixth element of this tuple.
func (m OfNine[A, B, C, D, E, F, G, H, I]) Get6() F { return m.elements[5].(F) }

// Get7 returns the seventh element of this tuple.
func (m OfNine[A, B, C, D, E, F, G, H, I]) Get7() G { return m.elements[6].(G) }

// Get8 returns the eighth element of this tuple.
func (m OfNine[A, B, C, D, E, F, G, H, I]) Get8() H { return m.elements[7].(H) }

// Get9 returns the ninth element of this tuple.
func (m OfNine[A, B, C, D, E, F, G, H, I]) Get9() I { return m.elements[8].(I) }
