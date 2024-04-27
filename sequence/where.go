package sequence

import "collections-go/contracts"

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
func WhereMust[T any](seq []T, f contracts.UnaryPredicateMust[T]) contracts.Sequence[T] {
	return FromSlicef(seq).WhereMust(f)
}

func (seq ContextualSequence[T]) Where(f contracts.ContextPredicate[T]) (contracts.Sequence[T], error) {
	var r Type[T]
	for _, v := range seq.Sequence {
		ok, err := f(seq.ctx, v)
		if err != nil {
			return nil, err
		}
		if ok {
			r = append(r, v)
		}
	}
	return r, nil
}
