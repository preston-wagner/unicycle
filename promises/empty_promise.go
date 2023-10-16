package promises

import (
	"errors"
	"sync"

	"github.com/preston-wagner/unicycle/slices"
)

// an EmptyPromise represents a function (that may return an error) that has not yet resolved but will in the future
type EmptyPromise struct {
	awaiters []chan error
	result   *error
	lock     *sync.RWMutex
}

func NewEmptyPromise() *EmptyPromise {
	return &EmptyPromise{
		awaiters: []chan error{},
		lock:     &sync.RWMutex{},
	}
}

func WrapInEmptyPromise(wrapped func() error) *EmptyPromise {
	promise := NewEmptyPromise()
	go func() {
		promise.Resolve(wrapped())
	}()
	return promise
}

func (promise *EmptyPromise) Await() error {
	promise.lock.Lock()
	if promise.result != nil {
		defer promise.lock.Unlock()
		return *promise.result
	}
	c := make(chan error)
	promise.awaiters = append(promise.awaiters, c)
	promise.lock.Unlock()
	return <-c
}

func resolveChannelEmpty(awaiter chan error, err error) {
	awaiter <- err
}

func (promise *EmptyPromise) Resolve(err error) {
	promise.lock.Lock()
	promise.result = &err
	for _, awaiter := range promise.awaiters {
		go resolveChannelEmpty(awaiter, err)
	}
	promise.awaiters = []chan error{} // empty the slice
	promise.lock.Unlock()
}

func AwaitAllEmpty(promises ...*EmptyPromise) error {
	return errors.Join(slices.Mapping(promises, func(promise *EmptyPromise) error {
		return promise.Await()
	})...)
}
