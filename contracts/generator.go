package contracts

// Generator is an interface that can be used to generate a sequence of type T
type Generator[T any] interface {
	Yield() error
	ToSlice() Sequence[T]
}
