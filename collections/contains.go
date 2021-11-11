package collections

// Contains determines if a collection contains a value.
func Contains[Type comparable](collection []Type, value Type) bool {
	for _, item := range collection {
		if item == value {
			return true
		}
	}

	return false
}
