package collections

// Filter calls the provided predicate function once for each element in the
// given collection slice, and constructs a new slice of all the elements for
// which the predicate returns true.
func Filter[Type any](collection []Type, predicate func(Type) bool) []Type {
	filtered := []Type{}

	for _, element := range collection {
		if predicate(element) {
			filtered = append(filtered, element)
		}
	}

	return filtered
}
