package collections

// ReduceLeft reduces the elements of the collection to a single value by
// combining them using the specified function associating left. When applying
// an associative function, there is no difference between this and ReduceRight.
// However, if the function is non-associative, the results are different due to
// the order of the associations. Panics if the collection is empty. See also:
// ReduceRight. Example:
//
//   subtract := func (a, b int) int { return a - b }
//   _ = ReduceLeft[int]([]int{1, 2, 3, 4}, subtract) == (((1 - 2) - 3) - 4)
func ReduceLeft[Type any](collection []Type, reducer func(Type, Type) Type) Type {
	if len(collection) == 0 {
		panic("cannot ReduceLeft an empty slice")
	}

	result := collection[0]
	for _, item := range collection[1:] {
		result = reducer(result, item)
	}

	return result
}

// ReduceRight reduces the elements of the collection to a single value by
// combining them using the specified function associating right. When applying
// an associative function, there is no difference between this and ReduceLeft.
// However, if the function is non-associative, the results are different due to
// the order of the associations. Panics if the collection is empty. See also:
// ReduceLeft. Example:
//
//   subtract := func (a, b int) int { return a - b }
//   _ = ReduceRight[int]([]int{1, 2, 3, 4}, subtract) == (1 - (2 - (3 - 4)))
func ReduceRight[Type any](collection []Type, reducer func(Type, Type) Type) Type {
	if len(collection) == 0 {
		panic("cannot ReduceRight an empty slice")
	}

	result := collection[len(collection)-1]
	for i := len(collection) - 2; i >= 0; i-- { //nolint:gomnd // 2 -> first-to-last position
		result = reducer(collection[i], result)
	}

	return result
}
