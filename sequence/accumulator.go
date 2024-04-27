package sequence

import "collections-go/contracts"

type accumulator[T any, U any] struct {
	Type[T]
	TransformFunc[U, T]
}

func NewAccumulator[T any, U any](seq []T, f TransformFunc[U, T]) contracts.Generator[T, U] {
	return &accumulator[T, U]{seq, f}
}

func (seq *accumulator[T, U]) Yield(value U) contracts.Generator[T, U] {
	return &accumulator[T, U]{append(seq.Unwrap(), seq.TransformFunc(value)), seq.TransformFunc}
}

func (seq *accumulator[T, U]) Unwrap() []T {
	return seq.Type
}
