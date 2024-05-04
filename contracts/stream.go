package contracts

type Channel[T any] interface {
	ToChannel[T]
	Send(T) error
	Receive() (T, error)
	Each(f ApplyFunc[T]) error
	Close() error
}
