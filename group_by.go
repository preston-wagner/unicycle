package unicycle

// Equivalent to lodash's _.groupBy()
func GroupBy[KEY_TYPE comparable, VALUE_TYPE any](input []VALUE_TYPE, keyGenerator func(VALUE_TYPE) KEY_TYPE) map[KEY_TYPE][]VALUE_TYPE {
	output := map[KEY_TYPE][]VALUE_TYPE{}
	for _, value := range input {
		key := keyGenerator(value)
		_, ok := output[key]
		if !ok {
			output[key] = []VALUE_TYPE{value}
		} else {
			output[key] = append(output[key], value)
		}
	}
	return output
}
