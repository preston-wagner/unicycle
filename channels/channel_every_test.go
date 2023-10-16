package channels

import (
	"testing"

	"github.com/preston-wagner/unicycle/test_ext"
)

func TestChannelEvery(t *testing.T) {
	if ChannelEvery(SliceToChannel([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}), test_ext.Odd) {
		t.Errorf("Every() returned false positive")
	}
	if !ChannelEvery(SliceToChannel([]int{1, 3, 5, 7, 9}), test_ext.Odd) {
		t.Errorf("Every() returned false negative")
	}
	if ChannelEvery(SliceToChannel([]int{2, 4, 6, 8, 0}), test_ext.Odd) {
		t.Errorf("Every() returned false positive")
	}
	if !ChannelEvery(SliceToChannel([]int{}), test_ext.Odd) {
		t.Errorf("Every() returned false negative")
	}
}
