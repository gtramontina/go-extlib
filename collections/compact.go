package collections

// Compact returns a slice with all nil elements removed. The order of the
// elements is preserved.
func Compact[T any](slice []*T) []*T {
	compacted := []*T{}

	for _, item := range slice {
		if item != nil {
			compacted = append(compacted, item)
		}
	}

	return compacted
}
