package sequence

import (
	"collections-go/contracts"
	"context"
)

type TransformMustFunc[ItemTypeIn any, ItemTypeOut any] func(ItemTypeIn) ItemTypeOut

type TransformWithContextFunc[ItemTypeIn any, ItemTypeOut any] func(context.Context, ItemTypeIn) (ItemTypeOut, error)

type TransformerMust[ItemTypeIn any, ItemTypeOut any] interface {
	TransformMust(f TransformMustFunc[ItemTypeIn, ItemTypeOut]) contracts.Sequence[ItemTypeOut]
}

// Transformer is an interface that can be used to transform a sequence of type T to a sequence of type U
type Transformer[T any, U any] interface {
	TransformMust(f TransformFunc[T, U]) Type[U]
}

// TransformFunc is a function that takes a value of type T and returns a value of type U
type TransformFunc[T any, U any] func(T) U

func NewTransformer[T any, U any](s contracts.Sequence[T]) TransformerMust[T, U] {
	return transformer[T, U]{s, nil}
}

type transformer[T any, U any] struct {
	Input  contracts.Sequence[T]
	Output Type[U]
}

// TransformMust takes a sequence of type T and a function that takes a value of type T and returns a value of type U. It returns a sequence of type U
func TransformMust[T any, U any](s Type[T], f TransformMustFunc[T, U]) contracts.Sequence[U] {
	output := make(Type[U], len(s))
	for i, v := range s {
		output[i] = f(v)
	}
	return output
}

// TransformMust takes a function that takes a value of type T and returns a value of type U. It returns a TransformMust that can be used to transform a sequence of type T to a sequence of type U
func (t transformer[T, U]) TransformMust(f TransformMustFunc[T, U]) contracts.Sequence[U] {
	in := t.Input.ToSlice()
	t.Output = make(Type[U], len(in))
	for i, v := range in {
		t.Output[i] = f(v)
	}
	return t.Output
}

func (t transformer[T, U]) ToSlice() []U {
	return t.Output.ToSlice()
}
