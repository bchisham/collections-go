package contracts

// Generator is an interface that can be used to generate a sequence of type T
type Generator[T any, U any] interface {
	Yield(value U) Generator[T, U]
}
