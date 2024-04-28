package contracts

import "github.com/google/uuid"

type IntegerType interface {
	int | int64 | int32 | int16 | int8 | uint | uint64 | uint32 | uint16 | uint8
}

type FloatType interface {
	float32 | float64
}

type NumericType interface {
	IntegerType | FloatType
}

type IdentityType interface {
	IntegerType | uuid.UUID | string
}

type SequenceType[T any] []T

type MapType[K comparable, V any] map[K]V

type ToSlice[T any] interface {
	ToSlice() []T
}

type ToMap[K comparable, V any] interface {
	ToMap() map[K]V
}
