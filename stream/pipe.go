package stream

import (
	"context"
	"github.com/bchisham/collections-go/contracts"
	"github.com/bchisham/collections-go/sequence"
)

// Pipe takes an input channel and an output channel and pipes the input channel to the output channel
func Pipe[T any](ctx context.Context, in InChanType[T], out OutChanType[T]) error {
	src := NewInChanWithContext(in, ctx)
	sink := NewOutChanWithContext(out, ctx)
	return src.Each(sink.Send)
}

// Tee takes an input channel and multiple output channels and pipes the input channel to the output channels
func Tee[T any](ctx context.Context, in InChanType[T], out ...OutChanType[T]) error {
	var fanout []contracts.WriteChannel[T]
	_ = sequence.FromSlice(out).Each(func(o OutChanType[T]) error {
		fanout = append(fanout, NewOutChanWithContext(o, ctx))
		return nil
	})
	output := sequence.FromSlice(fanout)
	input := NewInChanWithContext(in, ctx)
	return input.Each(func(t T) error {
		return output.Each(func(o contracts.WriteChannel[T]) error {
			return o.Send(t)
		})
	})
}

// Transform takes an input channel, a transform function, and an output channel and pipes the input channel to the output channel after applying the transform function
func Transform[T any, U any](ctx context.Context, transformFunc contracts.TransformFunc[T, U], in InChanType[T], out OutChanType[U]) error {
	src := NewInChanWithContext(in, ctx)
	sink := NewOutChanWithContext(out, ctx)
	return src.Each(func(t T) error {
		val, err := transformFunc(t)
		if err != nil {
			return err
		}
		return sink.Send(val)
	})
}
