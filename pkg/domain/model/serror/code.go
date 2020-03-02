package serror

import "net/http"

//go:generate go-enum -f=$GOFILE --marshal

/*
ENUM(Undefined, InvalidParam, NotFound, ImportDeleted)
*/
type Code int

func (c Code) HTTPStatusCode() int {
	switch c {
	case CodeInvalidParam:
		return http.StatusBadRequest
	case CodeNotFound:
		return http.StatusNotFound
	case CodeImportDeleted:
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}
