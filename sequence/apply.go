package sequence

import "collections-go/contracts"

func Each[T any](list []T, f contracts.ApplyFunc[T]) error {
	for _, item := range list {
		if err := f(item); err != nil {
			return err
		}
	}
	return nil
}

func (seq Type[T]) Each(f contracts.ApplyFunc[T]) error {
	return Each(seq, f)
}

func (seq ContextualSequence[T]) Each(f contracts.ApplyWithContextFunc[T]) error {
	for _, item := range seq.Sequence {
		if err := f(seq.ctx, item); err != nil {
			return err
		}
	}
	return nil
}
