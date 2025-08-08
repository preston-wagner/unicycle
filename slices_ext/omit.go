package slices_ext

import (
	"github.com/nuvi/unicycle/sets"
)

// Omit accepts two slices of data, and returns a new slice containing only the elements of the first slice that do not appear in the second
// Performance: ~O(n)
func Omit[T comparable](input []T, remove []T) []T {
	remove_set := sets.SetFromSlice(remove)
	return Filter(input, remove_set.Lacks)
}
