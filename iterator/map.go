package iterator

import "github.com/gtramontina/go-extlib/collections"

// Map calls the provided mapper function once for each element in the given
// iterator, in order, and constructs a new iterator from the results.
func Map[From any, To any](collection Iterator[From], mapper func(it From) To) Iterator[To] {
	return &iterableMap[From, To]{collection: collection, mapper: mapper}
}

type iterableMap[From any, To any] struct {
	collection Iterator[From]
	mapper     func(it From) To
}

func (i *iterableMap[From, To]) HasNext() bool {
	return i.collection.HasNext()
}

func (i *iterableMap[From, To]) Next() To {
	return i.mapper(i.collection.Next())
}

func (i *iterableMap[From, To]) Collect() []To {
	return collections.Map[From, To](i.collection.Collect(), i.mapper)
}
