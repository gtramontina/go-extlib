package collections

// Take returns the first n elements of a slice. If n is greater than the
// length of the slice, the entire slice is returned.
func Take[Type any](slice []Type, n uint) []Type {
	taken := []Type{}

	for i := uint(0); i < n && i < uint(len(slice)); i++ {
		taken = append(taken, slice[i])
	}

	return taken
}
