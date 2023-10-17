package multithread

import (
	"testing"

	"github.com/nuvi/unicycle/channels"
	"github.com/nuvi/unicycle/test_ext"
)

func TestChannelEveryMultithread(t *testing.T) {
	if ChannelEveryMultithread(channels.SliceToChannel([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}), test_ext.Odd, 10) {
		t.Errorf("Every() returned false positive")
	}
	if !ChannelEveryMultithread(channels.SliceToChannel([]int{1, 3, 5, 7, 9}), test_ext.Odd, 10) {
		t.Errorf("Every() returned false negative")
	}
	if ChannelEveryMultithread(channels.SliceToChannel([]int{2, 4, 6, 8, 0}), test_ext.Odd, 10) {
		t.Errorf("Every() returned false positive")
	}
	if !ChannelEveryMultithread(channels.SliceToChannel([]int{}), test_ext.Odd, 10) {
		t.Errorf("Every() returned false negative")
	}
}
