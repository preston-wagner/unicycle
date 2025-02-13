package promises

import (
	"sync"

	"github.com/preston-wagner/unicycle/defaults"
)

// an EmptySafePromise represents a function call that returns no data or errors, but we're waiting on the resolution of anyways
type EmptySafePromise struct {
	awaiters []chan struct{}
	resolved bool
	lock     *sync.RWMutex
}

func NewEmptySafePromise() *EmptySafePromise {
	return &EmptySafePromise{
		awaiters: []chan struct{}{},
		lock:     &sync.RWMutex{},
	}
}

func WrapInEmptySafePromise(wrapped func()) *EmptySafePromise {
	promise := NewEmptySafePromise()
	go func() {
		wrapped()
		promise.Resolve()
	}()
	return promise
}

func (promise *EmptySafePromise) Await() {
	promise.lock.Lock()
	if promise.resolved {
		defer promise.lock.Unlock()
		return
	}
	c := make(chan struct{})
	promise.awaiters = append(promise.awaiters, c)
	promise.lock.Unlock()
	<-c
}

func resolveChannelEmptySafe(awaiter chan struct{}) {
	awaiter <- defaults.Empty
}

func (promise *EmptySafePromise) Resolve() {
	promise.lock.Lock()
	promise.resolved = true
	for _, awaiter := range promise.awaiters {
		go resolveChannelEmptySafe(awaiter)
	}
	promise.awaiters = []chan struct{}{} // empty the slice
	promise.lock.Unlock()
}

func AwaitAllEmptySafe(promises ...*EmptySafePromise) {
	for _, prm := range promises {
		prm.Await()
	}
}
