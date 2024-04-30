package association

import (
	"collections-go/contracts"
	"context"
)

type mapTransform[K comparable, V1 any, V2 any] struct {
	Type[K, V1]
}

type contextualMapTransform[K comparable, V1 any, V2 any] struct {
	mapTransform[K, V1, V2]
	ctx context.Context
}

func NewMapTransform[K comparable, V1 any, V2 any](m contracts.MapType[K, V1]) contracts.MapTransformer[K, V1, V2] {
	return mapTransform[K, V1, V2]{Type: Type[K, V1](m)}
}

func NewContextualMapTransform[K comparable, V1 any, V2 any](m contracts.MapType[K, V1], ctx context.Context) contracts.MapWithContextTransformer[K, V1, V2] {
	return contextualMapTransform[K, V1, V2]{mapTransform: mapTransform[K, V1, V2]{Type: Type[K, V1](m)}, ctx: ctx}
}

func (m mapTransform[K, V1, V2]) Transform(f contracts.MapTransformFunc[K, V1, V2]) (contracts.Map[K, V2], error) {
	result := make(map[K]V2)
	for k, v := range m.Type {
		if transformed, err := f(k, v); err != nil {
			return nil, err
		} else {
			result[k] = transformed
		}
	}
	return FromMap[K, V2](result), nil
}

func (m mapTransform[K, V1, V2]) TransformMust(f contracts.MapTransformFuncMust[K, V1, V2]) contracts.Map[K, V2] {
	result := make(map[K]V2)
	for k, v := range m.Type {
		result[k] = f(k, v)
	}
	return FromMap[K, V2](result)
}

func (m contextualMapTransform[K, V1, V2]) Transform(f contracts.MapTransformFuncWithContext[K, V1, V2]) (contracts.Map[K, V2], error) {
	result := make(map[K]V2)
	var err error
	for k, v := range m.Type {
		result[k], err = f(m.ctx, k, v)
		if err != nil {
			return nil, err
		}
	}
	return FromMap[K, V2](result), nil
}
