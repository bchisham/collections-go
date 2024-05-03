package contracts

type Channel[T any] interface {
	ToChannel[T]
	Send(T) error
	Receive() (T, error)
	Close() error
}
