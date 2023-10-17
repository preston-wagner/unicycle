package slices

// KeyBy accepts a slice of data and key generator function, and returns a map correlating each value to the key the function generated.
// In the event of key conflicts, the last value is kept.
// Performance: O(n*log(n))
func KeyBy[KEY_TYPE comparable, VALUE_TYPE any](input []VALUE_TYPE, keyGenerator func(VALUE_TYPE) KEY_TYPE) map[KEY_TYPE]VALUE_TYPE {
	output := map[KEY_TYPE]VALUE_TYPE{}
	for _, value := range input {
		output[keyGenerator(value)] = value
	}
	return output
}
