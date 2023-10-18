# maps
This submodule contains functions meant for common operations on maps.

## Highlights:

### maps.Keys
Given a map, returns a slice containing all the keys of the map.

Note: because of how go implements maps, the keys will necessarily be in a random order.

Performance: `O(n)`

### maps.Values
Given a map, returns a slice containing all the values of the map.

Note: because of how go implements maps, the values will necessarily be in a random order.

Performance: `O(n)`

### maps.Merge
Given any number of maps, returns a new map containing all their combined key/value pairs, with later maps overriding prior ones in the event of key overlap.

Performance: `O(m*n*log(n))`

### maps.Invert
Given a map, returns a new map with the keys and values swapped.

Warning: because of how go implements maps, if values in the original map overlap, which of their corresponding keys is kept will be random.

Performance: `O(n*log(n))`

### maps.Pick
Given a map and a slice of keys, returns a new map containing only the subset of keys contained in the slice.

Performance: `O(n*log(n))`

### maps.FilterMap
Given a map and a filter function, returns a new map containing only the subset of key/value pairs that pass the filter function.
Conceptually similar to `slices.Filter`.

Performance: `O(n*log(n))`

## Note:
In several of the above functions, the big-O performance includes a `log(n)` term.
This is because, as more values are added to a map, the underlying array is periodically re-allocated.
This could be avoided by pre-allocating a larger array, but because there is no way to shrink a map's underlying array without instantiating a completely new one,
doing so could potentially lead to memory leaks by an unaware developer.
