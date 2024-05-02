package association

import (
	"collections-go/contracts"
	"context"
)

type mapTransform[U any, K comparable, V1 any] struct {
	Type[K, V1]
}

type contextualMapTransform[U any, K comparable, V1 any] struct {
	mapTransform[U, K, V1]
	ctx context.Context
}

func NewMapTransform[U any, K comparable, V1 any](m contracts.Map[K, V1]) contracts.MapTransformer[U, K, V1] {
	return mapTransform[U, K, V1]{Type: Type[K, V1](m.ToMap())}
}

func NewContextualMapTransform[U any, K comparable, V1 any](m contracts.MapType[K, V1], ctx context.Context) contracts.MapWithContextTransformer[U, K, V1] {
	return contextualMapTransform[U, K, V1]{mapTransform: mapTransform[U, K, V1]{Type: Type[K, V1](m)}, ctx: ctx}
}

func (m mapTransform[V2, K, V1]) Transform(f contracts.TransformFunc[V1, V2]) (contracts.Map[K, V2], error) {
	result := make(map[K]V2)
	for k, v := range m.Type {
		if transformed, err := f(v); err != nil {
			return nil, err
		} else {
			result[k] = transformed
		}
	}
	return FromMap[K, V2](result), nil
}

func (m mapTransform[V1, K, V2]) TransformMust(f contracts.TransformFuncMust[V2, V1]) contracts.Map[K, V1] {
	result := make(map[K]V1)
	for k, v := range m.Type {
		result[k] = f(v)
	}
	return FromMap(result)
}

func (m contextualMapTransform[V2, K, V1]) Transform(f contracts.TransformFuncWithContext[V1, V2]) (contracts.Map[K, V2], error) {
	result := make(map[K]V2)
	var err error
	for k, v := range m.Type {
		result[k], err = f(m.ctx, v)
		if err != nil {
			return nil, err
		}
	}
	return FromMap(result), nil
}
