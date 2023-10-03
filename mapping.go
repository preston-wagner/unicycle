package unicycle

// Mapping accepts a slice of any data type and a mutator function, then returns a slice of that same data with the mutator applied.
// Equivalent to JavaScript's Array.prototype.map()
func Mapping[INPUT_TYPE any, OUTPUT_TYPE any](input []INPUT_TYPE, mutator func(INPUT_TYPE) OUTPUT_TYPE) []OUTPUT_TYPE {
	output := make([]OUTPUT_TYPE, len(input))
	for index, value := range input {
		output[index] = mutator(value)
	}
	return output
}

// like Mapping(), but all mutator functions run in parallel in their own goroutines
func MappingMultithread[INPUT_TYPE any, OUTPUT_TYPE any](input []INPUT_TYPE, mutator func(INPUT_TYPE) OUTPUT_TYPE) []OUTPUT_TYPE {
	pending := Mapping(input, func(value INPUT_TYPE) *Promise[OUTPUT_TYPE] {
		return WrapInPromise(func() (OUTPUT_TYPE, error) {
			return mutator(value), nil
		})
	})
	return Mapping(pending, func(prm *Promise[OUTPUT_TYPE]) OUTPUT_TYPE {
		value, _ := prm.Await()
		return value
	})
}
