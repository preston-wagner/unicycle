package slices

// Mapping accepts a slice of any data type and a mutator function, then returns a slice of that same data with the mutator applied.
// Equivalent to JavaScript's Array.prototype.map()
func Mapping[INPUT_TYPE any, OUTPUT_TYPE any](input []INPUT_TYPE, mutator func(INPUT_TYPE) OUTPUT_TYPE) []OUTPUT_TYPE {
	output := make([]OUTPUT_TYPE, len(input))
	for index, value := range input {
		output[index] = mutator(value)
	}
	return output
}

// like Mapping(), but accepts a mutator that can return an error, and aborts on the first non-nil error returned by a mutator
func MappingWithError[INPUT_TYPE any, OUTPUT_TYPE any](input []INPUT_TYPE, mutator func(INPUT_TYPE) (OUTPUT_TYPE, error)) ([]OUTPUT_TYPE, error) {
	output := make([]OUTPUT_TYPE, len(input))
	for index, value := range input {
		mutated, err := mutator(value)
		if err != nil {
			return []OUTPUT_TYPE{}, err
		} else {
			output[index] = mutated
		}
	}
	return output, nil
}
