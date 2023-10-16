package multithread

import (
	"github.com/preston-wagner/unicycle/channels"
	"github.com/preston-wagner/unicycle/slices"
)

// Like ChannelMapping, but runs mutators concurrently up to a given limit
// WARNING: unlike ChannelMapping, order is not necessarily preserved
func ChannelMappingMultithread[INPUT_TYPE any, OUTPUT_TYPE any](input chan INPUT_TYPE, mutator func(INPUT_TYPE) OUTPUT_TYPE, threadCount int) chan OUTPUT_TYPE {
	return channels.MergeChannels(slices.Mapping(channels.SplitChannel(input, threadCount), func(inputChan chan INPUT_TYPE) chan OUTPUT_TYPE {
		return channels.ChannelMapping(inputChan, mutator)
	}), cap(input))
}
