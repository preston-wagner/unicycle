package slices

// Count accepts a slice of any data type and a test function, then returns the number of elements that passed the test
// Performance: O(n) (assuming a constant-time test function)
func Count[T any](input []T, filter func(T) bool) int {
	total := 0
	for _, value := range input {
		if filter(value) {
			total++
		}
	}
	return total
}

// like Count(), but accepts a test that can return an error, and aborts on the first non-nil error returned by a test, returning it
func CountWithError[T any](input []T, filter func(T) (bool, error)) (int, error) {
	total := 0
	for _, value := range input {
		ok, err := filter(value)
		if err != nil {
			return total, err
		}
		if ok {
			total++
		}
	}
	return total, nil
}
