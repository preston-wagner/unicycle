package slices

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

// like MappingFilter(), but accepts a mutator that can return an error, and aborts on the first non-nil error returned by a mutator
func MappingFilterWithError[INPUT_TYPE any, OUTPUT_TYPE any](input []INPUT_TYPE, mutatingFilter func(INPUT_TYPE) (OUTPUT_TYPE, bool, error)) ([]OUTPUT_TYPE, error) {
	keep := make([]OUTPUT_TYPE, 0, len(input))
	for _, value := range input {
		mutated, ok, err := mutatingFilter(value)
		if err != nil {
			return []OUTPUT_TYPE{}, err
		}
		if ok {
			keep = append(keep, mutated)
		}
	}
	return Trim(keep), nil
}
