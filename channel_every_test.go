package unicycle

import (
	"testing"
)

func TestChannelEvery(t *testing.T) {
	if ChannelEvery(sliceToChannel([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}), odd) {
		t.Errorf("Every() returned false positive")
	}
	if !ChannelEvery(sliceToChannel([]int{1, 3, 5, 7, 9}), odd) {
		t.Errorf("Every() returned false negative")
	}
	if ChannelEvery(sliceToChannel([]int{2, 4, 6, 8, 0}), odd) {
		t.Errorf("Every() returned false positive")
	}
	if !ChannelEvery(sliceToChannel([]int{}), odd) {
		t.Errorf("Every() returned false negative")
	}
}

func TestChannelEveryMultithread(t *testing.T) {
	if ChannelEveryMultithread(sliceToChannel([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}), odd, 10) {
		t.Errorf("Every() returned false positive")
	}
	if !ChannelEveryMultithread(sliceToChannel([]int{1, 3, 5, 7, 9}), odd, 10) {
		t.Errorf("Every() returned false negative")
	}
	if ChannelEveryMultithread(sliceToChannel([]int{2, 4, 6, 8, 0}), odd, 10) {
		t.Errorf("Every() returned false positive")
	}
	if !ChannelEveryMultithread(sliceToChannel([]int{}), odd, 10) {
		t.Errorf("Every() returned false negative")
	}
}
