package multithread

import (
	"github.com/preston-wagner/unicycle/promises"
	"github.com/preston-wagner/unicycle/slices"
)

// like Mapping(), but all mutator functions run in parallel in their own goroutines
func MappingMultithread[INPUT_TYPE any, OUTPUT_TYPE any](input []INPUT_TYPE, mutator func(INPUT_TYPE) OUTPUT_TYPE) []OUTPUT_TYPE {
	return promises.AwaitAllSafe(
		slices.Mapping(input, func(value INPUT_TYPE) *promises.SafePromise[OUTPUT_TYPE] {
			return promises.WrapInSafePromise(func() OUTPUT_TYPE {
				return mutator(value)
			})
		})...,
	)
}

// like MappingWithError(), but all mutator functions run in parallel in their own goroutines
func MappingMultithreadWithError[INPUT_TYPE any, OUTPUT_TYPE any](input []INPUT_TYPE, mutator func(INPUT_TYPE) (OUTPUT_TYPE, error)) ([]OUTPUT_TYPE, error) {
	pending := slices.Mapping(input, func(value INPUT_TYPE) *promises.Promise[OUTPUT_TYPE] {
		return promises.WrapInPromise(func() (OUTPUT_TYPE, error) {
			return mutator(value)
		})
	})
	return slices.MappingWithError(pending, func(prm *promises.Promise[OUTPUT_TYPE]) (OUTPUT_TYPE, error) {
		return prm.Await()
	})
}
