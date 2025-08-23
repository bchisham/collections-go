package iter

import "github.com/bchisham/collections-go/contracts"

// MapAppendMust applies a transformation function to each element of the input slice and appends the results to the output slice.
func MapAppendMust[OSlice ~[]U, U any, ISlice ~[]T, T any](s ISlice, u OSlice) func(xform contracts.TransformFuncMust[T, U]) OSlice {
	return func(xform contracts.TransformFuncMust[T, U]) OSlice {
		for _, v := range s {
			u = append(u, xform(v))
		}
		return u
	}
}

func Map[OSlice ~[]U, U any, ISlice ~[]T, T any](s ISlice) func(xf contracts.TransformFunc[T, U]) (OSlice, error) {
	return func(xf contracts.TransformFunc[T, U]) (OSlice, error) {
		var u = make(OSlice, 0, len(s))
		for _, v := range s {
			if v, err := xf(v); err != nil {
				return nil, err
			} else {
				u = append(u, v)
			}

		}
		return u, nil
	}
}

func MapMust[OSlice ~[]U, U any, ISlice ~[]T, T any](s ISlice) func(xf contracts.TransformFuncMust[T, U]) OSlice {
	return func(xf contracts.TransformFuncMust[T, U]) OSlice {
		var u = make(OSlice, 0, len(s))
		for _, v := range s {
			u = append(u, xf(v))
		}
		return u
	}
}
