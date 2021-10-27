package maybe

import (
	"fmt"
	"reflect"
)

type Maybe[Type any] interface {
	Equals(Maybe[Type]) bool
	String() string
	Unwrap() Type
	UnwrapOr(Type) Type
}

func Some[Type any](value Type) Maybe[Type] { return some[Type]{value} }

func None[Type any]() Maybe[Type] { return none[Type]{} }

func Of[Type any](value interface{}) Maybe[Type] {
	if value != nil {
		return some[Type]{value.(Type)}
	}
	return none[Type]{}
}

func Match[From any, To any](maybe Maybe[From], whenSome func(From) To, whenNone func() To) To {
	if it, ok := maybe.(some[From]); ok {
		return whenSome(it.value)
	}
	return whenNone()
}

func Map[From any, To any](maybe Maybe[From], mapper func(From) To) Maybe[To] {
	if it, ok := maybe.(some[From]); ok {
		return Of[To](mapper(it.value))
	}
	return None[To]()
}

func FlatMap[From any, To any](maybe Maybe[From], mapper func(From) Maybe[To]) Maybe[To] {
	if it, ok := maybe.(some[From]); ok {
		return mapper(it.value)
	}
	return None[To]()
}

// ---

type some[Type any] struct{ value Type }

func (s some[Type]) Equals(other Maybe[Type]) bool {
	if otherSome, ok := other.(some[Type]); ok {
		return reflect.DeepEqual(s.value, otherSome.value)
	}
	return false
}

func (s some[Type]) String() string {
	kind := reflect.TypeOf(s.value).String()
	return "Some[" + kind + "](" + fmt.Sprintf("%+v", s.value) + ")"
}

func (s some[Type]) Unwrap() Type {
	return s.value
}

func (s some[Type]) UnwrapOr(_ Type) Type {
	return s.value
}

type none[Type any] struct{}

func (none[Type]) Equals(other Maybe[Type]) bool {
	_, ok := other.(none[Type])
	return ok
}

func (none[Type]) String() string {
	return "None()"
}

func (none[Type]) Unwrap() Type {
	panic("nothing to unwrap from None()")
}

func (none[Type]) UnwrapOr(or Type) Type {
	return or
}
