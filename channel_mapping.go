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

// accepts a single input channel and returns a given number of unbuffered child channels that all pull from it
// the child channels close when the parent channel does
func splitChannel[T any](input chan T, splitCount int) []chan T {
	if splitCount < 1 {
		panic("splitChannel argument splitCount must be > 0")
	}
	output := []chan T{}
	for i := 0; i < splitCount; i++ {
		outChan := make(chan T)
		go func() {
			for value := range input {
				outChan <- value
			}
			close(outChan)
		}()
		output = append(output, outChan)
	}
	return output
}

// accepts any number of channels of the same type and returns a single channel that pulls from all of them at once
// the returned channel closes once all the source channels do
func mergeChannels[T any](input []chan T, capacity int) chan T {
	output := make(chan T, capacity)
	go func() {
		AwaitAll(Mapping(input, func(inputChan chan T) *Promise[bool] {
			return WrapInPromise(func() (bool, error) {
				for value := range inputChan {
					output <- value
				}
				return true, nil
			})
		})...)
		close(output)
	}()
	return output
}

// Like ChannelMapping, but runs mutators concurrently up to a given limit
// WARNING: unlike ChannelMapping, order is not necessarily preserved
func ChannelMappingMultithread[INPUT_TYPE any, OUTPUT_TYPE any](input chan INPUT_TYPE, mutator func(INPUT_TYPE) OUTPUT_TYPE, threadCount int) chan OUTPUT_TYPE {
	return mergeChannels(Mapping(splitChannel(input, threadCount), func(inputChan chan INPUT_TYPE) chan OUTPUT_TYPE {
		return ChannelMapping(inputChan, mutator)
	}), cap(input))
}
