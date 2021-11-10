package collections

// Chunk splits a collection into chunks of a given size. The given chunk size must be greater than zero, otherwise it
// panics. Example:
//
//	_ = Chunk([]int{1, 2, 3, 4, 5}, 2) == [][]int{{1, 2}, {3, 4}, {5}}
func Chunk[Type any](collection []Type, chunkSize int) [][]Type {
	if chunkSize < 1 {
		panic("chunk size must be greater than 1")
	}

	numberOfItems := len(collection)
	numberOfChunks := (numberOfItems + chunkSize - 1) / chunkSize
	chunks := make([][]Type, numberOfChunks)

	for i := 0; i < numberOfChunks; i++ {
		offset := i * chunkSize
		end := offset + chunkSize
		if end > numberOfItems {
			end = numberOfItems
		}
		chunks[i] = collection[offset:end]
	}

	return chunks
}
