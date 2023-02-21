package iterator

import "errors"

var ErrIteratorEmpty = errors.New("iterator is empty")

// Iterator is an interface that represents an iterator over a collection of
// elements of type T.
type Iterator[T any] interface {
	// HasNext returns true if the iteration has more elements.
	HasNext() bool

	// Next returns the next element in the iteration. If there are no more
	// elements, it panics with ErrIteratorEmpty.
	Next() T

	// Collect returns a slice containing all the remaining elements in the
	Collect() []T
}
