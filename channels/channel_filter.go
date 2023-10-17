package channels

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
