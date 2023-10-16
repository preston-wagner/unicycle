package channels

import (
	"reflect"
	"sync"
	"testing"

	"github.com/preston-wagner/unicycle/sets"
)

type accumulator struct {
	set  sets.Set[string]
	lock *sync.RWMutex
}

func (acc *accumulator) Add(value string) {
	acc.lock.Lock()
	defer acc.lock.Unlock()
	acc.set.Add(value)
}

func newAccumulator() accumulator {
	return accumulator{
		set:  sets.Set[string]{},
		lock: &sync.RWMutex{},
	}
}

func TestChannelForEach(t *testing.T) {
	input := []string{"a", "b", "c", "d", "e", "f", "g"}

	acc := newAccumulator()

	ChannelForEach(SliceToChannel(input), acc.Add)

	if !reflect.DeepEqual(acc.set, sets.SetFromSlice(input)) {
		t.Error("ChannelForEach() worker wasn't called on every value")
	}
}
