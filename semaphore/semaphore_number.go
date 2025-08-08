package semaphore

import (
	"sync"

	"github.com/nuvi/unicycle/number"
)

type SemaphoreNumber[T number.Number] Semaphore[T]

func (semaphore *SemaphoreNumber[T]) Set(value T) T {
	semaphore.lock.Lock()
	defer semaphore.lock.Unlock()
	prior := semaphore.value
	semaphore.value = value
	return prior
}

func (semaphore *SemaphoreNumber[T]) Get() T {
	semaphore.lock.RLock()
	defer semaphore.lock.RUnlock()
	return semaphore.value
}

func (semaphore *SemaphoreNumber[T]) Add(value T) T {
	semaphore.lock.Lock()
	defer semaphore.lock.Unlock()
	semaphore.value += value
	return semaphore.value
}

func (semaphore *SemaphoreNumber[T]) Multiply(value T) T {
	semaphore.lock.Lock()
	defer semaphore.lock.Unlock()
	semaphore.value *= value
	return semaphore.value
}

func (semaphore *SemaphoreNumber[T]) Divide(value T) T {
	semaphore.lock.Lock()
	defer semaphore.lock.Unlock()
	semaphore.value /= value
	return semaphore.value
}

func NewSemaphoreNumber[T number.Number](value T) SemaphoreNumber[T] {
	return SemaphoreNumber[T]{
		value: value,
		lock:  &sync.RWMutex{},
	}
}
