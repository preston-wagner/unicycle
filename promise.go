package unicycle

import (
	"sync"
)

type promissory[VALUE_TYPE any] struct {
	Value VALUE_TYPE
	Err   error
}

// a Promise represents data that is not yet available, but will be provided (most likely by a different goroutine) in the future
type Promise[VALUE_TYPE any] struct {
	awaiters []chan promissory[VALUE_TYPE]
	result   *promissory[VALUE_TYPE]
	lock     *sync.RWMutex
}

func NewPromise[VALUE_TYPE any]() *Promise[VALUE_TYPE] {
	return &Promise[VALUE_TYPE]{
		awaiters: []chan promissory[VALUE_TYPE]{},
		lock:     &sync.RWMutex{},
	}
}

func WrapInPromise[VALUE_TYPE any](wrapped func() (VALUE_TYPE, error)) *Promise[VALUE_TYPE] {
	promise := NewPromise[VALUE_TYPE]()
	go func() {
		promise.Resolve(wrapped())
	}()
	return promise
}

func (promise *Promise[VALUE_TYPE]) Await() (VALUE_TYPE, error) {
	c := make(chan promissory[VALUE_TYPE])
	promise.lock.Lock()
	if promise.result != nil {
		defer promise.lock.Unlock()
		return promise.result.Value, promise.result.Err
	}
	promise.awaiters = append(promise.awaiters, c)
	promise.lock.Unlock()
	result := <-c
	return result.Value, result.Err
}

func resolveChannel[VALUE_TYPE any](awaiter chan promissory[VALUE_TYPE], prm promissory[VALUE_TYPE]) {
	awaiter <- prm
}

func (promise *Promise[VALUE_TYPE]) Resolve(value VALUE_TYPE, err error) {
	prm := promissory[VALUE_TYPE]{
		Value: value,
		Err:   err,
	}
	promise.lock.Lock()
	promise.result = &prm
	for _, awaiter := range promise.awaiters {
		go resolveChannel(awaiter, prm)
	}
	promise.awaiters = []chan promissory[VALUE_TYPE]{} // empty the slice
	promise.lock.Unlock()
}
