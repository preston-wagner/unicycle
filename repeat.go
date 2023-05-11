package unicycle

import (
	"context"
	"time"
)

// Repeat runs a given function once per given duration, and returns a function that can be called to cancel the task
// if before == true, the wrapped function will be called immediately, instead of waiting for the ticker for the first run
func Repeat(wrapped func(), interval time.Duration, before bool) func() {
	ctx, cancel := context.WithCancel(context.Background())
	go repeatInner(interval, ctx, wrapped, before)
	return cancel
}

func repeatInner(interval time.Duration, ctx context.Context, wrapped func(), before bool) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		if before {
			wrapped()
		}
		select {
		case <-ticker.C:
			if !before {
				wrapped()
			}
		case <-ctx.Done():
			return
		}
	}
}

// Like Repeat, but each call runs in its own goroutine
func RepeatMultithread(wrapped func(), interval time.Duration, before bool) func() {
	return Repeat(func() {
		go wrapped()
	}, interval, before)
}
