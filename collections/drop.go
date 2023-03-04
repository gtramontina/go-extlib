package collections

// Drop returns a slice with the first n elements removed. If n is greater than
// the length of the slice, an empty slice is returned.
func Drop[T any](slice []T, n uint) []T {
	if n > uint(len(slice)) {
		return []T{}
	}

	return slice[n:]
}
