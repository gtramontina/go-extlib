package iterator

// Cycle returns an iterator that cycles through the given items. The returned
// iterator will yield the items in the same order as they were given, and
// will repeat the cycle indefinitely.
func Cycle[T any](items ...T) Iterator[T] {
	return &cycleIterator[T]{items: items}
}

type cycleIterator[T any] struct {
	items []T
}

func (i *cycleIterator[T]) HasNext() bool {
	return len(i.items) > 0
}

func (i *cycleIterator[T]) Next() T {
	if !i.HasNext() {
		panic(ErrIteratorEmpty)
	}

	item := i.items[0]
	i.items = i.items[1:]
	i.items = append(i.items, item)

	return item
}

func (i *cycleIterator[T]) Collect() []T {
	var collected []T
	for i.HasNext() {
		collected = append(collected, i.Next())
	}

	return collected
}
