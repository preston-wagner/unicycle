package channels

import (
	"reflect"
	"testing"
)

func TestSliceToChannel(t *testing.T) {
	input := []string{"a", "b", "c"}

	channel := SliceToChannel(input)

	if input[0] != <-channel {
		t.Error("SliceToChannel returned an incorrect value")
	}
	if input[1] != <-channel {
		t.Error("SliceToChannel returned an incorrect value")
	}
	if input[2] != <-channel {
		t.Error("SliceToChannel returned an incorrect value")
	}

	for range channel { // ensures channel was closed after all values were read

	}
}

func TestChannelToSlice(t *testing.T) {
	input := make(chan int)

	go func() {
		input <- 1
		input <- 2
		input <- 3
		close(input)
	}()

	slice := ChannelToSlice(input)

	if !reflect.DeepEqual(slice, []int{1, 2, 3}) {
		t.Error("ChannelToSlice did not return the correct slice of data")
	}
}
