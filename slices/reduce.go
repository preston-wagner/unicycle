package slices

// Reduce accepts a slice of data, an accumulator function, and an initial value, and applies the accumulator function to all the values of the slice, returning the accumulated data.
// Performance: O(n) (assuming a constant-time accumulator function)
// https://en.wikipedia.org/wiki/Fold_(higher-order_function)#On_lists
func Reduce[INPUT_TYPE any, OUTPUT_TYPE any](input []INPUT_TYPE, accumulator func(OUTPUT_TYPE, INPUT_TYPE) OUTPUT_TYPE, initial OUTPUT_TYPE) OUTPUT_TYPE {
	result := initial
	for _, value := range input {
		result = accumulator(result, value)
	}
	return result
}
