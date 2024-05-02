package contracts

// Generator is an interface that can be used to generate a sequence of type T
type Generator[T any, U any] interface {
	Yield(value U) Generator[T, U]
}

type AssociationGenerator[U any, K comparable, V any] interface {
	Yield(value V) (AssociationGenerator[U, K, V], error)
	ToSlice() Sequence[U]
}

type IteratorToSlice[T any] interface {
	Yield() (IteratorToSlice[T], error)
	ToSlice() Sequence[T]
}
