package tuple

// OfOne is a 1-tuple, also known as a monuple, single, singleton or monad.
//
// See also: OfTwo, OfThree, OfFour, OfFive, OfSix, OfSeven, OfEight, OfNine,
// OfTen.
type OfOne[A any] struct{ tuple }

// Of1 creates a OfOne 1-tuple.
func Of1[A any](a A) OfOne[A] {
	return OfOne[A]{tuple{[]any{a}}}
}

// Get1 returns the first element of this tuple.
func (m OfOne[A]) Get1() A { return m.elements[0].(A) }
