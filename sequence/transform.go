package sequence

//Transform takes a sequence of type T and a function that takes a value of type T and returns a value of type U. It returns a sequence of type U

// TransformFunc is a function that takes a value of type T and returns a value of type U
type TransformFunc[T any, U any] func(T) U

type transformer[T any, U any] struct {
	Type[T]
}

func NewTransformer[T any, U any](s Type[T]) Transformer[T, U] {
	return transformer[T, U]{s}
}

// Transform takes a sequence of type T and a function that takes a value of type T and returns a value of type U. It returns a sequence of type U
func Transform[T any, U any](s Type[T], f TransformFunc[T, U]) Type[U] {
	r := make(Type[U], len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// Transform takes a function that takes a value of type T and returns a value of type U. It returns a Transformer that can be used to transform a sequence of type T to a sequence of type U
func (t transformer[T, U]) Transform(f TransformFunc[T, U]) Type[U] {
	return Transform(t.Type, f)
}
