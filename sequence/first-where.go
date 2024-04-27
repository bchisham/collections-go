package sequence

import (
	"collections-go/contracts"
)

func FirstWhere[T any](list []T, predicate contracts.UnaryPredicateMust[T]) (result T, found bool) {
	for _, item := range list {
		if predicate(item) {
			return item, true
		}
	}
	return result, false
}

func (seq Type[T]) FirstWhereMust(predicate contracts.UnaryPredicateMust[T]) (result T, found bool) {
	return FirstWhere(seq, predicate)
}

func (seq ContextualSequence[T]) FirstWhere(predicate contracts.ContextPredicate[T]) (result T, found bool, err error) {
	for _, item := range seq.Sequence {
		ok, err := predicate(seq.ctx, item)
		if err != nil {
			return item, false, err
		}
		if ok {
			return item, true, nil
		}
	}
	return result, false, nil
}
