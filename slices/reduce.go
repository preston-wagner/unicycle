package slices

// Equivalent to JavaScript's Array.prototype.reduce()
func Reduce[INPUT_TYPE any, OUTPUT_TYPE any](input []INPUT_TYPE, reducer func(OUTPUT_TYPE, INPUT_TYPE) OUTPUT_TYPE, initial OUTPUT_TYPE) OUTPUT_TYPE {
	result := initial
	for _, value := range input {
		result = reducer(result, value)
	}
	return result
}
