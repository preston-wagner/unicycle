package multithread

import (
	"sync"
)

type SemaphoreBool struct {
	value bool
	lock  *sync.RWMutex
}

// SemaphoreBool.Set returns the prior value to simplify operations that should happen exactly once
func (semaphore *SemaphoreBool) Set(value bool) bool {
	semaphore.lock.Lock()
	defer semaphore.lock.Unlock()
	prior := semaphore.value
	semaphore.value = value
	return prior
}

func (semaphore *SemaphoreBool) Get() bool {
	semaphore.lock.RLock()
	defer semaphore.lock.RUnlock()
	return semaphore.value
}

func NewSemaphoreBool() SemaphoreBool {
	return SemaphoreBool{
		value: false,
		lock:  &sync.RWMutex{},
	}
}
