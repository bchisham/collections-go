package sequence

type Type[T any] []T

type IntegerType interface {
	int | int64 | int32 | int16 | int8 | uint | uint64 | uint32 | uint16 | uint8
}

type FloatType interface {
	float32 | float64
}

type NumericType interface {
	IntegerType | FloatType
}

// PredicateFunc is a function that takes a value of type T and returns a boolean value.
type PredicateFunc[T any] func(T) bool

type Unwrapable[T any] interface {
	Unwrap() []T
}

type Interface[T any] interface {
	Apply(f ApplyFunc[T]) error
	Where(f PredicateFunc[T]) Type[T]
	FirstWhere(list []T, predicate PredicateFunc[T]) (result T, found bool)
}

type Generator[T any, U any] interface {
	Yield(value U) Generator[T, U]
}

// New creates a new sequence of type T
func New[T any](list []T) Type[T] {
	return list
}

// Transformer is an interface that can be used to transform a sequence of type T to a sequence of type U
type Transformer[T any, U any] interface {
	Transform(f TransformFunc[T, U]) Type[U]
}

type NumericAggregator[T any, N NumericType] interface {
	Sum(f func(T) N) N
	Mean(f func(T) N) N
}

func (seq Type[T]) Unwrap() []T {
	return seq
}
