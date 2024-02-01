package multithread

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/preston-wagner/unicycle/sets"
)

func TestForMultithread(t *testing.T) {
	inputs := []int{}
	const size = 10000
	for i := 0; i < size; i++ {
		inputs = append(inputs, i)
	}

	set := sets.SetFromSlice([]int{})
	lock := &sync.Mutex{}

	ForMultithread(inputs, func(value int) {
		fraction := time.Duration(rand.Int())
		if fraction != 0 {
			time.Sleep(time.Second / fraction)
		}
		lock.Lock()
		defer lock.Unlock()
		set.Add(value)
	})

	for _, value := range inputs {
		if !set.Has(value) {
			t.Error("ForMultithread missed", value)
		}
	}

	ForMultithread(nil, func(value int) {}) // make sure calling with empty array doesn't block
}
