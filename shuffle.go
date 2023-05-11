package unicycle

import "reflect"

// Shuffle takes a slice and returns a copy of that slice with the elements in a random order
func Shuffle[INPUT_TYPE any](input []INPUT_TYPE) []INPUT_TYPE {
	if len(input) < 2 {
		return input
	}
	shuffled := shuffleInner(input)
	if len(input) > 2 {
		for reflect.DeepEqual(shuffled, input) { // if the order (by some coincidence) hasn't changed, reshuffle
			shuffled = shuffleInner(input)
		}
	}
	return shuffled
}

func shuffleInner[INPUT_TYPE any](input []INPUT_TYPE) []INPUT_TYPE {
	indexes := Set[int]{}
	for i := range input {
		indexes.Add(i)
	}
	// this works because range returns the keys of a map in a random order, and sets are maps
	return Mapping(indexes.Values(), func(i int) INPUT_TYPE {
		return input[i]
	})
}
