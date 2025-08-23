package iter

import "github.com/bchisham/collections-go/contracts"

func Each[ISlice ~[]T, T any](s ISlice) func(f contracts.ApplyFunc[T]) (ISlice, error) {
	return func(f contracts.ApplyFunc[T]) (ISlice, error) {
		for _, v := range s {
			if err := f(v); err != nil {
				return nil, err
			}
		}
		return s, nil
	}
}

// EachMust applies a function to each element of the input slice and returns the original slice.
func EachMust[ISlice ~[]T, T any](s ISlice) func(f contracts.ApplyFuncMust[T]) ISlice {
	return func(f contracts.ApplyFuncMust[T]) ISlice {
		for _, v := range s {
			f(v)
		}
		return s
	}
}
