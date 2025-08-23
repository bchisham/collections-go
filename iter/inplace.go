package iter

import "github.com/bchisham/collections-go/contracts"

// Inplace applies a mutation function to each element of the input slice and modifies the original slice in place.
func Inplace[ISlice ~[]T, T any](s ISlice) func(mutate contracts.MutateFunc[T]) ISlice {
	return func(mutate contracts.MutateFunc[T]) ISlice {
		if s == nil {
			return nil
		}
		for i, v := range s {
			s[i] = mutate(v)
		}
		return s
	}
}
