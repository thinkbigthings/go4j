package stream

// Map function that applies a transformation to each element in the slice.
func Map[T, R any](input []T, transform func(T) R) []R {
	result := make([]R, len(input))
	for i, item := range input {
		result[i] = transform(item)
	}
	return result
}

func Filter[T any](input []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range input {
		if predicate(item) {
			result = append(result, item)
		}
	}
	// Ensure result is not nil
	if result == nil {
		return []T{}
	}
	return result
}

// Reduce function that reduces the slice to a single value.
func Reduce[T any](input []T, initial T, reducer func(T, T) T) T {
	result := initial
	for _, item := range input {
		result = reducer(result, item)
	}
	return result
}
