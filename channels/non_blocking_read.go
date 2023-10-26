package channels

import "github.com/nuvi/unicycle/defaults"

func NonBlockingRead[T any](incoming chan T) (T, bool) {
	select {
	case value, ok := <-incoming:
		return value, ok
	default: // makes the above read non-blocking
		return defaults.ZeroValue[T](), false
	}
}
