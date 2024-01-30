package multithread

import (
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
	results := MappingMultithread(input, func(value INPUT_TYPE) errorResult[OUTPUT_TYPE] {
		mutated, err := mutator(value)
		return errorResult[OUTPUT_TYPE]{
			value: mutated,
			err:   err,
		}
	})
	return slices.MappingWithError(results, func(result errorResult[OUTPUT_TYPE]) (OUTPUT_TYPE, error) {
		return result.value, result.err
	})
}
