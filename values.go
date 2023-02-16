package unicycle

// Values returns a slice containing all the keys of the input map
func Values[DONTCARE comparable, VALUE_TYPE any](input map[DONTCARE]VALUE_TYPE) []VALUE_TYPE {
	result := make([]VALUE_TYPE, len(input))
	index := 0
	for _, value := range input {
		result[index] = value
		index += 1
	}
	return result
}
