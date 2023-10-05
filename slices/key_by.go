package slices

// Equivalent to lodash's _.keyBy()
// In the event of conflicts, the last value is kept
func KeyBy[KEY_TYPE comparable, VALUE_TYPE any](input []VALUE_TYPE, keyGenerator func(VALUE_TYPE) KEY_TYPE) map[KEY_TYPE]VALUE_TYPE {
	output := map[KEY_TYPE]VALUE_TYPE{}
	for _, value := range input {
		output[keyGenerator(value)] = value
	}
	return output
}
