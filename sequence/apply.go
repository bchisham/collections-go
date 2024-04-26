package sequence

type ApplyFunc[T any] func(item T) error

func Apply[T any](list []T, f ApplyFunc[T]) error {
	for _, item := range list {
		if err := f(item); err != nil {
			return err
		}
	}
	return nil
}

func (s Type[T]) Apply(f ApplyFunc[T]) error {
	return Apply(s, f)
}
