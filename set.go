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

// Difference returns the set of values contained in the base set but not in any others
func (set Set[T]) Difference(others ...Set[T]) Set[T] {
	result := Set[T]{}
	for value := range set {
		if !Some(others, func(other Set[T]) bool {
			return other.Has(value)
		}) {
			result.Add(value)
		}
	}
	return result
}

func (set Set[T]) intersection(others ...Set[T]) Set[T] {
	result := Set[T]{}
	for value := range set {
		if Every(others, func(other Set[T]) bool {
			return other.Has(value)
		}) {
			result.Add(value)
		}
	}
	return result
}

// Intersection returns the set of all values in all provided sets
func Union[T comparable](sets ...Set[T]) Set[T] {
	result := Set[T]{}
	for _, set := range sets {
		result.Add(set.Values()...)
	}
	return result
}

// Intersection returns the set of values shared by all provided sets
func Intersection[T comparable](sets ...Set[T]) Set[T] {
	if len(sets) == 0 {
		return Set[T]{}
	} else if len(sets) == 1 {
		return SetFromSlice(sets[0].Values())
	} else { // len(sets) > 1
		return sets[0].intersection(sets[1:]...)
	}
}
