package multithread

import (
	"github.com/preston-wagner/unicycle/promises"
	"github.com/preston-wagner/unicycle/slices"
)

func toMutatorChannel[INPUT_TYPE any, OUTPUT_TYPE any](mutator func(INPUT_TYPE) OUTPUT_TYPE) func(in INPUT_TYPE) chan OUTPUT_TYPE {
	return func(in INPUT_TYPE) chan OUTPUT_TYPE {
		out := make(chan OUTPUT_TYPE)
		go func() {
			out <- mutator(in)
		}()
		return out
	}
}

func fromMutatorChannel[OUTPUT_TYPE any](channel chan OUTPUT_TYPE) OUTPUT_TYPE {
	return <-channel
}

// like slices.Mapping(), but all mutator functions run in parallel in their own goroutines
func MappingMultithread[INPUT_TYPE any, OUTPUT_TYPE any](input []INPUT_TYPE, mutator func(INPUT_TYPE) OUTPUT_TYPE) []OUTPUT_TYPE {
	return slices.Mapping(slices.Mapping(input, toMutatorChannel(mutator)), fromMutatorChannel[OUTPUT_TYPE])
}

// like slices.MappingWithError(), but all mutator functions run in parallel in their own goroutines
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
