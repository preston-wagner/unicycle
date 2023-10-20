package channels

import (
	"log"
	"reflect"
	"testing"
	"time"
)

func TestChannelBatch(t *testing.T) {
	input := make(chan string)
	output := ChannelBatch(input, 5)
	input <- "lorem"
	input <- "ipsum"
	input <- "dolor"
	input <- "sit"
	input <- "amet"
	batch := <-output
	if !reflect.DeepEqual(batch, []string{"lorem", "ipsum", "dolor", "sit", "amet"}) {
		t.Error("ChannelBatch should have bundled all the provided entries into a single slice")
	}
	input <- "singleton"
	batch = <-output
	if !reflect.DeepEqual(batch, []string{"singleton"}) {
		t.Error("ChannelBatch should have returned a single item slice when only a single item was available")
	}
	input <- "final"
	close(input)
	time.Sleep(time.Second)
	batch = <-output
	if !reflect.DeepEqual(batch, []string{"final"}) {
		t.Error("ChannelBatch should have closed out the batch after the input channel was closed")
	}
	for range output { // ensures closing input closes output
	}

	input = make(chan string)
	output = ChannelBatch(input, 5)
	close(input)
	for range output { // ensures closing input closes output (even when nothing has been sent on the input channel)
	}

	input = make(chan string)
	close(input)
	output = ChannelBatch(input, 5)
	for range output { // ensures closing input closes output (even when input channel was closed before providing it)
	}
}

func TestChannelDebatch(t *testing.T) {
	input := make(chan []string)
	output := ChannelDebatch(input)
	go func() {
		input <- []string{
			"lorem",
			"ipsum",
			"dolor",
			"sit",
			"amet",
		}
	}()
	if <-output != "lorem" {
		t.Error("ChannelDebatch should have split the provided batch into constituent items, preserving order")
	}
	if <-output != "ipsum" {
		t.Error("ChannelDebatch should have split the provided batch into constituent items, preserving order")
	}
	if <-output != "dolor" {
		t.Error("ChannelDebatch should have split the provided batch into constituent items, preserving order")
	}
	if <-output != "sit" {
		t.Error("ChannelDebatch should have split the provided batch into constituent items, preserving order")
	}
	if <-output != "amet" {
		t.Error("ChannelDebatch should have split the provided batch into constituent items, preserving order")
	}
	log.Println("4")
	go func() {
		input <- []string{"singleton"}
	}()
	if <-output != "singleton" {
		t.Error("ChannelDebatch should have debatched the batch regardless of size")
	}
	log.Println("5")
	close(input)
	for range output { // ensures closing input closes output
	}
	log.Println("6")
}
