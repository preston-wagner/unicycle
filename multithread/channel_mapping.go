package multithread

import (
	"github.com/preston-wagner/unicycle/channels"
)

// Like ChannelMapping, but runs mutators concurrently up to a given limit
// WARNING: unlike ChannelMapping, order is not necessarily preserved
func ChannelMappingMultithread[INPUT_TYPE any, OUTPUT_TYPE any](input chan INPUT_TYPE, mutator func(INPUT_TYPE) OUTPUT_TYPE, threadCount int) chan OUTPUT_TYPE {
	results := make([]chan OUTPUT_TYPE, 0, threadCount)
	for i := 0; i < threadCount; i++ {
		results = append(results, channels.ChannelMapping(input, mutator))
	}
	return MergeChannels(results)
}
