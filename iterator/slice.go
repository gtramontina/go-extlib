package iterator

type slice[T any] struct {
	source []T
	index  int
}

// FromSlice returns an iterator over the given slice.
func FromSlice[T any](source []T) Iterator[T] {
	return &slice[T]{
		source: source,
		index:  0,
	}
}

func (i *slice[T]) HasNext() bool {
	return i.index < len(i.source)
}

func (i *slice[T]) Next() T {
	if i.index >= len(i.source) {
		panic(ErrIteratorEmpty)
	}

	value := i.source[i.index]
	i.index++

	return value
}
