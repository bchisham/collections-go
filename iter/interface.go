package iter

import (
	"iter"
)

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
