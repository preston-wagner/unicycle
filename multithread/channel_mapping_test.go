package multithread

import (
	"reflect"
	"testing"

	"github.com/nuvi/unicycle/channels"
	"github.com/nuvi/unicycle/sets"
	"github.com/nuvi/unicycle/test_ext"
)

func TestChannelMappingMultithread(t *testing.T) {
	result := channels.ChannelToSlice(ChannelMappingMultithread(channels.SliceToChannel([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}), test_ext.ToString, 10))
	if !reflect.DeepEqual(sets.SetFromSlice(result), sets.SetFromSlice([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"})) { // testing equality with sets since order of output is not preserved
		t.Errorf("ChannelMappingMultithread() returned unexpected %s", result)
	}

	if len(channels.ChannelToSlice(ChannelMappingMultithread(channels.SliceToChannel([]int{}), test_ext.ToString, 10))) != 0 {
		t.Error("ChannelMappingMultithread() with a closed channel should return a closed channel")
	}
}
