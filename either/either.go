package either

import "github.com/gtramontina/go-extlib/maybe"

// Either is a container for a value of one of two possible types (a disjoint
// union): Left or Right.
type Either[L any, R any] interface {
	// seal is used internally as a way of limiting external implementations. It
	// marks the container as being in one of the two possible states.
	seal() string

	// Equals checks if the container is equal to another container of the
	// same type.
	Equals(Either[L, R]) bool

	// String returns a string representation of the container.
	String() string

	// IsLeft returns true if the container is in the Left state. See also:
	// IsRight.
	IsLeft() bool

	// IsRight returns true if the container is in the Right state. See also:
	// IsLeft.
	IsRight() bool

	// Left returns a maybe.Maybe representation of the container. If it is in
	// the Left state, the value is returned wrapped in a maybe.Some. If it is
	// in the Right state, a maybe.None is returned. See also: LeftOr,
	// LeftOrElse, Right, RightOr, RightOrElse.
	Left() maybe.Maybe[L]

	// LeftOr returns the value if the container is in the Left state, or the
	// given default value. See also: Left, LeftOrElse, Right, RightOr,
	// RightOrElse.
	LeftOr(L) L

	// LeftOrElse returns the value if the container is in the Left state, or
	// the result of calling the given function. See also: Left, LeftOr, Right,
	// RightOr, RightOrElse.
	LeftOrElse(func() L) L

	// Right returns a maybe.Maybe representation of the container. If it is in
	// the Right state, the value is returned wrapped in a maybe.Some. If it is
	// in the Left state, a maybe.None is returned. See also: Left, LeftOr,
	// LeftOrElse, RightOr, RightOrElse.
	Right() maybe.Maybe[R]

	// RightOr returns the value if the container is in the Right state, or the
	// given default value. See also: Left, LeftOr, LeftOrElse, Right,
	// RightOrElse.
	RightOr(R) R

	// RightOrElse returns the value if the container is in the Right state, or
	// the result of calling the given function. See also: Left, LeftOr,
	// LeftOrElse, Right, RightOr.
	RightOrElse(func() R) R

	// UnwrapLeft returns the value if the container is in the Left state, or
	// panics. See also: UnwrapRight.
	UnwrapLeft() L

	// UnwrapRight returns the value if the container is in the Right state, or
	// panics. See also: UnwrapLeft.
	UnwrapRight() R

	// Flip returns a new container with the state of the current container
	// reversed.
	Flip() Either[R, L]
}

// Left returns a new container with the given value in the Left state. See
// also: Right.
func Left[L any, R any](value L) Either[L, R] {
	return left[L, R]{value}
}

// Right returns a new container with the given value in the Right state. See
// also: Left.
func Right[L any, R any](value R) Either[L, R] {
	return right[L, R]{value}
}

// Match pattern-matches on the given `either` and returns the result of the
// matching function. If `either` is Left[L, R], then `whenLeft` is evaluated
// and given the underlying value. If it is Right[L, R], then `whenRight` is
// evaluated and given the underlying value.
func Match[L any, R any, Out any](either Either[L, R], whenLeft func(L) Out, whenRight func(R) Out) Out {
	if either.IsLeft() {
		return whenLeft(either.(left[L, R]).value)
	}

	return whenRight(either.(right[L, R]).value)
}

// Map maps `mapper` over the contained value of `either` and return the result
// in the corresponding state: either Left or Right. See also: MapLeft,
// MapRight.
func Map[T any, Out any](either Either[T, T], mapper func(T) Out) Either[Out, Out] {
	if either.IsLeft() {
		return Left[Out, Out](mapper(either.(left[T, T]).value))
	}

	return Right[Out, Out](mapper(either.(right[T, T]).value))
}

// MapLeft applies the function `mapper` on the value in the Left variant, if it
// is the current state, re-wrapping the result in Left. See also: Map,
// MapRight.
func MapLeft[L any, R any, Out any](either Either[L, R], mapper func(L) Out) Either[Out, R] {
	if either.IsLeft() {
		return Left[Out, R](mapper(either.(left[L, R]).value))
	}

	return Right[Out, R](either.(right[L, R]).value)
}

// MapRight applies the function `mapper` on the value in the Right variant, if
// it is the current state, re-wrapping the result in Right. See also: Map,
// MapLeft.
func MapRight[L any, R any, Out any](either Either[L, R], mapper func(R) Out) Either[L, Out] {
	if either.IsLeft() {
		return Left[L, Out](either.(left[L, R]).value)
	}

	return Right[L, Out](mapper(either.(right[L, R]).value))
}
