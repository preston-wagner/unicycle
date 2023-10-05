package promises

import (
	"sync"

	"github.com/preston-wagner/unicycle/slices"
)

type Promissory[VALUE_TYPE any] struct {
	Value VALUE_TYPE
	Err   error
}

// a Promise represents data that is not yet available, but will be provided (most likely by a different goroutine) in the future
type Promise[VALUE_TYPE any] struct {
	awaiters []chan Promissory[VALUE_TYPE]
	result   *Promissory[VALUE_TYPE]
	lock     *sync.RWMutex
}

func NewPromise[VALUE_TYPE any]() *Promise[VALUE_TYPE] {
	return &Promise[VALUE_TYPE]{
		awaiters: []chan Promissory[VALUE_TYPE]{},
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
	c := make(chan Promissory[VALUE_TYPE])
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

func resolveChannel[VALUE_TYPE any](awaiter chan Promissory[VALUE_TYPE], prm Promissory[VALUE_TYPE]) {
	awaiter <- prm
}

func (promise *Promise[VALUE_TYPE]) Resolve(value VALUE_TYPE, err error) {
	prm := Promissory[VALUE_TYPE]{
		Value: value,
		Err:   err,
	}
	promise.lock.Lock()
	promise.result = &prm
	for _, awaiter := range promise.awaiters {
		go resolveChannel(awaiter, prm)
	}
	promise.awaiters = []chan Promissory[VALUE_TYPE]{} // empty the slice
	promise.lock.Unlock()
}

func AwaitAll[VALUE_TYPE any](promises ...*Promise[VALUE_TYPE]) []Promissory[VALUE_TYPE] {
	return slices.Mapping(promises, func(promise *Promise[VALUE_TYPE]) Promissory[VALUE_TYPE] {
		value, err := promise.Await()
		return Promissory[VALUE_TYPE]{
			Value: value,
			Err:   err,
		}
	})
}
