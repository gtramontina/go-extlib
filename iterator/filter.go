package iterator

// Filter returns an iterator that filters the items of the given iterator
// using the given predicate. The returned iterator will only yield items
// for which the predicate returns true.
func Filter[T any](iter Iterator[T], predicate func(T) bool) Iterator[T] {
	return &filterIterator[T]{iter: iter, predicate: predicate}
}

type filterIterator[T any] struct {
	iter      Iterator[T]
	next      *T
	predicate func(T) bool
}

func (i *filterIterator[T]) HasNext() bool {
	for i.iter.HasNext() {
		next := i.iter.Next()
		if i.predicate(next) {
			i.next = &next

			return true
		}
	}

	return false
}

func (i *filterIterator[T]) Next() T {
	if i.next != nil {
		next := *i.next
		i.next = nil

		return next
	}

	panic(ErrIteratorEmpty)
}

func (i *filterIterator[T]) Collect() []T {
	var collected []T
	for i.HasNext() {
		collected = append(collected, i.Next())
	}

	return collected
}
