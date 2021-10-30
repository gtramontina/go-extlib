package maybe

type none[Type any] struct{}

func (none[Type]) seal() string {
	return "None"
}

func (none[Type]) Equals(other Maybe[Type]) bool {
	_, ok := other.(none[Type])
	return ok
}

func (n none[Type]) String() string {
	return n.seal() + "()"
}

func (none[Type]) IsSome() bool {
	return false
}

func (none[Type]) IsNone() bool {
	return true
}

func (none[Type]) Unwrap() Type {
	panic("nothing to unwrap from None()")
}

func (none[Type]) UnwrapOr(or Type) Type {
	return or
}

func (none[Type]) UnwrapOrElse(orElse func() Type) Type {
	return orElse()
}
