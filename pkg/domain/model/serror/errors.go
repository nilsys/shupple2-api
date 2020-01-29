package serror

import (
	"fmt"

	"github.com/pkg/errors"
)

type (
	SError struct {
		Code    Code
		Message string
		Cause   error
	}
)

func New(cause error, code Code, message string, v ...interface{}) *SError {
	return &SError{
		Code:    code,
		Message: fmt.Sprintf(message, v...),
		Cause:   cause,
	}
}

func (e *SError) Error() string {
	if e.Cause == nil {
		return e.Message
	}
	return fmt.Sprintf("%s: %s", e.Message, e.Cause.Error())
}

func IsErrorCode(err error, codes ...Code) bool {
	if err == nil {
		return false
	}

	err = errors.Cause(err)
	if serr, ok := err.(*SError); ok {
		for _, code := range codes {
			if serr.Code == code {
				return true
			}
		}
	}

	return false
}
