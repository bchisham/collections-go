package contracts

import "context"

type TransformFunc[T any, U any] func(T) (U, error)
type TransformFuncMust[T any, U any] func(T) U
type TransformFuncWithContext[T any, U any] func(context.Context, T) (U, error)

type MapTransformFunc[U any, K comparable, V any] func(V) (U, error)
type MapTransformFuncMust[U any, K comparable, V any] func(V) U
type MapTransformFuncWithContext[U any, K comparable, V any] func(context.Context, V) (U, error)
