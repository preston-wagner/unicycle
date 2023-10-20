# multithread
This submodule contains functions to simplify common multithreading applications.

## Highlights:

### multithread.Repeat
`Repeat` simplifies the common task of setting up a task to be run on a given interval, like once every hour.
`RepeatMultithread` does the same thing, but gives each task its own goroutine.
The practical difference is that, if the task is long running, `Repeat` will wait until the prior task finishes to start the next iteration,
while `RepeatMultithread` will start it immediately.

### multithread.AwaitConcurrent
`AwaitConcurrent` simplifies the common task of waiting until several unrelated tasks on separate threads have concluded.
```go
AwaitConcurrent(func(){
  // perform task 1
}, func(){
  // perform task 2
}, func(){
  // perform task 3
})
```

### multithread.MappingMultithread
Many of the functions provided by this submodule are parallelized versions of functions provided by other submodules.
In this case, `multithread.MappingMultithread` acts like `slices.Mapping`, except that all mapping functions are run in parallel instead of serially.
