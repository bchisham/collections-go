package sequence

type Type[T any] []T

type Interface[T any] interface {
	Filter(f FilterFunc[T]) Type[T]
}

func New[T any](list []T) Type[T] {
	return list
}

type Transformer[T any, U any] interface {
	Transform(f TransformFunc[T, U]) Type[U]
}
