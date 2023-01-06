package unicycle

// Every accepts a slice of any data type and a test function, then returns true if all elements in the slice pass the test
// Empty slices always return true
// Equivalent to JavaScript's Array.prototype.every()
func Every[T any](input []T, filter func(T) bool) bool {
	for _, value := range input {
		if !filter(value) {
			return false
		}
	}
	return true
}
