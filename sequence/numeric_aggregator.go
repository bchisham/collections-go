package sequence

import "github.com/bchisham/collections-go/contracts"

type numericAggregator[T any, N contracts.NumericType] struct {
	Type[T]
}

func Mean[T any, N contracts.NumericType](seq []T, f func(T) N) N {
	return Sum(seq, f) / N(len(seq))
}

func (seq numericAggregator[T, N]) Mean(f func(T) N) N {
	return Mean(seq.ToSlice(), f)
}

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
