package contracts

import "context"

// ApplyFunc is a function that takes a value of type T and returns an error.
type ApplyFunc[T any] func(item T) error

// ApplyWithContextFunc is a function that takes a context and a value of type T and returns an error.
type ApplyWithContextFunc[T any] func(ctx context.Context, item T) error

// Sequence is a type that represents a sequence of type T
type Sequence[T any] interface {
	// ToSlice converts the sequence to a slice of type T
	ToSlice() []T
	// Each applies a function to each item in the sequence and returns an error if one occurs
	Each(f ApplyFunc[T]) error
	// Where filters the sequence based on a predicate
	Where(f UnaryPredicate[T]) (Sequence[T], error)
	// WhereMust filters the sequence based on a predicate
	WhereMust(f UnaryPredicateMust[T]) Sequence[T]
	// FirstWhereMust returns the first item in the sequence that satisfies the predicate.
	FirstWhereMust(predicate UnaryPredicateMust[T]) (result T, found bool)
	// WithContext returns a ContextualSequence with the given context
	EveryMust(f UnaryPredicateMust[T]) bool
}

// ContextualSequence is a type that represents a sequence of type T with a context.
type ContextualSequence[T any] interface {
	ToSlice() []T
	Each(f ApplyWithContextFunc[T]) error
	Where(f ContextPredicate[T]) (Sequence[T], error)
	FirstWhere(predicate ContextPredicate[T]) (T, bool, error)
}
