package multithread

import "github.com/preston-wagner/unicycle/slices_ext"

// like slices_ext.MappingFilter(), but all mutating/filter functions run in parallel in their own goroutines
func MappingFilterMultithread[INPUT_TYPE any, OUTPUT_TYPE any](input []INPUT_TYPE, mutatingFilter func(INPUT_TYPE) (OUTPUT_TYPE, bool)) []OUTPUT_TYPE {
	finished := MappingMultithread(input, func(value INPUT_TYPE) filterResult[OUTPUT_TYPE] {
		mutated, ok := mutatingFilter(value)
		return filterResult[OUTPUT_TYPE]{
			value: mutated,
			ok:    ok,
		}
	})
	return slices_ext.MappingFilter(finished, func(res filterResult[OUTPUT_TYPE]) (OUTPUT_TYPE, bool) {
		return res.value, res.ok
	})
}

// like slices_ext.MappingFilterWithError(), but all mutating/filter functions run in parallel in their own goroutines
func MappingFilterMultithreadWithError[INPUT_TYPE any, OUTPUT_TYPE any](input []INPUT_TYPE, mutatingFilter func(INPUT_TYPE) (OUTPUT_TYPE, bool, error)) ([]OUTPUT_TYPE, error) {
	finished, err := MappingMultithreadWithError(input, func(value INPUT_TYPE) (filterResult[OUTPUT_TYPE], error) {
		mutated, ok, err := mutatingFilter(value)
		return filterResult[OUTPUT_TYPE]{
			value: mutated,
			ok:    ok,
		}, err
	})
	if err != nil {
		return []OUTPUT_TYPE{}, err
	}
	return slices_ext.MappingFilter(finished, func(res filterResult[OUTPUT_TYPE]) (OUTPUT_TYPE, bool) {
		return res.value, res.ok
	}), nil
}
