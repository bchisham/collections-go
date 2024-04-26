package sequence

// FilterFunc is a function that takes a value of type T and returns a boolean value.
type FilterFunc[T any] func(T) bool

// Filter takes a sequence of type T and a function that takes a value of type T
// and returns a boolean value. It returns a sequence of type T with values that
// satisfy the function
func (seq Type[T]) Filter(f FilterFunc[T]) Type[T] {
	var r Type[T]
	for _, v := range seq {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

// Filter takes a sequence of type T and a function that takes a value of type T
// and returns a boolean value. It returns a sequence of type T with values that
// satisfy the function
func Filter[T any](seq []T, f FilterFunc[T]) Type[T] {
	return New(seq).Filter(f)
}
