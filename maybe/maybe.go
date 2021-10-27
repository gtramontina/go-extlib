package maybe

type Maybe[Type any] interface {
	seal() string
	Equals(Maybe[Type]) bool
	String() string
	Unwrap() Type
	UnwrapOr(Type) Type
	UnwrapOrElse(func() Type) Type
}

func Some[Type any](value Type) Maybe[Type] {
	return some[Type]{value}
}

func None[Type any]() Maybe[Type] {
	return none[Type]{}
}

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
