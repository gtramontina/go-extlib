package collections

// Map calls the provided mapper function once for each element in the given
// collection slice, in order, and constructs a new slice from the results.
func Map[From any, To any](collection []From, mapper func(it From) To) []To {
	mapped := make([]To, 0, len(collection))
	for _, element := range collection {
		mapped = append(mapped, mapper(element))
	}

	return mapped
}
