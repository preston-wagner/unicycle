package unicycle

// Keys returns a slice containing all the keys of the input map
func Keys[KEY_TYPE comparable, DONTCARE any](input map[KEY_TYPE]DONTCARE) []KEY_TYPE {
	result := make([]KEY_TYPE, len(input))
	index := 0
	for key := range input {
		result[index] = key
		index += 1
	}
	return result
}
