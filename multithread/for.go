package multithread

import (
	"github.com/nuvi/unicycle/defaults"
)

// like a for loop, but each loop runs in its own goroutine simultaneously, returning once all of them are finished
func ForMultithread[INPUT_TYPE any](input []INPUT_TYPE, apply func(INPUT_TYPE)) {
	total := len(input)
	if total == 0 {
		return
	}
	counter := NewSemaphoreInt()
	done := make(chan struct{})
	for index := range input {
		go func(index int) {
			apply(input[index])
			finished := counter.Add(1)
			if finished == total {
				done <- defaults.Empty
			}
		}(index)
	}
	<-done
}
