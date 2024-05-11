package contracts

type Channel[T any] interface {
	ToChannel[T]
	Send(T) error
	Receive() (T, error)
	Each(f ApplyFunc[T]) error
	Close() error
}

type ReadChannel[T any] interface {
	ToChannel() <-chan T
	Receive() (T, error)
	Each(f ApplyFunc[T]) error
}

type WriteChannel[T any] interface {
	ToChannel() chan<- T
	Send(T) error
	Close() error
}
