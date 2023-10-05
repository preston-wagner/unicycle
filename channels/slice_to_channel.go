package channels

import "github.com/preston-wagner/unicycle/slices"

// SliceToChannel returns an already-closed channel with the same contents and capacity of the input slice
func SliceToChannel[INPUT_TYPE any](input []INPUT_TYPE) chan INPUT_TYPE {
	output := make(chan INPUT_TYPE, cap(input))
	for _, value := range input {
		output <- value
	}
	close(output)
	return output
}

// ChannelToSlice reads from a channel until it closes, and returns what it read in a slice
// WARNING: this blocks until the input channel closes, so make sure it does close!
func ChannelToSlice[INPUT_TYPE any](input chan INPUT_TYPE) []INPUT_TYPE {
	output := make([]INPUT_TYPE, 0)
	for value := range input {
		output = append(output, value)
	}
	return slices.Trim(output)
}
