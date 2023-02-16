package must

// Return returns the result of a function that returns a value and an error.
// If the error is not nil, it panics.
func Return[T any](result T, err error) T {
	NoError(err)

	return result
}

// NoError panics if the error is not nil.
func NoError(err error) {
	if err != nil {
		panic(err)
	}
}
