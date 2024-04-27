package sequence

import (
	"collections-go/contracts"
	"context"
)

type Type[T any] contracts.SequenceType[T]

type ContextualSequence[T any] struct {
	Sequence Type[T]
	ctx      context.Context
}

// FromSlicef creates a new sequence of type T
func FromSlicef[T any](list []T) contracts.Sequence[T] {
	return Type[T](list)
}

func (seq Type[T]) WithContext(ctx context.Context) contracts.ContextualSequence[T] {
	return &ContextualSequence[T]{Sequence: seq, ctx: ctx}
}
