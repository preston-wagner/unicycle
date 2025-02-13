package multithread

import "github.com/preston-wagner/unicycle/slices_ext"

// like Filter, but multithreaded
func FilterMultithread[T any](input []T, filter func(T) bool) []T {
	results := MappingMultithread(input, func(value T) filterResult[T] {
		return filterResult[T]{
			value: value,
			ok:    filter(value),
		}
	})
	results = slices_ext.Filter(results, func(res filterResult[T]) bool {
		return res.ok
	})
	return slices_ext.Mapping(results, func(res filterResult[T]) T {
		return res.value
	})
}

// like FilterWithError, but multithreaded
func FilterMultithreadWithError[T any](input []T, filter func(T) (bool, error)) ([]T, error) {
	results, err := MappingMultithreadWithError(input, func(value T) (filterResult[T], error) {
		ok, err := filter(value)
		return filterResult[T]{
			value: value,
			ok:    ok,
		}, err
	})
	if err != nil {
		return []T{}, err
	}
	results = slices_ext.Filter(results, func(res filterResult[T]) bool {
		return res.ok
	})
	return slices_ext.Mapping(results, func(res filterResult[T]) T {
		return res.value
	}), nil
}
