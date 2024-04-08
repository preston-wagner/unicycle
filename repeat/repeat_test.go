package repeat

import (
	"testing"
	"time"
)

const loopTimes = 4
const duration = time.Second / 4

func TestRepeat(t *testing.T) {
	callCount := 0
	increment := func() {
		callCount++
	}
	kill := Repeat(increment, time.Duration(float64(duration)/loopTimes), false)
	time.Sleep(duration)
	kill()
	time.Sleep(duration)
	if callCount != loopTimes {
		t.Errorf("Repeat handler was not called the expected number of times (%v != %v)", loopTimes, callCount)
	}
	time.Sleep(duration)
	kill() // extra kills to make sure multiple calls don't block or cause errors
	kill()
	kill()
	if callCount != loopTimes {
		t.Errorf("Repeat handler was not called the expected number of times (%v != %v)", loopTimes, callCount)
	}
}

func TestRepeatBefore(t *testing.T) {
	callCount := 0
	increment := func() {
		callCount++
	}
	kill := Repeat(increment, time.Duration(float64(duration)/loopTimes), true)
	time.Sleep(duration)
	kill()
	time.Sleep(duration)
	if callCount != loopTimes+1 {
		t.Errorf("Repeat handler was not called the expected number of times (%v != %v)", loopTimes+1, callCount)
	}
	time.Sleep(duration)
	kill() // extra kills to make sure multiple calls don't block or cause errors
	kill()
	kill()
	if callCount != loopTimes+1 {
		t.Errorf("Repeat handler was not called the expected number of times (%v != %v)", loopTimes+1, callCount)
	}
}

func TestRepeatMultithread(t *testing.T) {
	callCount := 0
	increment := func() {
		callCount++
	}
	kill := RepeatMultithread(increment, time.Duration(float64(duration)/loopTimes), false)
	time.Sleep(duration)
	kill()
	time.Sleep(duration)
	if callCount != loopTimes {
		t.Errorf("RepeatMultithread handler was not called the expected number of times (%v != %v)", loopTimes, callCount)
	}
	time.Sleep(duration)
	kill() // extra kills to make sure multiple calls don't block or cause errors
	kill()
	kill()
	if callCount != loopTimes {
		t.Errorf("RepeatMultithread handler was not called the expected number of times (%v != %v)", loopTimes, callCount)
	}
}
