package unicycle

// Like Mapping and Filter at the same time
func MappingFilter[INPUT_TYPE any, OUTPUT_TYPE any](input []INPUT_TYPE, mutatingFilter func(INPUT_TYPE) (OUTPUT_TYPE, bool)) []OUTPUT_TYPE {
	keep := make([]OUTPUT_TYPE, 0, len(input))
	for _, value := range input {
		mutated, ok := mutatingFilter(value)
		if ok {
			keep = append(keep, mutated)
		}
	}
	return Trim(keep)
}

// like MappingFilter(), but all mutating/filter functions run in parallel in their own goroutines
func MappingFilterMultithread[INPUT_TYPE any, OUTPUT_TYPE any](input []INPUT_TYPE, mutatingFilter func(INPUT_TYPE) (OUTPUT_TYPE, bool)) []OUTPUT_TYPE {
	type result struct {
		mutated OUTPUT_TYPE
		ok      bool
	}
	pending := Mapping(input, func(value INPUT_TYPE) *Promise[result] {
		return WrapInPromise(func() (result, error) {
			mutated, ok := mutatingFilter(value)
			return result{
				mutated: mutated,
				ok:      ok,
			}, nil
		})
	})
	return MappingFilter(pending, func(prm *Promise[result]) (OUTPUT_TYPE, bool) {
		res, _ := prm.Await()
		return res.mutated, res.ok
	})
}
