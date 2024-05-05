package stream

import (
	"context"
	"github.com/bchisham/collections-go/contracts"
)

type ChanType[T any] chan T

type chanWrapper[T any] struct {
	ChanType[T]
	ctx context.Context
}

func NewChan[T any]() contracts.Channel[T] {
	return make(ChanType[T])
}

func FromChan[T any](c ChanType[T]) contracts.Channel[T] {
	return &chanWrapper[T]{c, context.Background()}
}

func FromChanWithContext[T any](c ChanType[T], ctx context.Context) contracts.Channel[T] {
	return &chanWrapper[T]{c, ctx}
}

func (c ChanType[T]) Send(value T) error {
	c <- value
	return nil
}

func (c ChanType[T]) Receive() (T, error) {
	return <-c, nil
}

func (c ChanType[T]) Close() error {
	close(c)
	return nil
}

func (c ChanType[T]) ToChannel() chan T {
	return c
}

func (c ChanType[T]) Each(f contracts.ApplyFunc[T]) error {
	for {
		select {
		case v := <-c:
			err := f(v)
			if err != nil {
				return err
			}
		}
	}
}

func (c *chanWrapper[T]) Send(value T) error {
	c.ChanType <- value
	return nil
}

func (c *chanWrapper[T]) Receive() (t T, _ error) {
	select {
	case <-c.ctx.Done():
		return t, contracts.ErrContextDone
	case v := <-c.ChanType:
		return v, nil
	}
}

func (c *chanWrapper[T]) Each(f contracts.ApplyFunc[T]) error {
	for {
		select {
		case <-c.ctx.Done():
			return contracts.ErrContextDone
		case v := <-c.ChanType:
			err := f(v)
			if err != nil {
				return err
			}
		}
	}
}
