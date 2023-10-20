package multithread

import (
	"github.com/preston-wagner/unicycle/channels"
	"github.com/preston-wagner/unicycle/slices"
)

// Like ChannelFilter, but runs filters concurrently up to a given limit
// WARNING: unlike ChannelFilter, order is not necessarily preserved
func ChannelFilterMultithread[T any](input chan T, filter func(T) bool, threadCount int) chan T {
	return MergeChannels(
		slices.Mapping(
			SplitChannel(input, threadCount),
			func(inputChan chan T) chan T {
				return channels.ChannelFilter(inputChan, filter)
			},
		),
	)
}
