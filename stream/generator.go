package stream

import (
	"collections-go/contracts"
	"collections-go/sequence"
	"context"
	"sync"
)

type chanGenerator[T any] struct {
	chanWrapper[T]
	sequence.Type[T]
}

func NewGenerator[T any](ctx context.Context, c contracts.Channel[T]) contracts.IteratorToSlice[T] {
	return &chanGenerator[T]{chanWrapper[T]{c.ToChannel(), ctx}, nil}
}

func (c *chanGenerator[T]) Yield() (contracts.IteratorToSlice[T], error) {
	select {
	case <-c.ctx.Done():
		return nil, contracts.ErrContextDone
	case value := <-c.ChanType:
		c.Type = append(c.Type, value)
	}
	return c, nil
}

func (c *chanGenerator[T]) ToSlice() contracts.Sequence[T] {
	return c.Type
}

// SendSequence sends a sequence of items to a channel.
func SendSequence[T any](c contracts.Channel[T], seq contracts.Sequence[T]) error {
	return seq.Each(func(item T) error {
		return c.Send(item)
	})
}

// IteratorToSlice reads from a channel and returns a sequence.
func IteratorToSlice[T any](ctx context.Context, c contracts.Channel[T]) (_ contracts.Sequence[T], err error) {
	reader := NewGenerator(ctx, c)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			_, err = reader.Yield()
			if err != nil {
				break
			}
		}
	}()
	wg.Wait()
	return reader.ToSlice(), nil
}
