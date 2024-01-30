package slices

import "github.com/nuvi/unicycle/defaults"

// Like Mapping and Find at the same time
func MappingFind[INPUT_TYPE any, OUTPUT_TYPE any](input []INPUT_TYPE, mutatingFilter func(INPUT_TYPE) (OUTPUT_TYPE, bool)) (OUTPUT_TYPE, bool) {
	for _, value := range input {
		mutated, ok := mutatingFilter(value)
		if ok {
			return mutated, true
		}
	}
	return defaults.ZeroValue[OUTPUT_TYPE](), false
}
