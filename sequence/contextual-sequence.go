package sequence

import (
	"context"
	"github.com/bchisham/collections-go/contracts"
)

func WithContext[T any](ctx context.Context) ContextualSequenceBuilder[T] {
	return ContextualSequenceBuilder[T]{ctx: ctx}
}

func (b ContextualSequenceBuilder[T]) Build(seq Type[T]) ContextualSequence[T] {
	return ContextualSequence[T]{Sequence: seq, ctx: b.ctx}
}

type ContextualSequence[T any] struct {
	Sequence Type[T]
	ctx      context.Context
}

type ContextualSequenceBuilder[T any] struct {
	ctx context.Context
}

func (seq ContextualSequence[T]) Each(f contracts.ApplyWithContextFunc[T]) error {
	for _, item := range seq.Sequence {
		if err := f(seq.ctx, item); err != nil {
			return err
		}
	}
	return nil
}

func (seq ContextualSequence[T]) Every(f contracts.ContextPredicate[T]) (bool, error) {
	for _, item := range seq.Sequence {
		ok, err := f(seq.ctx, item)
		if err != nil {
			return false, err
		}
		if !ok {
			return false, nil
		}
	}
	return true, nil
}

func (seq ContextualSequence[T]) FirstWhere(predicate contracts.ContextPredicate[T]) (result T, found bool, err error) {
	for _, item := range seq.Sequence {
		ok, err := predicate(seq.ctx, item)
		if err != nil {
			return item, false, err
		}
		if ok {
			return item, true, nil
		}
	}
	return result, false, nil
}

func (seq ContextualSequence[T]) Where(f contracts.ContextPredicate[T]) (contracts.Sequence[T], error) {
	var r Type[T]
	for _, v := range seq.Sequence.ToSlice() {
		ok, err := f(seq.ctx, v)
		if err != nil {
			return nil, err
		}
		if ok {
			r = append(r, v)
		}
	}
	return r, nil
}

func (seq ContextualSequence[T]) ToSlice() []T {
	return seq.Sequence
}

func (seq ContextualSequence[T]) ToSequence() contracts.Sequence[T] {
	return seq.Sequence
}

func (seq ContextualSequence[T]) Length() int {
	return len(seq.Sequence)
}
