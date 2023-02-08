package unicycle

import (
	"testing"
	"time"
)

func TestPromise(t *testing.T) {
	successes := 0

	prm := NewPromise[string]()

	logResult := func() {
		_, err := prm.Await()
		if err != nil {
			t.Errorf("no error should be returned when resolved without one")
		}
		successes += 1
	}

	for i := 0; i < loopTimes; i++ {
		go logResult()
	}

	prm.Resolve("hello world", nil)

	time.Sleep(duration)

	if successes != loopTimes {
		t.Errorf("Not all goroutines received resolution as expected (%v != %v)", loopTimes, successes)
	}

	newValue := "new value"

	prm.Resolve(newValue, nil)

	result, _ := prm.Await()

	if result != newValue {
		t.Errorf("re-resolution with new value did not work as expected")
	}
}

func TestWrapInPromise(t *testing.T) {
	prm := WrapInPromise(func() (string, error) {
		time.Sleep(duration)
		return "done", nil
	})

	result, err := prm.Await()
	if err != nil {
		t.Errorf("no error should be returned when resolved without one")
	}

	if result != "done" {
		t.Errorf("wrong result returned")
	}
}
