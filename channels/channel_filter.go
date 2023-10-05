package channels

import "github.com/preston-wagner/unicycle/slices"

// Like Filter, but instead of filtering the values of a slice, filters the values of a channel
// The output channel has the same capacity as the input channel, and is closed when the input channel is
// A single goroutine is spawned and order is preserved
func ChannelFilter[T any](input chan T, filter func(T) bool) chan T {
	keep := make(chan T, cap(input))
	go func() {
		for value := range input {
			if filter(value) {
				keep <- value
			}
		}
		close(keep)
	}()
	return keep
}

// Like ChannelFilter, but runs filters concurrently up to a given limit
// WARNING: unlike ChannelFilter, order is not necessarily preserved
func ChannelFilterMultithread[T any](input chan T, filter func(T) bool, threadCount int) chan T {
	return mergeChannels(slices.Mapping(splitChannel(input, threadCount), func(inputChan chan T) chan T {
		return ChannelFilter(inputChan, filter)
	}), cap(input))
}
