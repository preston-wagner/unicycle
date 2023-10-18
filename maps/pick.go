package maps

// Pick returns subset of a map filtered by a provided slice of keys
func Pick[KEY_TYPE comparable, VALUE_TYPE any](input map[KEY_TYPE]VALUE_TYPE, keys []KEY_TYPE) map[KEY_TYPE]VALUE_TYPE {
	result := map[KEY_TYPE]VALUE_TYPE{}
	for _, key := range keys {
		value, exists := input[key]
		if exists {
			result[key] = value
		}
	}
	return result
}
