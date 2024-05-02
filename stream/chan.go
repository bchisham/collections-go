package stream

import (
	"collections-go/contracts"
	"context"
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
