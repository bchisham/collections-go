package sequence

func FirstWhere[T any](list []T, predicate PredicateFunc[T]) (result T, found bool) {
	for _, item := range list {
		if predicate(item) {
			return item, true
		}
	}
	return result, false
}
