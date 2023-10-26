package multithread

import (
	"reflect"
	"sync"
	"testing"

	"github.com/nuvi/unicycle/sets"
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

func TestForEachMultithread(t *testing.T) {
	input := []string{"a", "b", "c", "d", "e", "f", "g"}

	acc := newAccumulator()

	ForEachMultithread(input, acc.Add, len(input)/3)

	if !reflect.DeepEqual(acc.set, sets.SetFromSlice(input)) {
		t.Error("ForEachMultithread() wasn't called on every value")
	}
}

func TestForEachMultithreadSingle(t *testing.T) {
	input := []string{"a", "b", "c", "d", "e", "f", "g"}

	acc := newAccumulator()

	ForEachMultithread(input, acc.Add, 1)

	if !reflect.DeepEqual(acc.set, sets.SetFromSlice(input)) {
		t.Error("ForEachMultithread() wasn't called on every value")
	}
}

func TestForEachMultithreadManyWorkers(t *testing.T) {
	input := []string{"a", "b", "c", "d", "e", "f", "g"}

	acc := newAccumulator()

	ForEachMultithread(input, acc.Add, len(input)+10)

	if !reflect.DeepEqual(acc.set, sets.SetFromSlice(input)) {
		t.Error("ForEachMultithread() wasn't called on every value")
	}
}
