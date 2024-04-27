package sequence

import "collections-go/contracts"

func (seq Type[T]) EveryMust(f contracts.UnaryPredicateMust[T]) bool {
	for _, item := range seq {
		if !f(item) {
			return false
		}
	}
	return true
}
