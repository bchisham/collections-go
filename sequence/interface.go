package sequence

import (
	"github.com/bchisham/collections-go/contracts"
)

type Type[T any] contracts.SequenceType[T]

// FromSlice creates a new sequence of type T
func FromSlice[T any](list []T) contracts.Sequence[T] {
	return Type[T](list)
}

func FromItems[T any](items ...T) contracts.Sequence[T] {
	return Type[T](items)
}

func (seq Type[T]) Length() int {
	return len(seq)
}
