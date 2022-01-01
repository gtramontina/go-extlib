package maybe

// Maybe is a polymorphic type that represents the presence (Some) or absence
// (None) of a value.
type Maybe[Type any] interface {
	// seal is used internally as a way of limiting external implementations
	seal() string

	// Equals allows comparing two Maybe[Type]. Returns true when they are equal
	// to each other; false otherwise.
	Equals(Maybe[Type]) bool

	// String renders itself as a string.
	String() string

	// IsSome returns true if the Maybe[Type] is Some; false otherwise.
	IsSome() bool

	// IsNone returns true if the Maybe[Type] is None; false otherwise.
	IsNone() bool

	// Unwrap returns the wrapped value, if Some, panics if None. See also:
	// UnwrapOr, UnwrapOrElse.
	Unwrap() Type

	// UnwrapOr returns the wrapped value, if Some, or the given default value.
	// The usage of this function forces an early evaluation of the default
	// value, in contrast to UnwrapOrElse. See also: Unwrap,  UnwrapOrElse.
	UnwrapOr(Type) Type

	// UnwrapOrElse returns the wrapped value, if Some, or evaluates the given
	// function which provides a default.
	// The usage of this function lazily evaluates the default value, in
	// contrast to UnwrapOr. See also: Unwrap, UnwrapOr.
	UnwrapOrElse(func() Type) Type
}

// Some wraps the given value with Maybe[Type]. It represents the presence of
// the given value. See also: None, Of.
func Some[Type any](value Type) Maybe[Type] {
	return some[Type]{value}
}

// None represents to absence of a value of the given type. See also None, Of.
func None[Type any]() Maybe[Type] {
	return none[Type]{}
}

// Of wraps the given value with Maybe[Type]. If the given value is nil,
// None[Type] will be returned. Otherwise, Some[Type] is returned. This function
// is useful when the source or state of the given value is unknown. See also:
// Some, None.
func Of[Type any](value any) Maybe[Type] {
	if value != nil {
		return some[Type]{value.(Type)}
	}
	return none[Type]{}
}

// Match pattern-matches on the given Maybe[Type] and returns the result of the
// matching function. If `maybe` is Some[Type], then `whenSome` is evaluated and
// given the underlying value. If it is None[Type], then `whenNone` is
// evaluated.
func Match[From any, To any](maybe Maybe[From], whenSome func(From) To, whenNone func() To) To {
	if it, ok := maybe.(some[From]); ok {
		return whenSome(it.value)
	}
	return whenNone()
}

// Map allows converting the given `maybe`'s value into a different type via the
// `mapper` function. The result of `mapper` is then wrapped with Maybe[Type],
// or None[Type] if value is absent. See also: FlatMap.
func Map[From any, To any](maybe Maybe[From], mapper func(From) To) Maybe[To] {
	if it, ok := maybe.(some[From]); ok {
		return Of[To](mapper(it.value))
	}
	return None[To]()
}

// FlatMap allows converting given `maybe`'s value into a different type via the
// `mapper` function. The result of `mapper` is then wrapped with Maybe[Type] if
// not yet a Maybe[Type], or None[Type] if value is absent. See also: Map.
func FlatMap[From any, To any](maybe Maybe[From], mapper func(From) any) Maybe[To] {
	if it, ok := maybe.(some[From]); ok {
		result := mapper(it.value)
		if maybeResult, ok := result.(Maybe[To]); ok {
			return maybeResult
		}
		return Of[To](result)
	}
	return None[To]()
}
