package serror

import "net/http"

//go:generate go-enum -f=code.go --marshal

/*
ENUM(Undefined, InvalidParam, NotFound)
*/
type Code int

func (c Code) HTTPStatusCode() int {
	switch c {
	case CodeInvalidParam:
		return http.StatusBadRequest
	case CodeNotFound:
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}
