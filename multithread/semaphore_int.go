package multithread

import (
	"sync"
)

type SemaphoreInt struct {
	value int
	lock  *sync.RWMutex
}

func (semaphore *SemaphoreInt) Add(value int) int {
	semaphore.lock.Lock()
	defer semaphore.lock.Unlock()
	semaphore.value += value
	return semaphore.value
}

func (semaphore *SemaphoreInt) Set(value int) {
	semaphore.lock.Lock()
	defer semaphore.lock.Unlock()
	semaphore.value = value
}

func (semaphore *SemaphoreInt) Get() int {
	semaphore.lock.RLock()
	defer semaphore.lock.RUnlock()
	return semaphore.value
}

func NewSemaphoreInt() SemaphoreInt {
	return SemaphoreInt{
		value: 0,
		lock:  &sync.RWMutex{},
	}
}
