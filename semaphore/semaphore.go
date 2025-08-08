package semaphore

import (
	"sync"
)

// Semaphore is a type that simplifies safely reading and modifying single types across multiple threads
type Semaphore[T any] struct {
	value T
	lock  *sync.RWMutex
}

// Semaphore.Set returns the prior value to simplify operations that should happen exactly once
func (semaphore *Semaphore[T]) Set(value T) T {
	semaphore.lock.Lock()
	defer semaphore.lock.Unlock()
	prior := semaphore.value
	semaphore.value = value
	return prior
}

func (semaphore *Semaphore[T]) Get() T {
	semaphore.lock.RLock()
	defer semaphore.lock.RUnlock()
	return semaphore.value
}

func NewSemaphore[T any](value T) Semaphore[T] {
	return Semaphore[T]{
		value: value,
		lock:  &sync.RWMutex{},
	}
}
