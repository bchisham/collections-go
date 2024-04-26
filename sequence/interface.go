package sequence

type Type[T any] []T

type Interface[T any] interface {
	Apply(f ApplyFunc[T]) error
	Filter(f FilterFunc[T]) Type[T]
}

// New creates a new sequence of type T
func New[T any](list []T) Type[T] {
	return list
}

// Transformer is an interface that can be used to transform a sequence of type T to a sequence of type U
type Transformer[T any, U any] interface {
	Transform(f TransformFunc[T, U]) Type[U]
}
