package multithread

import (
	"github.com/preston-wagner/unicycle/slices"
)

// accepts any number of channels of the same type and returns a single unbuffered channel that pulls from all of them at once
// the returned channel closes once all the source channels do
func MergeChannels[T any](input []chan T) chan T {
	output := make(chan T)
	go func() {
		AwaitConcurrent(
			slices.Mapping(
				input,
				func(inputChan chan T) func() {
					return func() {
						for value := range inputChan {
							output <- value
						}
					}
				},
			)...,
		)
		close(output)
	}()
	return output
}
