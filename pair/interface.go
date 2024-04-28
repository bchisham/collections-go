package pair

type Type[T any, U any] struct {
	first  T
	second U
}

func New[T any, U any](first T, second U) Type[T, U] {
	return Type[T, U]{first: first, second: second}
}

func (p Type[T, U]) First() T {
	return p.first
}

func (p Type[T, U]) Second() U {
	return p.second
}
