package collections

// FoldLeft folds the elements of the collection to a single value by combining
// them using the specified function associating left and the initial value.
// When applying an associative function, there is no difference between this
// and FoldRight. However, if the function is non-associative, the results are
// different due to the order of the associations. See also: FoldRight. Example:
//
//	subtract := func (a, b int) int { return a - b }
//	_ = FoldLeft[int]([]int{1, 2, 3, 4}, subtract, 9) == ((((9 - 1) - 2) - 3) - 4)
func FoldLeft[In any, Out any](collection []In, f func(Out, In) Out, initial Out) Out {
	result := initial
	for _, item := range collection {
		result = f(result, item)
	}

	return result
}

// FoldRight folds the elements of the collection to a single value by combining
// them using the specified function associating right and the initial value.
// When applying an associative function, there is no difference between this
// and FoldLeft. However, if the function is non-associative, the results are
// different due to the order of the associations. See also: FoldLeft. Example:
//
//	subtract := func (a, b int) int { return a - b }
//	_ = FoldRight[int]([]int{1, 2, 3, 4}, subtract, 9) == (1 - (2 - (3 - (4 - 9))))
func FoldRight[In any, Out any](collection []In, f func(In, Out) Out, initial Out) Out {
	result := initial
	for i := len(collection) - 1; i >= 0; i-- {
		result = f(collection[i], result)
	}

	return result
}
