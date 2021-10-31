package either

import (
	"fmt"
	"github.com/gtramontina/go-collections/maybe"
	"reflect"
)

type left[L any, R any] struct {
	value L
}

func (left[L, R]) seal() string {
	return "Left"
}

func (l left[L, R]) Equals(other Either[L, R]) bool {
	return reflect.DeepEqual(l, other)
}

func (l left[L, R]) String() string {
	kind := reflect.TypeOf(l.value).String()
	return l.seal() + "[" + kind + "](" + fmt.Sprintf("%+v", l.value) + ")"
}

func (left[L, R]) IsLeft() bool {
	return true
}

func (left[L, R]) IsRight() bool {
	return false
}

func (l left[L, R]) Left() maybe.Maybe[L] {
	return maybe.Some(l.value)
}

func (l left[L, R]) LeftOr(_ L) L {
	return l.value
}

func (l left[L, R]) LeftOrElse(_ func() L) L {
	return l.value
}

func (left[L, R]) Right() maybe.Maybe[R] {
	return maybe.None[R]()
}

func (left[L, R]) RightOr(or R) R {
	return or
}

func (left[L, R]) RightOrElse(orElse func() R) R {
	return orElse()
}

func (l left[L, R]) UnwrapLeft() L {
	return l.value
}

func (left[L, R]) UnwrapRight() R {
	panic("nothing to unwrap right from Left")
}

func (l left[L, R]) Flip() Either[R, L] {
	return Right[R, L](l.value)
}
