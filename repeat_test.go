package unicycle

import (
	"log"
	"testing"
	"time"
)

const loopTimes = 5
const duration = time.Second

func TestRepeat(t *testing.T) {
	callCount := 0
	increment := func() {
		callCount++
	}
	kill := Repeat(increment, time.Duration(float64(duration)/loopTimes))
	time.Sleep(duration)
	kill()
	if callCount != loopTimes {
		t.Errorf("Repeat was not called the expected number of times (%v != %v)", loopTimes, callCount)
	}
	time.Sleep(duration)
	kill() // extra kills to make sure multiple calls don't block or cause errors
	kill()
	kill()
	if callCount != loopTimes {
		t.Errorf("Repeat was not called the expected number of times (%v != %v)", loopTimes, callCount)
	}
}
