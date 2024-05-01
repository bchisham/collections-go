package association

import (
	"collections-go/contracts"
	"collections-go/sequence"
)

type Type[T comparable, U any] contracts.MapType[T, U]

func FromMap[T comparable, U any](m map[T]U) contracts.Map[T, U] {
	return Type[T, U](m)
}

func (m Type[T, U]) Each(f contracts.ApplyFunc[U]) error {
	for _, v := range m {
		if err := f(v); err != nil {
			return err
		}
	}
	return nil
}

func (m Type[T, U]) Every(f contracts.MapPredicate[T, U]) (bool, error) {
	for k, v := range m {
		if ok, err := f(k, v); err != nil {
			return false, err
		} else if !ok {
			return false, nil
		}
	}
	return true, nil
}

func (m Type[T, U]) EveryMust(f contracts.MapPredicateMust[T, U]) bool {
	for k, v := range m {
		if !f(k, v) {
			return false
		}
	}
	return true

}

func (m Type[T, U]) Where(f contracts.MapPredicate[T, U]) (contracts.Map[T, U], error) {
	result := make(map[T]U)
	for k, v := range m {
		if ok, err := f(k, v); err != nil {
			return nil, err
		} else if ok {
			result[k] = v
		}
	}
	return FromMap(result), nil
}

func (m Type[T, U]) WhereMust(f contracts.MapPredicateMust[T, U]) contracts.Map[T, U] {
	result := make(map[T]U)
	for k, v := range m {
		if f(k, v) {
			result[k] = v
		}
	}
	return FromMap(result)
}

func (m Type[T, U]) Keys() contracts.Sequence[T] {
	var keys sequence.Type[T]
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (m Type[T, U]) Values() contracts.Sequence[U] {
	var values sequence.Type[U]
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func (m Type[T, U]) ToMap() map[T]U {
	return m
}
