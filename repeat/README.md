# repeat
This submodule contains functions to simplify common recurring background tasks.

## Highlights:

### repeat.Repeat
`Repeat` simplifies the common task of setting up a task to be run on a given interval, like once every hour.
`RepeatMultithread` does the same thing, but gives each task its own goroutine.
The practical difference is that, if the task is long running, `Repeat` will wait until the prior task finishes to start the next iteration,
while `RepeatMultithread` will start it immediately.
