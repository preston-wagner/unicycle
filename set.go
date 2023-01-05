package unicycle

var Empty struct{}

// Sets are basically maps with no values; empty structs have a width of 0 bytes
// https://dave.cheney.net/2014/03/25/the-empty-struct
type Set[KEY_TYPE comparable] map[KEY_TYPE]struct{}

func (set Set[KEY_TYPE]) Add(value KEY_TYPE) {
	set[value] = Empty
}

func (set Set[KEY_TYPE]) Remove(value KEY_TYPE) {
	delete(set, value)
}

func (set Set[KEY_TYPE]) Has(value KEY_TYPE) bool {
	_, ok := set[value]
	return ok
}

func (set Set[KEY_TYPE]) Values() []KEY_TYPE {
	return Keys(set)
}
