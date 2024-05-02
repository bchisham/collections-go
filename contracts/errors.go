package contracts

type ErrorType struct {
	Code    int
	Message string
}

func (e ErrorType) Error() string {
	return e.Message
}

var ErrAssertionFailed = ErrorType{Code: 1, Message: "Assertion failed"}
var ErrContextDone = ErrorType{Code: 2, Message: "Context done"}
