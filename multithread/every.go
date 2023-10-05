package multithread

import "github.com/preston-wagner/unicycle/channels"

// Like Every, but runs tests concurrently up to a given limit
func EveryMultithread[T any](input []T, test func(T) bool, threadCount int) bool {
	return channels.ChannelEveryMultithread(channels.SliceToChannel(input), test, threadCount)
}
