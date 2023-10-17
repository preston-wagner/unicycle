package channels

import (
	"reflect"
	"testing"

	"github.com/nuvi/unicycle/test_ext"
)

func TestChannelFilter(t *testing.T) {
	result := ChannelToSlice(ChannelFilter(SliceToChannel([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}), test_ext.Odd))
	if !reflect.DeepEqual(result, []int{1, 3, 5, 7, 9}) {
		t.Errorf("ChannelFilter() returned unexpected %v", result)
	}
}
