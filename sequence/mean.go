package sequence

func Mean[T any, N NumericType](seq []T, f func(T) N) N {
	return Sum(seq, f) / N(len(seq))
}

func (seq numericAggregator[T, N]) Mean(f func(T) N) N {
	return Mean(seq.Unwrap(), f)
}
