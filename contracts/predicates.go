package contracts

import "context"

type UnaryPredicate[T any] func(T) (bool, error)

// UnaryPredicateMust is a function that takes a value of type T and returns a boolean value.
type UnaryPredicateMust[T any] func(T) bool

// ContextPredicate is a function that takes a context and a value of type T and returns a boolean value.
type ContextPredicate[T any] func(ctx context.Context, item T) (bool, error)
