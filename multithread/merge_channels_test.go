package multithread

import (
	"testing"
	"time"
)

func TestMergeChannels(t *testing.T) {
	inputs := []chan int{
		make(chan int),
		make(chan int),
		make(chan int),
	}

	go func() {
		inputs[0] <- 1
	}()
	go func() {
		inputs[1] <- 2
	}()
	go func() {
		inputs[2] <- 3
	}()

	output := MergeChannels(inputs)

	time.Sleep(time.Second)

	<-output
	<-output
	<-output
}
