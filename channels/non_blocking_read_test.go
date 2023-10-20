package channels

import "testing"

func TestNonBlockingRead(t *testing.T) {
	input := make(chan string, 1)

	_, ok := NonBlockingRead(input)
	if ok {
		t.Error("NonBlockingRead should return ok=false when no items are available in the channel")
	}

	input <- "lorem"

	value, ok := NonBlockingRead(input)
	if !ok {
		t.Error("NonBlockingRead should return ok=true when an item is available in the channel")
	}
	if value != "lorem" {
		t.Error("NonBlockingRead returned the wrong value")
	}

	_, ok = NonBlockingRead(input)
	if ok {
		t.Error("NonBlockingRead should return ok=false when no items are available in the channel")
	}
}
