package unicycle

import (
	"reflect"
	"testing"
)

func TestChannelFilter(t *testing.T) {
	result := ChannelToSlice(ChannelFilter(SliceToChannel([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}), odd))
	if !reflect.DeepEqual(result, []int{1, 3, 5, 7, 9}) {
		t.Errorf("ChannelFilter() returned unexpected %v", result)
	}
}

func TestChannelFilterMultithread(t *testing.T) {
	result := ChannelToSlice(ChannelFilterMultithread(SliceToChannel([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}), odd, 10))
	if !reflect.DeepEqual(SetFromSlice(result), SetFromSlice([]int{1, 3, 5, 7, 9})) { // testing equality with sets since order of output is not preserved
		t.Errorf("ChannelFilterMultithread() returned unexpected %v", result)
	}

	if len(ChannelToSlice(ChannelFilterMultithread(SliceToChannel([]int{}), odd, 10))) != 0 {
		t.Error("ChannelFilterMultithread with a closed channel should return a closed channel")
	}
}
