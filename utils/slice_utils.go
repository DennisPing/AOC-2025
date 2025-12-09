package utils

// GetElementFromReverse gets the N-th element from the end of a slice.
// N is 1-based (1 for the last element, 2 for the second to last, etc.).
func GetElementFromReverse[T any](slice []T, n int) (T, bool) {
	var zero T // Zero value for the type T
	if n <= 0 || n > len(slice) {
		return zero, false
	}
	return slice[len(slice)-n], true
}

// FindAllIndices finds all the indices of element T in a slice.
func FindAllIndices[T comparable](slice []T, element T) []int {
	var indices []int
	for i, v := range slice {
		if v == element {
			indices = append(indices, i)
		}
	}
	return indices
}
