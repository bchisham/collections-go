package iter

import "github.com/bchisham/collections-go/contracts"

// Filter applies a filter function to each element of the input slice and
// returns a new slice containing only the elements that satisfy the filter
// condition.
func Filter[ISlice ~[]T, T any](s ISlice) func(filter contracts.FilterFunc[T]) ISlice {
	return func(filter contracts.FilterFunc[T]) ISlice {
		var u = make(ISlice, 0, len(s))
		for _, v := range s {
			if filter(v) {
				u = append(u, v)
			}
		}
		return u
	}
}
