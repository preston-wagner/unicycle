package maps

// Like Filter, but for maps instead of slices
// Performance: O(n*log(n))
func FilterMap[KEY_TYPE comparable, VALUE_TYPE any](input map[KEY_TYPE]VALUE_TYPE, filter func(KEY_TYPE, VALUE_TYPE) bool) map[KEY_TYPE]VALUE_TYPE {
	keep := map[KEY_TYPE]VALUE_TYPE{}
	for key, value := range input {
		if filter(key, value) {
			keep[key] = value
		}
	}
	return keep
}
