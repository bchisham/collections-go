package sequence

import (
	"errors"
	"fmt"
)

type Error struct {
	Message string
	Index   int
	Length  int
	inner   error
}

func (e Error) Error() string {
	return fmt.Sprintf("Error: %s, index: %d, length: %d", e.Message, e.Index, e.Length)
}

func Unwrap(err error) error {
	var e Error
	if errors.As(err, &e) {
		return e.inner
	}
	return nil
}

func Wrap(err error, message string, index int, length int) error {
	if err == nil {
		return nil
	}
	return Error{message, index, length, err}

}

func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return Error{fmt.Sprintf(format, args...), -1, -1, err}

}

func (e Error) WithMessage(message string) error {
	e.Message = message
	return e
}

func (e Error) WithIndex(index int) error {
	e.Index = index
	return e
}

func (e Error) WithLength(length int) error {
	e.Length = length
	return e

}

var ErrEndOfSequence = Error{Message: "End of sequence"}
