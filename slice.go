package gotils

// MapSlice applies the provided mapping function to each element in the input slice
// and returns a new slice containing the transformed elements.
// TIn is the type of elements in the input slice.
// TOut is the type of elements in the output slice.
// The mapping function takes a TIn value and returns a TOut value.
func MapSlice[TIn any, TOut any](input []TIn, mapFn func(a TIn) TOut) []TOut {
	output := make([]TOut, len(input))
	for i, item := range input {
		output[i] = mapFn(item)
	}
	return output
}

// FilterSlice creates a new slice containing only the elements from the input slice
// that satisfy the provided predicate function.
// T is the type of elements in both the input and output slices.
// The predicate function takes a T value and returns true if the element should be included
// in the output slice, or false if it should be excluded.
func FilterSlice[T any](input []T, fn func(a T) bool) []T {
	output := make([]T, 0)
	for _, item := range input {
		if fn(item) {
			output = append(output, item)
		}
	}
	return output
}

// AnySlice converts a typed slice to a slice of the empty interface (any).
// This is useful when you need to work with heterogeneous data or when
// interfacing with APIs that expect a slice of interface{} values.
// T is the type of elements in the input slice.
// The returned slice contains the same elements as the input slice,
// but with the type information erased.
func AnySlice[T any](slice []T) []any {
	result := make([]any, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}

// FindSlice returns the first element in the slice that satisfies the provided predicate function,
// along with a boolean indicating whether such an element was found.
// T is the type of elements in the slice.
// The predicate function takes a T value and returns true if the element matches the search criteria.
// If an element is found, it returns that element and true.
// If no element is found, it returns the zero value of T and false.
func FindSlice[T any](slice []T, fn func(a T) bool) (T, bool) {
	for _, item := range slice {
		if fn(item) {
			return item, true
		}
	}
	var zero T
	return zero, false
}
