package sequence

import (
	"context"
)

func (seq ContextualSequence[T]) ToSlice() []T {
	return seq.Sequence
}

func (seq Type[T]) ToSlice() []T {
	return seq
}

func (seq ContextualSequence[T]) continueWithContext(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}
