package promises

import (
	"testing"
	"time"
)

func TestEmptySafePromise(t *testing.T) {
	successes := 0

	prm := NewEmptySafePromise()

	logResult := func() {
		prm.Await()
		successes += 1
	}

	for i := 0; i < loopTimes; i++ {
		go logResult()
	}

	prm.Resolve()

	time.Sleep(duration)

	if successes != loopTimes {
		t.Errorf("Not all goroutines received resolution as expected (%v != %v)", loopTimes, successes)
	}

	prm.Resolve()

	prm.Await()
}

func TestWrapInEmptySafePromise(t *testing.T) {
	successes := 0

	prm := WrapInEmptySafePromise(func() {
		time.Sleep(duration)
		successes += 1
	})

	prm.Await()
	prm.Await()
	prm.Await()

	if successes != 1 {
		t.Errorf("Not all goroutines received resolution as expected (%v != %v)", 1, successes)
	}
}

func TestAwaitAllEmptySafe(t *testing.T) {
	successes := 0

	AwaitAllEmptySafe(
		WrapInEmptySafePromise(func() { time.Sleep(duration * 3); successes += 1 }),
		WrapInEmptySafePromise(func() { time.Sleep(duration * 2); successes += 1 }),
		WrapInEmptySafePromise(func() { time.Sleep(duration * 1); successes += 1 }),
	)

	if successes != 3 {
		t.Errorf("Not all goroutines received resolution as expected (%v != %v)", 3, successes)
	}
}
