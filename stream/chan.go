package stream

import (
	"context"
	"github.com/bchisham/collections-go/contracts"
)

type ChanType[T any] chan T

type InChanType[T any] <-chan T
type OutChanType[T any] chan<- T

type chanWrapper[T any] struct {
	ChanType[T]
	ctx context.Context
}

type chanReader[T any] struct {
	InChanType[T]
	ctx context.Context
}

type chanWriter[T any] struct {
	OutChanType[T]
	ctx context.Context
}

func NewChan[T any]() contracts.Channel[T] {
	return make(ChanType[T])
}

func NewInChan[T any](c InChanType[T]) contracts.ReadChannel[T] {
	return &chanReader[T]{c, context.Background()}
}

func NewInChanWithContext[T any](c InChanType[T], ctx context.Context) contracts.ReadChannel[T] {
	return &chanReader[T]{c, ctx}
}

func NewOutChan[T any](c OutChanType[T]) contracts.WriteChannel[T] {
	return &chanWriter[T]{c, context.Background()}
}

func NewOutChanWithContext[T any](c OutChanType[T], ctx context.Context) contracts.WriteChannel[T] {
	return &chanWriter[T]{c, ctx}
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

func (c *chanReader[T]) Receive() (t T, _ error) {
	select {
	case <-c.ctx.Done():
		return t, contracts.ErrContextDone
	case v := <-c.InChanType:
		return v, nil
	}
}

func (c *chanReader[T]) Each(f contracts.ApplyFunc[T]) error {
	for {
		select {
		case <-c.ctx.Done():
			return contracts.ErrContextDone
		case v := <-c.InChanType:
			err := f(v)
			if err != nil {
				return err
			}
		}
	}
}

func (c *chanReader[T]) ToChannel() <-chan T {
	return c.InChanType
}

func (c *chanWriter[T]) ToChannel() chan<- T {
	return c.OutChanType
}

func (c *chanWriter[T]) Send(t T) error {
	c.OutChanType <- t
	return nil
}

func (c *chanWriter[T]) Close() error {
	close(c.OutChanType)
	return nil
}
