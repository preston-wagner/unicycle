package unicycle

import (
	"sync"
)

// Equivalent to lodash's _.groupBy()
func GroupBy[KEY_TYPE comparable, VALUE_TYPE any](input []VALUE_TYPE, keyGenerator func(VALUE_TYPE) KEY_TYPE) map[KEY_TYPE][]VALUE_TYPE {
	output := map[KEY_TYPE][]VALUE_TYPE{}
	for _, value := range input {
		key := keyGenerator(value)
		addOrAppend(output, key, value)
	}
	return output
}

// like GroupBy, but runs the keyGenerator function concurrently (order of results is not guaranteed)
func GroupByConcurrently[KEY_TYPE comparable, VALUE_TYPE any](input []VALUE_TYPE, keyGenerator func(VALUE_TYPE) KEY_TYPE) map[KEY_TYPE][]VALUE_TYPE {
	output := map[KEY_TYPE][]VALUE_TYPE{}
	mux := &sync.Mutex{}
	AwaitConcurrent(Mapping(input, func(value VALUE_TYPE) func() {
		return func() {
			key := keyGenerator(value)
			mux.Lock()
			defer mux.Unlock()
			addOrAppend(output, key, value)
		}
	})...)
	return output
}

func addOrAppend[KEY_TYPE comparable, VALUE_TYPE any](output map[KEY_TYPE][]VALUE_TYPE, key KEY_TYPE, value VALUE_TYPE) {
	_, ok := output[key]
	if !ok {
		output[key] = []VALUE_TYPE{value}
	} else {
		output[key] = append(output[key], value)
	}
}
