package slice

// ForEach iterates over a slice of elements of type T, and calls the callback
// function on each element
func ForEach[T any](data []T, callback func(*T, int, []T)) {
	for i := range data {
		callback(&data[i], i, data)
	}
}

// While continuously executes the provided callback function as long as the
// predicate function returns true. The predicate function is evaluated before
// callback function
func While(predicate func() bool, callback func()) {
	for predicate() {
		callback()
	}
}
