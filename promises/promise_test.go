package promises

import (
	"reflect"
	"testing"
	"time"

	"github.com/nuvi/unicycle/slices"
)

const loopTimes = 4
const duration = time.Second / 4

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

func TestAwaitAll(t *testing.T) {
	result := slices.Mapping(AwaitAll(
		WrapInPromise(func() (int, error) { time.Sleep(duration * 3); return 1, nil }),
		WrapInPromise(func() (int, error) { time.Sleep(duration * 2); return 2, nil }),
		WrapInPromise(func() (int, error) { time.Sleep(duration * 1); return 3, nil }),
	), func(prm Promissory[int]) int {
		return prm.Value
	})
	if !reflect.DeepEqual(result, []int{1, 2, 3}) {
		t.Errorf("AwaitAll() returned unexpected %v", result)
	}

	if len(AwaitAll[int]()) != 0 {
		t.Error("AwaitAll() with no args should return a slice with length 0")
	}
}
