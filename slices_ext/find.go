package slices_ext

import "github.com/nuvi/unicycle/defaults"

// Find accepts a slice of any data type and a filter function, then returns the first item that passes the filter, or false if none do
// Performance: O(n) (assuming a constant-time filter function)
func Find[T any](input []T, filter func(T) bool) (T, bool) {
	for _, value := range input {
		if filter(value) {
			return value, true
		}
	}
	return defaults.ZeroValue[T](), false
}
