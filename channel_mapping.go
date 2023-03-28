package unicycle

// Like Mapping, but instead of mapping the values of a slice, maps the values of a channel
// The output channel has the same capacity as the input channel, and is closed when the input channel is
func ChannelMapping[INPUT_TYPE any, OUTPUT_TYPE any](input chan INPUT_TYPE, mutator func(INPUT_TYPE) OUTPUT_TYPE) chan OUTPUT_TYPE {
	output := make(chan OUTPUT_TYPE, cap(input))
	go func() {
		for value := range input {
			output <- mutator(value)
		}
		close(output)
	}()
	return output
}

// Like ChannelMapping, but accepts mutator functions that can return errors in addition to the data type
func ChannelMappingFallible[INPUT_TYPE any, OUTPUT_TYPE any](input chan INPUT_TYPE, mutator func(INPUT_TYPE) (OUTPUT_TYPE, error)) chan Promissory[OUTPUT_TYPE] {
	return ChannelMapping(input, func(value INPUT_TYPE) Promissory[OUTPUT_TYPE] {
		result, err := mutator(value)
		return Promissory[OUTPUT_TYPE]{Value: result, Err: err}
	})
}
