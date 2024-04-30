package contracts

import "context"

type MapTransformFunc[K comparable, V any, U any] func(K, V) (U, error)
type MapTransformFuncMust[K comparable, V any, U any] func(K, V) U
type MapTransformFuncWithContext[K comparable, V any, U any] func(context.Context, K, V) (U, error)
