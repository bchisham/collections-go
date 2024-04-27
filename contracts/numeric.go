package contracts

type NumericAggregator[T any, N NumericType] interface {
	Sum(f func(T) N) N
	Mean(f func(T) N) N
}
