package maps

// Keys returns a slice containing all the keys of the input map
// Performance: O(n)
func Keys[KEY_TYPE comparable, DONTCARE any](input map[KEY_TYPE]DONTCARE) []KEY_TYPE {
	result := make([]KEY_TYPE, 0, len(input))
	for key := range input {
		result = append(result, key)
	}
	return result
}
