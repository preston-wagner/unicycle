package unicycle

func ChannelBatch[T any](input chan T, batchSize, buffer int) chan []T {
	output := make(chan []T, buffer)
	go func() {
		batch := make([]T, 0, batchSize)
		for value := range input {
			batch = append(batch, value)
			if (len(batch) % batchSize) == 0 {
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

func ChannelDebatch[T any](input chan []T, buffer int) chan T {
	output := make(chan T, buffer)
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
