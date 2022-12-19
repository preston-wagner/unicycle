package unicycle

// Merge takes any number of maps and returns a map containing their combined values, overwriting prior values with later ones
func Merge[KEY_TYPE comparable, VALUE_TYPE any](input ...map[KEY_TYPE]VALUE_TYPE) map[KEY_TYPE]VALUE_TYPE {
	result := make(map[KEY_TYPE]VALUE_TYPE)
	for _, inputMap := range input {
		for key, value := range inputMap {
			result[key] = value
		}
	}
	return result
}
