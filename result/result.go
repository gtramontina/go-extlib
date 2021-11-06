package result

import "github.com/gtramontina/go-collections/maybe"

// Result represents a computation that may or may not have succeeded.
type Result[Type any] interface {
	// seal is used internally as a way of limiting external implementations. It
	// marks the container as being in one of the two possible states.
	seal() string

	// Equals returns true if the result is equal to the given result.
	Equals(Result[Type]) bool

	// String returns a string representation of the result.
	String() string

	// IsOk returns true if the result is Ok. See also: IsErr.
	IsOk() bool

	// IsErr returns true if the result is Err. See also: IsOk.
	IsErr() bool

	// Ok returns a maybe.Maybe[Type] representation of the result. If it is Ok,
	// a maybe.Some[Type] is returned, otherwise a maybe.None[Type] is returned.
	Ok() maybe.Maybe[Type]

	// Unwrap returns the value contained in the result. It panics if the result
	// is Err. See also: UnwrapErr, UnwrapOr, UnwrapOrElse.
	Unwrap() Type

	// UnwrapErr returns the error contained in the result. It panics if the
	// result is Ok. See also: Unwrap, UnwrapOr, UnwrapOrElse.
	UnwrapErr() error

	// UnwrapOr returns the value contained in the result or the default value
	// if the result is Err. See also: Unwrap, UnwrapErr, UnwrapOrElse.
	UnwrapOr(Type) Type

	// UnwrapOrElse returns the value contained in the result or the result of
	// calling the function if the result is Err. See also: Unwrap, UnwrapErr,
	// UnwrapOr.
	UnwrapOrElse(func() Type) Type
}

// Ok returns a new Ok result.
func Ok[Type any](value Type) Result[Type] {
	return ok[Type]{value}
}

// Err returns a new Err result.
func Err[Type any](value error) Result[Type] {
	return err[Type]{value}
}

// Of returns a new Ok result if the value is not nil, otherwise an Err result.
// See also: Ok, Err.
func Of[Type any](value Type, err error) Result[Type] {
	if err != nil {
		return Err[Type](err)
	}
	return Ok(value)
}

// Match pattern-matches on the given `result` and returns the output of the
// matching function. If `result` is OK[Type], then `whenOk` is evaluated and
// given the underlying value. If it is Err[Type], then `whenErr` is evaluated
// and given the underlying error.
func Match[Type any, Out any](result Result[Type], whenOk func(Type) Out, whenErr func(error) Out) Out {
	if result.IsOk() {
		return whenOk(result.(ok[Type]).value)
	}
	return whenErr(result.(err[Type]).value)
}

// Map maps a Result[Type] to Result[Out] by applying the given `mapper`
// function to a contained Ok value, leaving an Err value untouched.
func Map[Type any, Out any](result Result[Type], mapper func(Type) Out) Result[Out] {
	if result.IsOk() {
		return Ok[Out](mapper(result.(ok[Type]).value))
	}
	return Err[Out](result.(err[Type]).value)
}

// MapErr maps a Result[Type] to Result[Type] by applying the given `mapper`
// function to a contained Err value, leaving an Ok value untouched.
func MapErr[Type any](result Result[Type], mapper func(error) error) Result[Type] {
	if result.IsOk() {
		return Ok[Type](result.(ok[Type]).value)
	}
	return Err[Type](mapper(result.(err[Type]).value))
}
