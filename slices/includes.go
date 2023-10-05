package slices

// Equivalent to JavaScript's Array.prototype.includes()
func Includes[INPUT_TYPE comparable](input []INPUT_TYPE, value INPUT_TYPE) bool {
	for _, inputValue := range input {
		if inputValue == value {
			return true
		}
	}
	return false
}
