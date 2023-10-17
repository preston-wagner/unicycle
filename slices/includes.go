package slices

// Includes tells you whether or not a given slice contains a given value
// Performance: O(n)
func Includes[INPUT_TYPE comparable](input []INPUT_TYPE, value INPUT_TYPE) bool {
	for _, inputValue := range input {
		if inputValue == value {
			return true
		}
	}
	return false
}
