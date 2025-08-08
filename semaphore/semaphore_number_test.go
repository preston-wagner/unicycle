package semaphore

import (
	"testing"
)

func TestSemaphoreNumber(t *testing.T) {
	sema := NewSemaphoreNumber(0)

	sema.Add(5)

	if result := sema.Get(); result != 5 {
		t.Errorf("sema.Get() returned wrong result %v", result)
	}

	sema.Multiply(3)

	if result := sema.Get(); result != 15 {
		t.Errorf("sema.Get() returned wrong result %v", result)
	}

	sema.Divide(5)

	if result := sema.Get(); result != 3 {
		t.Errorf("sema.Get() returned wrong result %v", result)
	}
}
