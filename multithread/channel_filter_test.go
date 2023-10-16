package multithread

import (
	"reflect"
	"testing"

	"github.com/preston-wagner/unicycle/channels"
	"github.com/preston-wagner/unicycle/sets"
	"github.com/preston-wagner/unicycle/test_ext"
)

func TestChannelFilterMultithread(t *testing.T) {
	result := channels.ChannelToSlice(ChannelFilterMultithread(channels.SliceToChannel([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}), test_ext.Odd, 10))
	if !reflect.DeepEqual(sets.SetFromSlice(result), sets.SetFromSlice([]int{1, 3, 5, 7, 9})) { // testing equality with sets since order of output is not preserved
		t.Errorf("ChannelFilterMultithread() returned unexpected %v", result)
	}

	if len(channels.ChannelToSlice(ChannelFilterMultithread(channels.SliceToChannel([]int{}), test_ext.Odd, 10))) != 0 {
		t.Error("ChannelFilterMultithread with a closed channel should return a closed channel")
	}
}
