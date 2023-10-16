package promises

import (
	"sync"

	"github.com/preston-wagner/unicycle/slices"
)

// a SafePromise represents data that is not yet available, but will be provided (most likely by a different goroutine) in the future
type SafePromise[VALUE_TYPE any] struct {
	awaiters []chan VALUE_TYPE
	result   *VALUE_TYPE
	lock     *sync.RWMutex
}

func NewSafePromise[VALUE_TYPE any]() *SafePromise[VALUE_TYPE] {
	return &SafePromise[VALUE_TYPE]{
		awaiters: []chan VALUE_TYPE{},
		lock:     &sync.RWMutex{},
	}
}

func WrapInSafePromise[VALUE_TYPE any](wrapped func() VALUE_TYPE) *SafePromise[VALUE_TYPE] {
	promise := NewSafePromise[VALUE_TYPE]()
	go func() {
		promise.Resolve(wrapped())
	}()
	return promise
}

func (promise *SafePromise[VALUE_TYPE]) Await() VALUE_TYPE {
	promise.lock.Lock()
	if promise.result != nil {
		defer promise.lock.Unlock()
		return *promise.result
	}
	c := make(chan VALUE_TYPE)
	promise.awaiters = append(promise.awaiters, c)
	promise.lock.Unlock()
	return <-c
}

func resolveChannelSafe[VALUE_TYPE any](awaiter chan VALUE_TYPE, prm VALUE_TYPE) {
	awaiter <- prm
}

func (promise *SafePromise[VALUE_TYPE]) Resolve(value VALUE_TYPE) {
	promise.lock.Lock()
	promise.result = &value
	for _, awaiter := range promise.awaiters {
		go resolveChannelSafe(awaiter, value)
	}
	promise.awaiters = []chan VALUE_TYPE{} // empty the slice
	promise.lock.Unlock()
}

func AwaitAllSafe[VALUE_TYPE any](promises ...*SafePromise[VALUE_TYPE]) []VALUE_TYPE {
	return slices.Mapping(promises, func(prm *SafePromise[VALUE_TYPE]) VALUE_TYPE {
		return prm.Await()
	})
}
