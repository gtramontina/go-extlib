package collections

// Every runs the given predicate function against each element in the
// collection. If the result of calling the predicate on all elements is `true`,
// then it returns `true`. If, at least one returns `false`, then it returns
// `false`.
func Every[Type any](collection []Type, predicate func(Type) bool) bool {
	for _, element := range collection {
		if !predicate(element) {
			return false
		}
	}

	return true
}
