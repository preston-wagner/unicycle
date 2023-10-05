package aggregate

import (
	"github.com/nuvi/unicycle/sets"
	"github.com/nuvi/unicycle/slices"
)

// Unique returns a copy of a slice with the duplicate values omitted, preserving order based on the first instance of each element
func Unique[T comparable](input []T) []T {
	set := sets.Set[T]{}
	output := make([]T, 0, len(input))
	for _, value := range input {
		if !set.Has(value) {
			output = append(output, value)
			set.Add(value)
		}
	}
	return slices.Trim(output)
}
