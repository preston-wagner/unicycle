package slices

// Filter accepts a slice of any data type and a filter function, then returns a slice of the data that passes the filter, preserving the original ordering
// Performance: O(n) (assuming a constant-time filter function)
// https://en.wikipedia.org/wiki/Filter_(higher-order_function)
func Filter[T any](input []T, filter func(T) bool) []T {
	keep := make([]T, 0, len(input))
	for _, value := range input {
		if filter(value) {
			keep = append(keep, value)
		}
	}
	return Trim(keep)
}

// like Filter(), but accepts a filter that can return an error, and aborts on the first non-nil error returned by a filter, returning it
func FilterWithError[T any](input []T, filter func(T) (bool, error)) ([]T, error) {
	keep := make([]T, 0, len(input))
	for _, value := range input {
		ok, err := filter(value)
		if err != nil {
			return []T{}, err
		}
		if ok {
			keep = append(keep, value)
		}
	}
	return Trim(keep), nil
}
