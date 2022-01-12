package collections

// Some runs the given predicate function against each element in the
// collection. If the result of calling the predicate on at least one of the
// elements is `true`, then it returns `true`. If no elements satisfied the
// predicate, returns `false`.
func Some[Type any](collection []Type, predicate func(Type) bool) bool {
	for _, element := range collection {
		if predicate(element) {
			return true
		}
	}

	return false
}
