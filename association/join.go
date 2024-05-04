package association

import (
	"github.com/bchisham/collections-go/contracts"
	"github.com/bchisham/collections-go/pair"
)

type Joiner[U any, T comparable, V any] struct {
	Type[T, V]
}

func NewJoiner[U any, T comparable, V any](m contracts.Map[T, V]) contracts.MapJoiner[U, T, V] {
	return Joiner[U, T, V]{Type: Type[T, V](m.ToMap())}
}

func (m Joiner[U, T, V]) JoinMust(other contracts.Map[T, U]) contracts.Map[T, pair.Type[V, U]] {
	result := make(map[T]pair.Type[V, U])
	otherMap := other.ToMap()
	for k, v := range m.Type {
		if otherV, ok := otherMap[k]; ok {
			result[k] = pair.New(v, otherV)
		}
	}
	return FromMap(result)
}

func (m Joiner[U, T, V]) Join(other contracts.Map[T, U]) (contracts.Map[T, pair.Type[V, U]], error) {
	result := make(map[T]pair.Type[V, U])
	otherMap := other.ToMap()
	for k, v := range m.Type {
		if otherV, ok := otherMap[k]; ok {
			result[k] = pair.New(v, otherV)
		}
	}
	return FromMap(result), nil
}
