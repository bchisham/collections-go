package sequence

import (
	"fmt"
	"github.com/bchisham/collections-go/contracts"
	"strings"
)

type vector[T contracts.NumericType] []T

func FromNumericSlice[N contracts.NumericType](nseq []N) contracts.Vector[N] {
	return vector[N](nseq)
}

func (v vector[T]) ToSlice() []T {
	return v
}

func (v vector[T]) ToSequence() contracts.Sequence[T] {
	return Type[T](v)
}

func (v vector[T]) Length() int {
	return len(v)
}

func (v vector[T]) DotProduct(other contracts.Vector[T]) T {
	otherSeq := other.ToSequence().ToSlice()
	if v.Length() != len(otherSeq) {
		return 0
	}
	var result T
	for i, item := range v {
		result += item * otherSeq[i]
	}
	return result
}

func (v vector[T]) Add(other contracts.Vector[T]) contracts.Vector[T] {
	otherSeq := other.ToSequence().ToSlice()
	if v.Length() != len(otherSeq) {
		return nil
	}
	result := make(vector[T], v.Length())
	for i, item := range v {
		result[i] = item + otherSeq[i]
	}
	return result
}

func (v vector[T]) Subtract(other contracts.Vector[T]) contracts.Vector[T] {
	otherSeq := other.ToSequence().ToSlice()
	if v.Length() != len(otherSeq) {
		return nil
	}
	result := make(vector[T], v.Length())
	for i, item := range v {
		result[i] = item - otherSeq[i]
	}
	return result
}

func (v vector[T]) CrossProduct(other contracts.Vector[T]) contracts.Vector[T] {
	otherSeq := other.ToSequence().ToSlice()
	if v.Length() != len(otherSeq) {
		return nil
	}
	result := make(vector[T], v.Length())
	for i, item := range v {
		result[i] = item * otherSeq[i]
	}
	return result
}
func (v vector[T]) Scale(factor T) contracts.Vector[T] {
	result := make(vector[T], v.Length())
	for i, item := range v {
		result[i] = item * factor
	}
	return result
}

func (v vector[T]) String() string {
	buf := strings.Builder{}
	buf.WriteString("[")
	for i, item := range v {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(fmt.Sprintf("%v", item))
	}
	buf.WriteString("]")
	return buf.String()
}
