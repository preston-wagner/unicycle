package multithread

import (
	"github.com/nuvi/unicycle/defaults"
)

// like slices_ext.MappingFind(), but all mutating/filter functions run in parallel in their own goroutines
// WARNING: unlike Find(), the returned result is not guaranteed to be the first in array order; only the first whose mutator returns
func MappingFindMultithread[INPUT_TYPE any, OUTPUT_TYPE any](input []INPUT_TYPE, mutatingFilter func(INPUT_TYPE) (OUTPUT_TYPE, bool)) (OUTPUT_TYPE, bool) {
	total := len(input)
	if total == 0 {
		return defaults.ZeroValue[OUTPUT_TYPE](), false
	}
	counter := NewSemaphoreInt()
	done := NewSemaphoreBool()
	success := make(chan OUTPUT_TYPE)
	failure := make(chan struct{})
	for _, value := range input {
		go func(value INPUT_TYPE) {
			if !done.Get() {
				result, ok := mutatingFilter(value)
				if ok && !done.Set(true) { // .Set() returns the prior value, so this ensures we only send on the success channel once
					success <- result
				} else {
					if counter.Add(1) == total {
						failure <- defaults.Empty
					}
				}
			}
		}(value)
	}
	select {
	case result := <-success:
		return result, true
	case <-failure:
		return defaults.ZeroValue[OUTPUT_TYPE](), false
	}
}
