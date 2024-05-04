package sequence

import "github.com/bchisham/collections-go/contracts"

// Chunk takes a sequence and a chunk size and returns a slice of sequences of the given chunk size.
func Chunk[T any](sequenceType contracts.Sequence[T], chunkSize int) contracts.Sequence[contracts.Sequence[T]] {
	if chunkSize <= 0 {
		return nil
	}
	var chunks []contracts.Sequence[T]
	seqSlice := sequenceType.ToSlice()
	for i := 0; i < sequenceType.Length(); i += chunkSize {
		end := i + chunkSize
		if end > sequenceType.Length() {
			end = sequenceType.Length()
		}
		chunks = append(chunks, FromSlice(seqSlice[i:end]))
	}
	return FromSlice(chunks)
}
