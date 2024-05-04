package sequence

import "github.com/bchisham/collections-go/contracts"

type numericAggregator[T any, N contracts.NumericType] struct {
	Type[T]
}
