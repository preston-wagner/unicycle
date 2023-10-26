# slices
This submodule contains functions meant for common operations on slices.

## Highlights:

### slices.Mapping
A higher-order function that accepts a slice of any data type and a mutator function, then returns the slice of that data with the mutator applied.

Performance: `O(n)` (assuming a constant-time mutator function)

See also: https://en.wikipedia.org/wiki/Map_(higher-order_function)

### slices.Filter
A higher-order function that accepts a slice of any data type and a filter function, then returns a slice of the data that passes the filter, preserving the original ordering.

Performance: `O(n)` (assuming a constant-time filter function)

See also: https://en.wikipedia.org/wiki/Filter_(higher-order_function)

### slices.Reduce
A higher-order function that accepts a slice of data, an accumulator function, and an initial value, and applies the accumulator function to all the values of the slice, returning the accumulated data.

Performance: `O(n)` (assuming a constant-time accumulator function)

See also: https://en.wikipedia.org/wiki/Fold_(higher-order_function)#On_lists

### slices.Unique
Accepts a slice of data, and returns a new slice containing only the first instance of each unique value.

Performance: `~O(n)`

### slices.Concatenate
Accepts any number of slices and copies their values into a new slice, essentially joining the slices together.

Performance: `O(m*n)`

### slices.Includes
Accepts a slice and a value, and returns whether or not that value can be found in the provided slice.

Performance: `O(n)`

### slices.KeyBy
Accepts a slice of data and key generator function, and returns a map correlating each value to the key the function generated.

In the event that two values share a key, the last value is kept.

Useful for optimizing lookups.

Performance: `O(n*log(n))`

### slices.GroupBy
Accepts a slice of data and key generator function, and returns a map of keys to slices of values (i.e. grouping values by their keys).

Useful for optimizing lookups.

Performance: `O(n*log(n))`

### slices.Every
Accepts a slice of any data type and a test function, then returns true if all elements in the slice pass the test.

Empty slices always return true.

Performance: `O(n)`

### slices.Some
Accepts a slice of any data type and a test function, then returns true if at least one element in the slice passes the test.

Empty slices always return false.

Performance: `O(n)`
