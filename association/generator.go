package association

import (
	"collections-go/contracts"
	"collections-go/sequence"
	"context"
)

type accumulator[U any, V any] struct {
	sequence.Type[U]
	TransformFunc contracts.TransformFunc[V, U]
}

type generator[U any, K comparable, V any] struct {
	accumulator[U, V]
	Source chan V
	ctx    context.Context
}

func NewGenerator[U any, K comparable, V any](ctx context.Context, f contracts.TransformFunc[V, U], source chan V) contracts.AssociationGenerator[U, K, V] {
	return &generator[U, K, V]{accumulator[U, V]{TransformFunc: f}, source, ctx}
}

func (seq *generator[U, K, V]) Yield(value V) (contracts.AssociationGenerator[U, K, V], error) {
	select {
	case <-seq.ctx.Done():
		return nil, contracts.ErrContextDone
	case seq.Source <- value:
		v1, err := seq.TransformFunc(value)
		if err != nil {
			return nil, err
		}
		seq.Type = append(seq.Type, v1)
	}
	return seq, nil
}

func (seq *generator[U, K, V]) ToSlice() contracts.Sequence[U] {
	return seq.Type
}
