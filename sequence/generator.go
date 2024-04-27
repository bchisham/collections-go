package sequence

import "context"

type generator[T any, U any] struct {
	accumulator[T, U]
	Source chan U
	ctx    context.Context
}

func NewGenerator[T any, U any](ctx context.Context, f TransformFunc[U, T], source chan U) Generator[T, U] {
	return &generator[T, U]{accumulator[T, U]{[]T{}, f}, source, ctx}
}

func (seq *generator[T, U]) Yield(value U) Generator[T, U] {
	select {
	case <-seq.ctx.Done():
		return seq
	case seq.Source <- value:
		seq.Type = append(seq.Type, seq.TransformFunc(value))
	}
	return seq
}

func (seq *generator[T, U]) Unwrap() []T {
	return seq.Type
}
