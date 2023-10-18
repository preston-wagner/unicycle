# promises
This submodule contains types and functions meant to simplify certain asynchronous/multithreaded tasks.

An instance of `promises.Promise` represents a piece of data that will become available at some point in the future.
It is best used in a write-once-read-once or write-once-read-many context, such as parallelizing network calls or caching.

The easiest and most useful way to create a promise is via the `promises.WrapInPromise` function:

```go
prm := promises.WrapInPromise(func()(string, error){
  // network call or similar long-running task we don't want to block
  return "whatever", nil
})

// do other tasks

value, err := prm.Await() // "whatever", nil
```

Calling `.Await()` blocks the consumer thread until the promise is resolved, but the resolving thread is never blocked.
