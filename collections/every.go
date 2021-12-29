package collections

func Every[Type any](collection []Type, predicate func(Type) bool) bool {
	for _, element := range collection {
		if !predicate(element) {
			return false
		}
	}

	return true
}
