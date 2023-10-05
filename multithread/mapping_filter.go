package multithread

import "github.com/nuvi/unicycle/slices"

type mappingFilterResult[OUTPUT_TYPE any] struct {
	mutated OUTPUT_TYPE
	ok      bool
}

// like MappingFilter(), but all mutating/filter functions run in parallel in their own goroutines
func MappingFilterMultithread[INPUT_TYPE any, OUTPUT_TYPE any](input []INPUT_TYPE, mutatingFilter func(INPUT_TYPE) (OUTPUT_TYPE, bool)) []OUTPUT_TYPE {
	finished := MappingMultithread(input, func(value INPUT_TYPE) mappingFilterResult[OUTPUT_TYPE] {
		mutated, ok := mutatingFilter(value)
		return mappingFilterResult[OUTPUT_TYPE]{
			mutated: mutated,
			ok:      ok,
		}
	})
	return slices.MappingFilter(finished, func(res mappingFilterResult[OUTPUT_TYPE]) (OUTPUT_TYPE, bool) {
		return res.mutated, res.ok
	})
}

// like MappingFilterWithError(), but all mutating/filter functions run in parallel in their own goroutines
func MappingFilterMultithreadWithError[INPUT_TYPE any, OUTPUT_TYPE any](input []INPUT_TYPE, mutatingFilter func(INPUT_TYPE) (OUTPUT_TYPE, bool, error)) ([]OUTPUT_TYPE, error) {
	finished, err := MappingMultithreadWithError(input, func(value INPUT_TYPE) (mappingFilterResult[OUTPUT_TYPE], error) {
		mutated, ok, err := mutatingFilter(value)
		return mappingFilterResult[OUTPUT_TYPE]{
			mutated: mutated,
			ok:      ok,
		}, err
	})
	if err != nil {
		return []OUTPUT_TYPE{}, err
	}
	return slices.MappingFilter(finished, func(res mappingFilterResult[OUTPUT_TYPE]) (OUTPUT_TYPE, bool) {
		return res.mutated, res.ok
	}), nil
}
