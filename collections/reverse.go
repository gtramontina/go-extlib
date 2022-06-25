package collections

// Reverse reverses the order of the elements in the given collection.
func Reverse[Type any](collection []Type) []Type {
	reversed := make([]Type, len(collection))
	for i, v := range collection {
		reversed[len(collection)-i-1] = v
	}

	return reversed
}
