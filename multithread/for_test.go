package multithread

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/nuvi/unicycle/sets"
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

func BenchmarkForMultithread10(b *testing.B)      { benchmarkForMultithread(b, 10) }
func BenchmarkForMultithread100(b *testing.B)     { benchmarkForMultithread(b, 100) }
func BenchmarkForMultithread1000(b *testing.B)    { benchmarkForMultithread(b, 1000) }
func BenchmarkForMultithread10000(b *testing.B)   { benchmarkForMultithread(b, 10000) }
func BenchmarkForMultithread100000(b *testing.B)  { benchmarkForMultithread(b, 100000) }
func BenchmarkForMultithread1000000(b *testing.B) { benchmarkForMultithread(b, 1000000) }

func benchmarkForMultithread(b *testing.B, size int) {
	inputs := make([]int, 0, size)
	for i := 0; i < size; i++ {
		inputs = append(inputs, rand.Int())
	}
	for i := 0; i < b.N; i++ {
		ForMultithread(inputs, func(value int) {
			// not doing anything except benchmarking the wrapper
		})
	}
}
