package iterator

import "sync"

// Tee returns two iterators that iterate over the same underlying iterator.
// Both iterators can consume the items independently.
func Tee[T any](it Iterator[T]) (Iterator[T], Iterator[T]) {
	buffer := &teeBuffer[T]{}

	return &teeIterator[T]{it: it, buffer: buffer}, &teeIterator[T]{it: it, buffer: buffer}
}

type teeBuffer[T any] struct {
	sync.RWMutex
	buffer []T
	zero   T
}

func (b *teeBuffer[T]) Append(item T) {
	b.Lock()
	defer b.Unlock()

	b.buffer = append(b.buffer, item)
}

func (b *teeBuffer[T]) Get(index int) (T, bool) {
	b.RLock()
	defer b.RUnlock()

	if index < len(b.buffer) {
		return b.buffer[index], true
	}

	return b.zero, false
}

type teeIterator[T any] struct {
	it     Iterator[T]
	index  int
	buffer *teeBuffer[T]
}

func (i *teeIterator[T]) HasNext() bool {
	if i.index < len(i.buffer.buffer) {
		return true
	}

	return i.it.HasNext()
}

func (i *teeIterator[T]) Next() T {
	if item, ok := i.buffer.Get(i.index); ok {
		i.index++

		return item
	}

	i.buffer.Append(i.it.Next())

	return i.Next()
}

func (i *teeIterator[T]) Collect() []T {
	var collected []T
	for i.HasNext() {
		collected = append(collected, i.Next())
	}

	return collected
}
