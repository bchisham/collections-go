package sequence

import "github.com/bchisham/collections-go/contracts"

func (seq Type[T]) Assert(f contracts.UnaryPredicateMust[T]) error {
	for _, item := range seq {
		if !f(item) {
			return ErrAssertionFailed
		}
	}
	return nil
}

func (seq Type[T]) First() (f T, _ error) {
	if len(seq) == 0 {
		return f, ErrEmptySequence
	}
	return seq[0], nil
}

func (seq Type[T]) At(index int) (f T, _ error) {
	if index < 0 || index >= len(seq) {
		return f, ErrIndexOutOfBounds
	}
	return seq[index], nil
}

func (seq Type[T]) Last() (f T, _ error) {
	if len(seq) == 0 {
		return f, ErrEmptySequence
	}
	return seq[len(seq)-1], nil
}

func (seq Type[T]) Push(v T) contracts.Sequence[T] {
	return append(seq, v)
}

func (seq Type[T]) TrimTo(newLen int) (contracts.Sequence[T], error) {
	if newLen < 0 || newLen > len(seq) {
		return nil, ErrInvalidLength
	}
	return seq[:newLen], nil
}

func (seq Type[T]) Subsequence(start, end int) (contracts.Sequence[T], error) {
	if start < 0 || end > len(seq) || start > end {
		return nil, ErrInvalidRange
	}
	return seq[start:end], nil
}

func (seq Type[T]) Each(f contracts.ApplyFunc[T]) error {
	for _, item := range seq {
		if err := f(item); err != nil {
			return err
		}
	}
	return nil
}

func (seq Type[T]) EachMust(f contracts.ApplyFuncMust[T]) {
	for _, item := range seq {
		f(item)
	}
}

func (seq Type[T]) EveryMust(f contracts.UnaryPredicateMust[T]) bool {
	for _, item := range seq {
		if !f(item) {
			return false
		}
	}
	return true
}

func (seq Type[T]) Every(f contracts.UnaryPredicate[T]) (bool, error) {
	for _, item := range seq {
		ok, err := f(item)
		if err != nil {
			return false, err
		}
		if !ok {
			return false, nil
		}
	}
	return true, nil

}

func (seq Type[T]) FirstWhereMust(predicate contracts.UnaryPredicateMust[T]) (result T, found bool) {
	for _, item := range seq {
		if predicate(item) {
			return item, true
		}
	}
	return result, false
}

func (seq Type[T]) Mutate(index int, f contracts.ApplyAtFunc[T]) error {
	if index < 0 || index >= len(seq) {
		return ErrIndexOutOfBounds
	}
	return f(index, seq[index])
}

func (seq Type[T]) Where(f contracts.UnaryPredicate[T]) (contracts.Sequence[T], error) {
	var r Type[T]
	for _, v := range seq {
		ok, err := f(v)
		if err != nil {
			return nil, err
		}
		if ok {
			r = append(r, v)
		}
	}
	return r, nil
}

// WhereMust takes a sequence of type T and a function that takes a value of type T
// and returns a boolean value. It returns a sequence of type T with values that
// satisfy the function
func (seq Type[T]) WhereMust(f contracts.UnaryPredicateMust[T]) contracts.Sequence[T] {
	var r Type[T]
	for _, v := range seq {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func (seq Type[T]) ToSlice() []T {
	return seq
}
