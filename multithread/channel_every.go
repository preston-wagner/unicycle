package multithread

import "github.com/preston-wagner/unicycle/channels"

// Like ChannelEvery, but runs tests concurrently up to a given limit
func ChannelEveryMultithread[T any](input chan T, test func(T) bool, threadCount int) bool {
	if threadCount < 1 {
		return false
	}
	return channels.ChannelEvery(ChannelMappingMultithread(input, test, threadCount), func(passed bool) bool { return passed })
}
