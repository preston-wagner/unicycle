package maps

// Values returns a slice containing all the keys of the input map
// Performance: O(n)
func Values[DONTCARE comparable, VALUE_TYPE any](input map[DONTCARE]VALUE_TYPE) []VALUE_TYPE {
	result := make([]VALUE_TYPE, 0, len(input))
	for _, value := range input {
		result = append(result, value)
	}
	return result
}
