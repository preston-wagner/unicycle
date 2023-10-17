package slices

import (
	"github.com/preston-wagner/unicycle/sets"
)

// Unique accepts a slice of data, and returns a new slice containing only the first instance of each unique value.
// Performance: ~O(n)
func Unique[T comparable](input []T) []T {
	set := make(sets.Set[T], len(input))
	output := make([]T, 0, len(input))
	for _, value := range input {
		if !set.Has(value) {
			output = append(output, value)
			set.Add(value)
		}
	}
	return Trim(output)
}
