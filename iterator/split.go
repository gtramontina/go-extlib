package iterator

// Split returns two iterators that split the items of the given iterator
// using the given predicate. The first iterator will only yield items for
// which the predicate returns true, while the second iterator will only
// yield items for which the predicate returns false.
func Split[T any](iter Iterator[T], predicate func(T) bool) (Iterator[T], Iterator[T]) {
	left, right := Tee[T](iter)

	return Filter[T](left, predicate), Filter[T](right, func(item T) bool { return !predicate(item) })
}
