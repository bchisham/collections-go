package contracts

type Channel[T any] interface {
	ToChannel[T]
	Send(T) error
	Receive() (T, error)
}

type ChannelWithClose[T any] interface {
	Channel[T]
	Close() error
}
