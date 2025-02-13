package multithread

import (
	"github.com/preston-wagner/unicycle/channels"
)

// Like ChannelFilter, but runs filters concurrently up to a given limit
// WARNING: unlike ChannelFilter, order is not necessarily preserved
func ChannelFilterMultithread[T any](input chan T, filter func(T) bool, threadCount int) chan T {
	results := make([]chan T, 0, threadCount)
	for i := 0; i < threadCount; i++ {
		results = append(results, channels.ChannelFilter(input, filter))
	}
	return MergeChannels(results)
}
