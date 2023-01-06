package unicycle

var Empty struct{}

// Sets are basically maps with no values; empty structs have a width of 0 bytes
// https://dave.cheney.net/2014/03/25/the-empty-struct
type Set[T comparable] map[T]struct{}

func SetFromSlice[T comparable](input []T) Set[T] {
	set := Set[T]{}
	set.Add(input...)
	return set
}

func (set Set[T]) Add(values ...T) {
	for _, value := range values {
		set[value] = Empty
	}
}

func (set Set[T]) Remove(values ...T) {
	for _, value := range values {
		delete(set, value)
	}
}

func (set Set[T]) Has(value T) bool {
	_, ok := set[value]
	return ok
}

func (set Set[T]) Values() []T {
	return Keys(set)
}

// func (set Set[T]) Union(others ...Set[T]) Set[T] {
// 	return Merge(set, Merge(others...))
// }

// func (set Set[T]) Intersection(others ...Set[T]) Set[T] {
// }
