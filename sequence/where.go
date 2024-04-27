package sequence

// Where takes a sequence of type T and a function that takes a value of type T
// and returns a boolean value. It returns a sequence of type T with values that
// satisfy the function
func (seq Type[T]) Where(f PredicateFunc[T]) Type[T] {
	var r Type[T]
	for _, v := range seq {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

// Where takes a sequence of type T and a function that takes a value of type T
// and returns a boolean value. It returns a sequence of type T with values that
// satisfy the function
func Where[T any](seq []T, f PredicateFunc[T]) Type[T] {
	return New(seq).Where(f)
}
