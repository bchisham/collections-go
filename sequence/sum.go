package sequence

func Sum[T any, N NumericType](list []T, f func(T) N) N {
	var sum N
	for _, item := range list {
		sum += f(item)
	}
	return sum
}

func (seq numericAggregator[T, N]) Sum(f func(T) N) N {
	return Sum(seq.Unwrap(), f)
}
