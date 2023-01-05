package unicycle

var Empty struct{}

// Sets are basically maps with no values; empty structs have a width of 0 bytes
// https://dave.cheney.net/2014/03/25/the-empty-struct
type Set[KEY_TYPE comparable] map[KEY_TYPE]struct{}

func (set Set[KEY_TYPE]) Add(values ...KEY_TYPE) {
	for _, value := range values {
		set[value] = Empty
	}
}

func (set Set[KEY_TYPE]) Remove(values ...KEY_TYPE) {
	for _, value := range values {
		delete(set, value)
	}
}

func (set Set[KEY_TYPE]) Has(value KEY_TYPE) bool {
	_, ok := set[value]
	return ok
}

func (set Set[KEY_TYPE]) Values() []KEY_TYPE {
	return Keys(set)
}
