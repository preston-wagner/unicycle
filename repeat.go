package unicycle

import (
	"context"
	"time"
)

// Repeat runs a given function in its own goroutine once per given duration, and returns a function that can be called to cancel the task
func Repeat(wrapped func(), interval time.Duration) func() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				go wrapped()
			case <-ctx.Done():
				return
			}
		}
	}()
	return cancel
}
