package unicycle

// Filter accepts a slice of any data type and a filter function, then returns a slice of the data that passes the filter.
// Equivalent to JavaScript's Array.prototype.filter()
func Filter[INPUT_TYPE any](input []INPUT_TYPE, filter func(INPUT_TYPE) bool) []INPUT_TYPE {
	var keepIndexes Queue[int]
	for index, value := range input {
		if filter(value) {
			keepIndexes.Push(index)
		}
	}
	return Mapping(keepIndexes.PopAll(), func(index int) INPUT_TYPE {
		return input[index]
	})
}
