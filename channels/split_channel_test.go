package channels

import (
	"testing"
)

func TestSplitChannel(t *testing.T) {
	input := make(chan int)

	go func() {
		input <- 1
	}()
	go func() {
		input <- 2
	}()
	go func() {
		input <- 3
	}()

	outputs := SplitChannel(input, 3)

	<-outputs[0]
	<-outputs[1]
	<-outputs[2]
}

func TestMergeChannels(t *testing.T) {
	inputs := []chan int{
		make(chan int),
		make(chan int),
		make(chan int),
	}

	go func() {
		inputs[0] <- 1
		inputs[1] <- 2
		inputs[2] <- 3
	}()

	output := MergeChannels(inputs, 0)

	<-output
	<-output
	<-output
}
