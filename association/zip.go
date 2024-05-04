package association

import "collections-go/contracts"

// Zip takes two sequences and returns a map of the two sequences zipped together.
func Zip[T comparable, U any](seq1 contracts.Sequence[T], seq2 contracts.Sequence[U]) contracts.Map[T, U] {
	if seq1.Length() != seq2.Length() {
		return nil
	}
	m := make(Type[T, U])
	vals := seq2.ToSlice()
	for i, v := range seq1.ToSlice() {
		m[v] = vals[i]
	}
	return m
}
