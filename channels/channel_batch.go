package channels

// ChannelBatch accepts a channel and returns a channel of slices representing batches of data from the input channel.
// This is useful for applications like query batching.
// ChannelBatch prioritizes filling up a batch, but if a receiver is waiting on the output batch and no new inputs are available, a partial batch will be sent instead.
// When the input channel is closed, the batch will be closed out and sent, and then the output channel will be closed as well.
func ChannelBatch[T any](input chan T, batchSize int) chan []T {
	output := make(chan []T)
	go func() {
		batch := make([]T, 0, batchSize)
	BatchLoop:
		for {
			if len(batch) == 0 {
				value, ok := <-input
				if !ok {
					break BatchLoop
				}
				batch = append(batch, value)
			} else if len(batch) < batchSize {
				value, ok := NonBlockingRead(input)
				if ok {
					batch = append(batch, value)
				} else {
					select {
					case value, ok := <-input:
						if !ok {
							break BatchLoop
						}
						batch = append(batch, value)
					case output <- batch:
						batch = make([]T, 0, batchSize)
					}
				}
			} else {
				output <- batch
				batch = make([]T, 0, batchSize)
			}
		}
		if len(batch) > 0 {
			output <- batch
		}
		close(output)
	}()
	return output
}

// ChannelDebatch is the inverse of ChannelBatch; accepts a channel of slices, and splits each batch slice up into its constituent values.
func ChannelDebatch[T any](input chan []T) chan T {
	output := make(chan T)
	go func() {
		for chunk := range input {
			for _, value := range chunk {
				output <- value
			}
		}
		close(output)
	}()
	return output
}
