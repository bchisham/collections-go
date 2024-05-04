package sequence

import "github.com/bchisham/collections-go/contracts"

// PartitionMust takes a sequence of type T and a function that takes a value of type T and partitions the set into two sets that satisfy the predicate and the complement of the predicate.
func PartitionMust[T any](seq contracts.Sequence[T], f contracts.UnaryPredicateMust[T]) (contracts.Sequence[T], contracts.Sequence[T]) {
	left := seq.WhereMust(f)
	right := seq.WhereMust(func(v T) bool { return !f(v) })
	return left, right
}
