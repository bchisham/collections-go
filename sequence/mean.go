package sequence

import "collections-go/contracts"

func Mean[T any, N contracts.NumericType](seq []T, f func(T) N) N {
	return Sum(seq, f) / N(len(seq))
}

func (seq numericAggregator[T, N]) Mean(f func(T) N) N {
	return Mean(seq.ToSlice(), f)
}
