package channels

// Like Every, but instead of testing the values of a slice, tests the values of a channel
// Returns false on the first failed test return, or true if the input channel closes without any failures
func ChannelEvery[T any](input chan T, test func(T) bool) bool {
	for value := range input {
		if !test(value) {
			return false
		}
	}
	return true
}
