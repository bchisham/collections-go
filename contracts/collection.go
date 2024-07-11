package contracts

import (
	"context"
	"github.com/bchisham/collections-go/pair"
)

// ApplyFunc is a function that takes a value of type T and returns an error.
type ApplyFunc[T any] func(item T) error

// ApplyWithContextFunc is a function that takes a context and a value of type T and returns an error.
type ApplyWithContextFunc[T any] func(ctx context.Context, item T) error

// Sequence is a type that represents a sequence of type T
type Sequence[T any] interface {
	ToSlice[T]
	// Each applies a function to each item in the sequence and returns an error if one occurs
	Each[T]
	// Where filters the sequence based on a predicate
	Where[T]
	// WhereMust filters the sequence based on a predicate
	WhereMust[T]

	Every[T]
	// EveryMust returns true if every item in the sequence satisfies the predicate.
	EveryMust[T]
	// FirstWhereMust returns the first item in the sequence that satisfies the predicate.
	FirstWhereMust(predicate UnaryPredicateMust[T]) (result T, found bool)
	First() (T, error)
	Last() (T, error)
	Length() int
}

// ContextualSequence is a type that represents a sequence of type T with a context.
type ContextualSequence[T any] interface {
	ToSlice[T]
	ToSequence[T]
	Each(f ApplyWithContextFunc[T]) error
	Every(f ContextPredicate[T]) (bool, error)
	Where(f ContextPredicate[T]) (Sequence[T], error)
	FirstWhere(predicate ContextPredicate[T]) (T, bool, error)
	Length() int
}

type Map[K comparable, V any] interface {
	ToMap() map[K]V
	Each[V]
	Every[V]
	EveryMust[V]
	Where(f UnaryPredicate[V]) (Map[K, V], error)
	WhereMust(f UnaryPredicateMust[V]) Map[K, V]
	Keys() Sequence[K]
	Values() Sequence[V]
}

// MapWithContext is a type that represents a map of type K and V with a context.
type MapWithContext[K comparable, V any] interface {
	Each(f ApplyWithContextFunc[V]) error
	Every(f ContextPredicate[V]) (bool, error)
	Where(f ContextPredicate[V]) (MapWithContext[K, V], error)
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
