package sequence

import "collections-go/contracts"

func Sum[T any, N contracts.NumericType](list []T, f func(T) N) N {
	var sum N
	for _, item := range list {
		sum += f(item)
	}
	return sum
}

func (seq numericAggregator[T, N]) Sum(f func(T) N) N {
	return Sum(seq.ToSlice(), f)
}
