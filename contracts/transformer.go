package contracts

import "context"

type MapTransformFunc[K comparable, V any, U any] func(V) (U, error)
type MapTransformFuncMust[K comparable, V any, U any] func(V) U
type MapTransformFuncWithContext[K comparable, V any, U any] func(context.Context, V) (U, error)
