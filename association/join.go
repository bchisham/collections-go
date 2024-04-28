package association

import (
	"collections-go/contracts"
	"collections-go/pair"
)

type Joiner[T comparable, V any, U any] struct {
	Type[T, V]
}

func (m Joiner[T, V, U]) JoinMust(other contracts.MapType[T, U]) contracts.Map[T, pair.Type[V, U]] {
	result := make(map[T]pair.Type[V, U])
	for k, v := range m.Type {
		if otherV, ok := other[k]; ok {
			result[k] = pair.New(v, otherV)
		}
	}
	return FromMap(result)
}

func (m Joiner[T, V, U]) Join(other contracts.MapType[T, U]) (contracts.Map[T, pair.Type[V, U]], error) {
	result := make(map[T]pair.Type[V, U])
	for k, v := range m.Type {
		if otherV, ok := other[k]; ok {
			result[k] = pair.New(v, otherV)
		}
	}
	return FromMap(result), nil
}
