package solution31

import (
	"errors"
)

type nonCriticalError struct{}

func (e nonCriticalError) Error() string {
	return "validation error"
}

var (
	errBadConnection = errors.New("bad connection")
	errBadRequest    = errors.New("bad request")
)

const unknownErrorMsg = "unknown error"

func GetErrorMsg(err error) string {
	if errors.As(err, &nonCriticalError{}) {
		return ""
	} else if errors.Is(err, errBadConnection) {
		return errBadConnection.Error()
	} else if errors.Is(err, errBadRequest) {
		return errBadRequest.Error()
	}
	return (unknownErrorMsg)
}
