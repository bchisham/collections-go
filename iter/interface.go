package iter

import (
	"iter"
)

// MapToSeq applies a transformation function to each element of the input slice and returns a sequence of the results.
func MapToSeq[U any, ISlice ~[]T, T any](s ISlice) func(xf func(T) U) iter.Seq[U] {
	return func(xf func(T) U) iter.Seq[U] {
		return func(yield func(U) bool) {
			for i := range s {
				if !yield(xf(s[i])) {
					return
				}
			}
		}
	}
}
