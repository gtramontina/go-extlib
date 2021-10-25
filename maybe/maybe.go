package maybe

import (
	"fmt"
	"reflect"
)

type Maybe[Type comparable] interface {
	Equals(Maybe[Type]) bool
	String() string
}

func Some[Type comparable](value Type) Maybe[Type] { return some[Type]{value} }

func None[Type comparable]() Maybe[Type] { return none[Type]{} }

func Of[Type comparable](value interface{}) Maybe[Type] {
	if value != nil {
		return some[Type]{value.(Type)}
	}
	return none[Type]{}
}

func Match[From comparable, To comparable](maybe Maybe[From], whenSome func(From) To, whenNone func() To) To {
	if it, ok := maybe.(some[From]); ok {
		return whenSome(it.value)
	}
	return whenNone()
}

func Map[From comparable, To comparable](maybe Maybe[From], mapper func(From) To) Maybe[To] {
	if it, ok := maybe.(some[From]); ok {
		return Of[To](mapper(it.value))
	}
	return None[To]()
}

func FlatMap[From comparable, To comparable](maybe Maybe[From], mapper func(From) Maybe[To]) Maybe[To] {
	if it, ok := maybe.(some[From]); ok {
		return mapper(it.value)
	}
	return None[To]()
}

// ---

type some[Type comparable] struct{ value Type }

func (s some[Type]) Equals(other Maybe[Type]) bool {
	if otherSome, ok := other.(some[Type]); ok {
		return s.value == otherSome.value
	}
	return false
}

func (s some[Type]) String() string {
	kind := reflect.TypeOf(s.value).String()
	return "Some[" + kind + "](" + fmt.Sprintf("%+v", s.value) + ")"
}

type none[Type comparable] struct{}

func (n none[Type]) Equals(other Maybe[Type]) bool {
	_, ok := other.(none[Type])
	return ok
}

func (s none[Type]) String() string {
	return "None()"
}
