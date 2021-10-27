package maybe

import (
	"fmt"
	"reflect"
)

type some[Type any] struct {
	value Type
}

func (some[Type]) seal() string {
	return "Some"
}

func (s some[Type]) Equals(other Maybe[Type]) bool {
	if otherSome, ok := other.(some[Type]); ok {
		return reflect.DeepEqual(s.value, otherSome.value)
	}
	return false
}

func (s some[Type]) String() string {
	kind := reflect.TypeOf(s.value).String()
	return s.seal() + "[" + kind + "](" + fmt.Sprintf("%+v", s.value) + ")"
}

func (s some[Type]) Unwrap() Type {
	return s.value
}

func (s some[Type]) UnwrapOr(_ Type) Type {
	return s.value
}

func (s some[Type]) UnwrapOrElse(_ func() Type) Type {
	return s.value
}
