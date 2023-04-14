package unicycle

import (
	"reflect"
	"testing"
)

func TestChannelMapping(t *testing.T) {
	result := ChannelToSlice(ChannelMapping(SliceToChannel([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}), toString))
	if !reflect.DeepEqual(result, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}) {
		t.Errorf("ChannelMapping() returned unexpected %s", result)
	}

	if len(ChannelToSlice(ChannelMapping(SliceToChannel([]int{}), toString))) != 0 {
		t.Error("ChannelMapping with a closed channel should return a closed channel")
	}
}

func TestChannelMappingMultithread(t *testing.T) {
	result := ChannelToSlice(ChannelMappingMultithread(SliceToChannel([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}), toString, 10))
	if !reflect.DeepEqual(SetFromSlice(result), SetFromSlice([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"})) { // testing equality with sets since order of output is not preserved
		t.Errorf("ChannelMapping() returned unexpected %s", result)
	}

	if len(ChannelToSlice(ChannelMappingMultithread(SliceToChannel([]int{}), toString, 10))) != 0 {
		t.Error("ChannelMappingMultithread with a closed channel should return a closed channel")
	}
}
