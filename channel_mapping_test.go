package unicycle

import (
	"reflect"
	"testing"
)

func sliceToChannel[INPUT_TYPE any](input []INPUT_TYPE) chan INPUT_TYPE {
	output := make(chan INPUT_TYPE)
	go func() {
		for _, value := range input {
			output <- value
		}
		close(output)
	}()
	return output
}

func channelToSlice[INPUT_TYPE any](input chan INPUT_TYPE) []INPUT_TYPE {
	output := make([]INPUT_TYPE, 0)
	for value := range input {
		output = append(output, value)
	}
	return Trim(output)
}

func TestChannelMapping(t *testing.T) {
	result := channelToSlice(ChannelMapping(sliceToChannel([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}), toString))
	if !reflect.DeepEqual(result, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}) {
		t.Errorf("ChannelMapping() returned unexpected %s", result)
	}
}
