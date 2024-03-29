package promises

import (
	"reflect"
	"testing"
	"time"
)

func TestSafePromise(t *testing.T) {
	successes := 0

	prm := NewSafePromise[string]()

	for i := 0; i < loopTimes; i++ {
		go func() {
			prm.Await()
			successes += 1
		}()
	}

	value := "hello world"

	prm.Resolve(value)

	time.Sleep(duration)

	result := prm.Await()
	if result != value {
		t.Errorf("SafePromise.Await() returned wrong result; expected %v, got %v", value, result)
	}

	if successes != loopTimes {
		t.Errorf("Not all goroutines received resolution as expected (%v != %v)", loopTimes, successes)
	}

	value = "new value"

	prm.Resolve(value)

	result = prm.Await()

	if result != value {
		t.Errorf("re-resolution with new value did not work as expected")
	}
}

func TestWrapInSafePromise(t *testing.T) {
	prm := WrapInSafePromise(func() string {
		time.Sleep(duration)
		return "done"
	})

	result := prm.Await()

	if result != "done" {
		t.Errorf("wrong result returned")
	}
}

func TestAwaitAllSafe(t *testing.T) {
	result := AwaitAllSafe(
		WrapInSafePromise(func() int { time.Sleep(duration * 3); return 1 }),
		WrapInSafePromise(func() int { time.Sleep(duration * 2); return 2 }),
		WrapInSafePromise(func() int { time.Sleep(duration * 1); return 3 }),
	)
	if !reflect.DeepEqual(result, []int{1, 2, 3}) {
		t.Errorf("AwaitAllSafe() returned unexpected %v", result)
	}

	if len(AwaitAll[int]()) != 0 {
		t.Error("AwaitAllSafe() with no args should return a slice with length 0")
	}
}
