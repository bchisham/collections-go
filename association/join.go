package association

import (
	"collections-go/contracts"
	"collections-go/pair"
)

type Joiner[T comparable, V any, U any] struct {
	Type[T, V]
}

func NewJoiner[T comparable, V any, U any](m contracts.Map[T, V]) contracts.MapJoiner[T, V, U] {
	return Joiner[T, V, U]{Type: Type[T, V](m.ToMap())}
}

func (m Joiner[T, V, U]) JoinMust(other contracts.Map[T, U]) contracts.Map[T, pair.Type[V, U]] {
	result := make(map[T]pair.Type[V, U])
	otherMap := other.ToMap()
	for k, v := range m.Type {
		if otherV, ok := otherMap[k]; ok {
			result[k] = pair.New(v, otherV)
		}
	}
	return FromMap(result)
}

func (m Joiner[T, V, U]) Join(other contracts.Map[T, U]) (contracts.Map[T, pair.Type[V, U]], error) {
	result := make(map[T]pair.Type[V, U])
	otherMap := other.ToMap()
	for k, v := range m.Type {
		if otherV, ok := otherMap[k]; ok {
			result[k] = pair.New(v, otherV)
		}
	}
	return FromMap(result), nil
}
