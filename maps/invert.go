package maps

// Invert accepts a map and returns a new map with the keys and values swapped. If multiple values conflict, the last one returned by range will be kept
// Performance: O(n*log(n))
func Invert[KEY_TYPE comparable, VALUE_TYPE comparable](input map[KEY_TYPE]VALUE_TYPE) map[VALUE_TYPE]KEY_TYPE {
	output := map[VALUE_TYPE]KEY_TYPE{}
	for key, value := range input {
		output[value] = key
	}
	return output
}
