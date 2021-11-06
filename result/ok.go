package result

import (
	"fmt"
	"github.com/gtramontina/go-collections/maybe"
	"reflect"
)

type ok[Type any] struct {
	value Type
}

func (ok[Type]) seal() string {
	return "Ok"
}

func (o ok[Type]) Equals(other Result[Type]) bool {
	return reflect.DeepEqual(o, other)
}

func (o ok[Type]) String() string {
	kind := reflect.TypeOf(o.value).String()
	return o.seal() + "[" + kind + "](" + fmt.Sprintf("%+v", o.value) + ")"
}

func (ok[Type]) IsOk() bool {
	return true
}

func (ok[Type]) IsErr() bool {
	return false
}

func (o ok[Type]) Ok() maybe.Maybe[Type] {
	return maybe.Some(o.value)
}

func (o ok[Type]) Unwrap() Type {
	return o.value
}

func (o ok[Type]) UnwrapErr() error {
	panic(o.value)
}

func (o ok[Type]) UnwrapOr(_ Type) Type {
	return o.value
}

func (o ok[Type]) UnwrapOrElse(_ func() Type) Type {
	return o.value
}
