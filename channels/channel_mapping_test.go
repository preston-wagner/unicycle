package channels

import (
	"errors"
	"reflect"
	"testing"

	"github.com/nuvi/unicycle/slices"
	"github.com/nuvi/unicycle/test_ext"
)

func TestChannelMapping(t *testing.T) {
	inputSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	result := ChannelToSlice(ChannelMapping(SliceToChannel(inputSlice), test_ext.ToString))
	if !reflect.DeepEqual(result, slices.Mapping(inputSlice, test_ext.ToString)) {
		t.Errorf("ChannelMapping() returned unexpected %s", result)
	}

	if len(ChannelToSlice(ChannelMapping(SliceToChannel([]int{}), test_ext.ToString))) != 0 {
		t.Error("ChannelMapping with a closed channel should return a closed channel")
	}
}

func TestChannelMappingWithError(t *testing.T) {
	valueFromResult := func(in ResultWithError[string]) string {
		return in.Value
	}
	errFromResult := func(in ResultWithError[string]) error {
		return in.Err
	}

	inputSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	result := slices.Mapping(ChannelToSlice(ChannelMappingWithError(SliceToChannel(inputSlice), test_ext.ToStringErrIfNegative)), valueFromResult)
	if !reflect.DeepEqual(result, slices.Mapping(inputSlice, test_ext.ToString)) {
		t.Errorf("ChannelMappingWithError() returned unexpected %s", result)
	}

	inputSlice = []int{1, 2, 3, 4, -5, 6, 7, 8, 9, 0}
	err := errors.Join(slices.Mapping(ChannelToSlice(ChannelMappingWithError(SliceToChannel(inputSlice), test_ext.ToStringErrIfNegative)), errFromResult)...)
	if err == nil {
		t.Errorf("ChannelMappingWithError() should have returned an error when a wrapped function did")
	}

	if len(ChannelToSlice(ChannelMappingWithError(SliceToChannel([]int{}), test_ext.ToStringErrIfNegative))) != 0 {
		t.Error("ChannelMappingWithError with a closed channel should return a closed channel")
	}
}
