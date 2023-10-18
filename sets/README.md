# sets
This submodule contains the definition for a `Set`.

This is implemented as a `map` with no value (or more correctly, a value of `struct{}`, which has a width of 0 bytes and thus no memory overhead). Since a `Set` is a `map`, you can use it in a range statement, as long as you keep in mind that the key of the map doubles as the value:

```go
set := make(sets.Set[string])

for value := range set {
  // ...
}
```

Sets expose self-explanatory `.Add()`, `.Remove()`, `.Has()` methods,
as well as `.Values()`, which returns all the entries in the set as a slice,
and `.Difference()`, which allows you to compare a set with any number of others.
`sets.Union` and `sets.Intersection` exist as standalone functions.

`sets.SetFromSlice` is also provided to simplify instantiating sets with data.

See also: https://en.wikipedia.org/wiki/Set_(abstract_data_type)
