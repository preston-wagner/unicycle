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
