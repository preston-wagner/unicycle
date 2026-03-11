# sets
This submodule contains the definition for a `Set`.

This is implemented as a `map` with no value
(or more correctly, a value of `struct{}`, which has a width of 0 bytes and thus no memory overhead).
Since a `Set` is a `map`, you can use it in a range statement, as long as you keep in mind that the key of the map doubles as the value:
```go
set := make(sets.Set[string])

for value := range set {
  // ...
}
```

Sets expose self-explanatory `.Add()`, `.Remove()`, `.Has()`, and `.Copy()` methods,
as well as `.Values()`, which returns all the entries in the set as a slice,
and `.Difference()`, which allows you to compare a set with any number of others.
`sets.Union` and `sets.Intersection` exist as standalone functions.

`sets.SetFromSlice` is also provided to simplify instantiating sets with data.
Anyone who has used a raw `map[T]struct{}` to implement a set has probably noticed that because of golang's strong opinions
regarding whitespace, adding or removing entries from a hardcoded set _a la_
```go
map[string]struct{
  "short": struct{},
}
```
to
```go
map[string]struct{
  "short":           struct{},
  "very long value": struct{},
}
```
can make git diffs difficult to read pretty quickly; `SetFromSlice` saves you from both that and typing `struct{}` over and over:
```go
sets.SetFromSlice([]string{
  "short",
  "very long value",
})
```

See also: https://en.wikipedia.org/wiki/Set_(abstract_data_type)
