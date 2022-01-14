package collections

// IsEmpty tells whether the given collection is empty.
func IsEmpty[Type any](collection []Type) bool {
	return len(collection) == 0
}
