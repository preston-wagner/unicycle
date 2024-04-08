# multithread
This submodule contains functions to simplify common multithreading applications.

## Highlights:

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
