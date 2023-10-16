package multithread

import (
	"github.com/preston-wagner/unicycle/channels"
	"github.com/preston-wagner/unicycle/promises"
	"github.com/preston-wagner/unicycle/slices"
)

// like ChannelForEach, but runs concurrently up to a given limit
func ChannelForEachMultithread[INPUT_TYPE any](input chan INPUT_TYPE, worker func(INPUT_TYPE), threadCount int) {
	promises.AwaitAllSafe(
		slices.Mapping(
			channels.SplitChannel(input, threadCount),
			func(inputChan chan INPUT_TYPE) *promises.SafePromise[bool] {
				return promises.WrapInSafePromise(func() bool {
					channels.ChannelForEach(inputChan, worker)
					return true
				})
			},
		)...,
	)
}

func ForEachMultithread[INPUT_TYPE any](input []INPUT_TYPE, worker func(INPUT_TYPE), threadCount int) {
	ChannelForEachMultithread(channels.SliceToChannel(input), worker, threadCount)
}
