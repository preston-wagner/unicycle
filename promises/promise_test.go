package promises

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/preston-wagner/unicycle/slices_ext"
)

const loopTimes = 4
const duration = time.Second / 4

func TestPromise(t *testing.T) {
	successes := 0

	prm := NewPromise[string]()

	for i := 0; i < loopTimes; i++ {
		go func() {
			_, err := prm.Await()
			if err != nil {
				t.Errorf("no error should be returned when resolved without one")
			}
			successes += 1
		}()
	}

	value := "hello world"
	prm.Resolve(value, nil)

	time.Sleep(duration)

	result, err := prm.Await()
	if err != nil {
		t.Error(err)
	}
	if result != value {
		t.Errorf("Promise.Await() returned wrong result; expected %v, got %v", value, result)
	}

	if successes != loopTimes {
		t.Errorf("Not all goroutines received resolution as expected (%v != %v)", loopTimes, successes)
	}

	value = "new value"

	prm.Resolve(value, nil)

	result, err = prm.Await()

	if err != nil {
		t.Error(err)
	}

	if result != value {
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

	prm = WrapInPromise(func() (string, error) {
		time.Sleep(duration)
		return "", errors.New("bad time")
	})

	result, err = prm.Await()
	if err == nil {
		t.Errorf("When the wrapped function returns an error, Promise.Resolve() should as well")
	}

	if result != "" {
		t.Errorf("wrong result returned")
	}
}

func TestAwaitAll(t *testing.T) {
	result := slices_ext.Mapping(AwaitAll(
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
