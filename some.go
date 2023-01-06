package unicycle

// Some accepts a slice of any data type and a test function, then returns true if at least 1 element in the slice passes the test
// Empty slices always return false
// Equivalent to JavaScript's Array.prototype.some()
func Some[T any](input []T, filter func(T) bool) bool {
	for _, value := range input {
		if filter(value) {
			return true
		}
	}
	return false
}
