package association

import (
	"context"
	"github.com/bchisham/collections-go/contracts"
)

type contextualMap[T comparable, U any] struct {
	m   Type[T, U]
	ctx context.Context
}

type ContextualMapBuilder[T comparable, U any] struct {
	ctx context.Context
}

func WithContext[T comparable, U any](ctx context.Context) ContextualMapBuilder[T, U] {
	return ContextualMapBuilder[T, U]{ctx: ctx}
}

func (b ContextualMapBuilder[T, U]) FromMap(m Type[T, U]) contracts.MapWithContext[T, U] {
	return contextualMap[T, U]{m: m, ctx: b.ctx}
}

func (m contextualMap[T, U]) Each(f contracts.ApplyWithContextFunc[U]) error {
	for _, v := range m.m {
		if err := f(m.ctx, v); err != nil {
			return err
		}
	}
	return nil
}

func (m contextualMap[T, U]) Every(f contracts.ContextPredicate[U]) (bool, error) {
	for _, v := range m.m {
		if ok, err := f(m.ctx, v); err != nil {
			return false, err
		} else if !ok {
			return false, nil
		}
	}
	return true, nil
}

func (m contextualMap[T, U]) Where(f contracts.ContextPredicate[U]) (contracts.MapWithContext[T, U], error) {
	result := make(map[T]U)
	for k, v := range m.m {
		if ok, err := f(m.ctx, v); err != nil {
			return nil, err
		} else if ok {
			result[k] = v
		}
	}
	return contextualMap[T, U]{m: result, ctx: m.ctx}, nil
}
