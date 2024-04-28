package contracts

import "context"

type UnaryPredicate[T any] func(T) (bool, error)
type BinaryPredicate[T any, U any] func(T, U) (bool, error)

// UnaryPredicateMust is a function that takes a value of type T and returns a boolean value.
type UnaryPredicateMust[T any] func(T) bool
type BinaryPredicateMust[T any, U any] func(T, U) bool

// ContextPredicate is a function that takes a context and a value of type T and returns a boolean value.
type ContextPredicate[T any] func(ctx context.Context, item T) (bool, error)

type MapPredicate[K comparable, V any] func(K, V) (bool, error)
type MapPredicateMust[K comparable, V any] func(K, V) bool
type MapPredicateWithContext[K comparable, V any] func(context.Context, K, V) (bool, error)
