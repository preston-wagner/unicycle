package multithread

import (
	"github.com/nuvi/unicycle/channels"
	"github.com/nuvi/unicycle/slices"
)

// like ChannelForEach, but runs concurrently up to a given limit
func ChannelForEachMultithread[INPUT_TYPE any](input chan INPUT_TYPE, worker func(INPUT_TYPE), threadCount int) {
	AwaitConcurrent(
		slices.Mapping(
			SplitChannel(input, threadCount),
			func(inputChan chan INPUT_TYPE) func() {
				return func() {
					channels.ChannelForEach(inputChan, worker)
				}
			},
		)...,
	)
}

func ForEachMultithread[INPUT_TYPE any](input []INPUT_TYPE, worker func(INPUT_TYPE), threadCount int) {
	ChannelForEachMultithread(channels.SliceToChannel(input), worker, threadCount)
}
