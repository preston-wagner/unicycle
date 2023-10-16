package promises

import (
	"errors"
	"testing"
	"time"
)

func TestEmptyPromise(t *testing.T) {
	successes := 0

	prm := NewEmptyPromise()

	logResult := func() {
		err := prm.Await()
		if err != nil {
			t.Errorf("no error should be returned when resolved without one")
		}
		successes += 1
	}

	for i := 0; i < loopTimes; i++ {
		go logResult()
	}

	prm.Resolve(nil)

	time.Sleep(duration)

	if successes != loopTimes {
		t.Errorf("Not all goroutines received resolution as expected (%v != %v)", loopTimes, successes)
	}

	prm.Resolve(nil)

	err := prm.Await()

	if err != nil {
		t.Error(err)
	}
}

func TestWrapInEmptyPromise(t *testing.T) {
	prm := WrapInEmptyPromise(func() error {
		time.Sleep(duration)
		return nil
	})

	err := prm.Await()
	if err != nil {
		t.Errorf("no error should be returned when resolved without one")
	}
}

func TestAwaitAllEmpty(t *testing.T) {
	err := AwaitAllEmpty(
		WrapInEmptyPromise(func() error { return nil }),
		WrapInEmptyPromise(func() error { return nil }),
		WrapInEmptyPromise(func() error { return nil }),
	)
	if err != nil {
		t.Errorf("AwaitAllEmpty() returned unexpected error %v", err)
	}

	err = AwaitAllEmpty(
		WrapInEmptyPromise(func() error { return nil }),
		WrapInEmptyPromise(func() error { return errors.New("bad result 1") }),
		WrapInEmptyPromise(func() error { return nil }),
	)
	if err == nil {
		t.Errorf("AwaitAllEmpty() should have returned an error when a resolver did")
	}

	if AwaitAllEmpty() != nil {
		t.Error("AwaitAllEmpty() with no args should return nil")
	}
}
