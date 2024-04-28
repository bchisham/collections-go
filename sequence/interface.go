package sequence

import (
	"collections-go/contracts"
)

type Type[T any] contracts.SequenceType[T]

// FromSlice creates a new sequence of type T
func FromSlice[T any](list []T) contracts.Sequence[T] {
	return Type[T](list)
}
