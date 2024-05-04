package sequence

import "collections-go/contracts"

type accumulator[T any, U any] struct {
	Type[T]
	Output        Type[U]
	TransformFunc TransformFunc[T, U]
	idx           int
}

func NewAccumulator[T any, U any](seq []T, f TransformFunc[T, U]) contracts.Generator[U] {
	output := make([]U, 0)
	return &accumulator[T, U]{seq, output, f, 0}
}

func (seq *accumulator[T, U]) Yield() error {
	if seq.idx >= seq.Type.Length() {
		return ErrEndOfSequence
	}
	value := seq.Type[seq.idx]
	seq.idx++
	seq.Output = append(seq.Output, seq.TransformFunc(value))
	return nil
}

func (seq *accumulator[T, U]) ToSlice() contracts.Sequence[U] {

	return FromSlice(seq.Output)
}
