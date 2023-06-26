package unicycle

// ChannelForEach runs the provided worker on each value passed and blocks until the provided channel closes
func ChannelForEach[INPUT_TYPE any](input chan INPUT_TYPE, worker func(INPUT_TYPE)) {
	for value := range input {
		worker(value)
	}
}

// like ChannelForEach, but runs concurrently up to a given limit
func ChannelForEachMultithread[INPUT_TYPE any](input chan INPUT_TYPE, worker func(INPUT_TYPE), threadCount int) {
	// TODO: create an alternate version of Promise for functions that will never return errors
	AwaitAll(Mapping(splitChannel(input, threadCount), func(inputChan chan INPUT_TYPE) *Promise[bool] {
		return WrapInPromise(func() (bool, error) {
			ChannelForEach(inputChan, worker)
			return true, nil
		})
	})...)
}

func ForEachMultithread[INPUT_TYPE any](input []INPUT_TYPE, worker func(INPUT_TYPE), threadCount int) {
	ChannelForEachMultithread(SliceToChannel(input), worker, threadCount)
}
