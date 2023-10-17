package multithread

import (
	"github.com/nuvi/unicycle/channels"
	"github.com/nuvi/unicycle/slices"
)

// Like ChannelMapping, but runs mutators concurrently up to a given limit
// WARNING: unlike ChannelMapping, order is not necessarily preserved
func ChannelMappingMultithread[INPUT_TYPE any, OUTPUT_TYPE any](input chan INPUT_TYPE, mutator func(INPUT_TYPE) OUTPUT_TYPE, threadCount int) chan OUTPUT_TYPE {
	return MergeChannels(slices.Mapping(SplitChannel(input, threadCount), func(inputChan chan INPUT_TYPE) chan OUTPUT_TYPE {
		return channels.ChannelMapping(inputChan, mutator)
	}), cap(input))
}
