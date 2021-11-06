package result

import (
	"fmt"
	"github.com/gtramontina/go-collections/maybe"
	"reflect"
)

type err[Type any] struct {
	value error
}

func (err[Type]) seal() string {
	return "Err"
}

func (e err[Type]) Equals(other Result[Type]) bool {
	return reflect.DeepEqual(e, other)
}

func (e err[Type]) String() string {
	return e.seal() + "(" + fmt.Sprintf("%+v", e.value) + ")"
}

func (err[Type]) IsOk() bool {
	return false
}

func (err[Type]) IsErr() bool {
	return true
}

func (err[Type]) Ok() maybe.Maybe[Type] {
	return maybe.None[Type]()
}

func (e err[Type]) Unwrap() Type {
	panic(e.value)
}

func (e err[Type]) UnwrapErr() error {
	return e.value
}

func (e err[Type]) UnwrapOr(or Type) Type {
	return or
}

func (e err[Type]) UnwrapOrElse(orElse func() Type) Type {
	return orElse()
}
