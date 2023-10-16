package channels

// ChannelForEach runs the provided worker on each value passed and blocks until the provided channel closes
func ChannelForEach[INPUT_TYPE any](input chan INPUT_TYPE, worker func(INPUT_TYPE)) {
	for value := range input {
		worker(value)
	}
}
