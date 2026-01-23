package multithread

import (
	"errors"

	"github.com/preston-wagner/unicycle/defaults"
	"github.com/preston-wagner/unicycle/semaphore"
)

// like a for loop, but each loop runs in its own goroutine simultaneously, returning once all of them are finished
func ForMultithread[INPUT_TYPE any](input []INPUT_TYPE, apply func(INPUT_TYPE)) {
	total := len(input)
	if total == 0 {
		return
	}
	counter := semaphore.NewSemaphoreNumber(0)
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

// like For, but apply function can return an error
func ForMultithreadWithError[INPUT_TYPE any](input []INPUT_TYPE, apply func(INPUT_TYPE) error) error {
	total := len(input)
	if total == 0 {
		return nil
	}
	counter := semaphore.NewSemaphoreNumber(0)
	done := make(chan struct{})
	errs := make([]error, len(input))
	for index := range input {
		go func(index int) {
			errs[index] = apply(input[index])
			finished := counter.Add(1)
			if finished == total {
				done <- defaults.Empty
			}
		}(index)
	}
	<-done
	return errors.Join(errs...)
}
