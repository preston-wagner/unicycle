package multithread

import (
	"github.com/preston-wagner/unicycle/channels"
)

// like ChannelForEach, but runs concurrently up to a given limit
func ChannelForEachMultithread[INPUT_TYPE any](input chan INPUT_TYPE, worker func(INPUT_TYPE), threadCount int) {
	workers := make([]func(), 0, threadCount)
	for i := 0; i < threadCount; i++ {
		workers = append(workers, func() {
			channels.ChannelForEach(input, worker)
		})
	}
	AwaitConcurrent(workers...)
}

func ForEachMultithread[INPUT_TYPE any](input []INPUT_TYPE, worker func(INPUT_TYPE), threadCount int) {
	ChannelForEachMultithread(channels.SliceToChannel(input), worker, threadCount)
}
