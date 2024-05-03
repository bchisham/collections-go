package contracts

import (
	"collections-go/pair"
	"context"
)

// ApplyFunc is a function that takes a value of type T and returns an error.
type ApplyFunc[T any] func(item T) error

// ApplyWithContextFunc is a function that takes a context and a value of type T and returns an error.
type ApplyWithContextFunc[T any] func(ctx context.Context, item T) error

// Sequence is a type that represents a sequence of type T
type Sequence[T any] interface {
	ToSlice() []T
	// Each applies a function to each item in the sequence and returns an error if one occurs
	Each(f ApplyFunc[T]) error
	// Where filters the sequence based on a predicate
	Where(f UnaryPredicate[T]) (Sequence[T], error)
	// WhereMust filters the sequence based on a predicate
	WhereMust(f UnaryPredicateMust[T]) Sequence[T]
	// FirstWhereMust returns the first item in the sequence that satisfies the predicate.
	FirstWhereMust(predicate UnaryPredicateMust[T]) (result T, found bool)
	Every(f UnaryPredicate[T]) (bool, error)
	// EveryMust returns true if every item in the sequence satisfies the predicate.
	EveryMust(f UnaryPredicateMust[T]) bool
	Length() int
}

// ContextualSequence is a type that represents a sequence of type T with a context.
type ContextualSequence[T any] interface {
	ToSlice() []T
	Each(f ApplyWithContextFunc[T]) error
	Where(f ContextPredicate[T]) (Sequence[T], error)
	FirstWhere(predicate ContextPredicate[T]) (T, bool, error)
}

type Map[K comparable, V any] interface {
	ToMap() map[K]V
	Each(f ApplyFunc[V]) error
	Every(f MapPredicate[K, V]) (bool, error)
	EveryMust(f MapPredicateMust[K, V]) bool
	Where(f MapPredicate[K, V]) (Map[K, V], error)
	WhereMust(f MapPredicateMust[K, V]) Map[K, V]
	Keys() Sequence[K]
	Values() Sequence[V]
}

// MapWithContext is a type that represents a map of type K and V with a context.
type MapWithContext[K comparable, V any] interface {
	Each(f MapPredicateWithContext[K, V]) error
	Every(f MapPredicateWithContext[K, V]) (bool, error)
	Where(f MapPredicateWithContext[K, V]) (MapWithContext[K, V], error)
}

// MapJoiner is a type that represents a map of type K and V that can be joined with another map of type K and U.
type MapJoiner[U any, K comparable, V any] interface {
	Join(other Map[K, U]) (Map[K, pair.Type[V, U]], error)
	JoinMust(other Map[K, U]) Map[K, pair.Type[V, U]]
}

// MapTransformer is a type that represents a map of type K and V that can be transformed into a map of type K and U.
type MapTransformer[U any, K comparable, V any] interface {
	Transform(f TransformFunc[V, U]) (Map[K, U], error)
	TransformMust(f TransformFuncMust[V, U]) Map[K, U]
}

// MapWithContextTransformer is a type that represents a map of type K and V that can be transformed into a map of type K and U with a context.
type MapWithContextTransformer[U any, K comparable, V any] interface {
	Transform(f TransformFuncWithContext[V, U]) (Map[K, U], error)
}
