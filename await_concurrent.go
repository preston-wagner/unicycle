package unicycle

import (
	"errors"
)

// AwaitConcurrent simplifies the common task of waiting until tasks on multiple threads have finished
func AwaitConcurrent(funcs ...func()) {
	pending := make(chan struct{})
	finished := 0
	for _, wrapped := range funcs {
		go awaitSafe(pending, wrapped)
	}
	for range pending {
		finished++
		if finished == len(funcs) {
			return
		}
	}
}

func awaitSafe(pending chan struct{}, wrapped func()) {
	wrapped()
	pending <- Empty
}

// Like AwaitConcurrent, but accepts functions that return errors, and returns the first error if there is one
func AwaitConcurrentWithErrors(funcs ...func() error) error {
	pending := make(chan error)
	finished := 0
	for _, wrapped := range funcs {
		go awaitUnsafe(pending, wrapped)
	}

	errs := []error{}
	for err := range pending {
		if err != nil {
			errs = append(errs, err)
		}
		finished++
		if finished == len(funcs) {
			break
		}
	}
	return errors.Join(errs...)
}

func awaitUnsafe(pending chan error, wrapped func() error) {
	pending <- wrapped()
}

var ErrAwaitConcurrentWithErrorsPanic = errors.New("panicking goroutine in AwaitConcurrentWithErrors recovered")
