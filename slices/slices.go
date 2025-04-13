package slices

// Filter iterates over a slice and returns a new slice containing
// elements that satisfy the given condition. i.e if the condition
// returns true, the element is included in the new slice.
func Filter[T any](s []T, condition func(T) bool) []T {
	// Create new slices with cap of the len of
	// the original slice to save in memory allocation.
	result := make([]T, 0, len(s))
	for _, v := range s {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}
