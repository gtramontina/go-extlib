package either

import (
	"fmt"
	"reflect"

	"github.com/gtramontina/go-extlib/maybe"
)

type right[L any, R any] struct {
	value R
}

func (right[L, R]) seal() string {
	return "Right"
}

func (r right[L, R]) Equals(other Either[L, R]) bool {
	return reflect.DeepEqual(r, other)
}

func (r right[L, R]) String() string {
	kind := reflect.TypeOf(r.value).String()

	return r.seal() + "[" + kind + "](" + fmt.Sprintf("%+v", r.value) + ")"
}

func (right[L, R]) IsLeft() bool {
	return false
}

func (right[L, R]) IsRight() bool {
	return true
}

func (right[L, R]) Left() maybe.Maybe[L] {
	return maybe.None[L]()
}

func (right[L, R]) LeftOr(or L) L {
	return or
}

func (right[L, R]) LeftOrElse(orElse func() L) L {
	return orElse()
}

func (r right[L, R]) Right() maybe.Maybe[R] {
	return maybe.Some(r.value)
}

func (r right[L, R]) RightOr(_ R) R {
	return r.value
}

func (r right[L, R]) RightOrElse(_ func() R) R {
	return r.value
}

func (right[L, R]) UnwrapLeft() L {
	panic("nothing to unwrap left from Right")
}

func (r right[L, R]) UnwrapRight() R {
	return r.value
}

func (r right[L, R]) Flip() Either[R, L] {
	return Left[R, L](r.value)
}
