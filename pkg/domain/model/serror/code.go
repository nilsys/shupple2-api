package serror

import "net/http"

//go:generate go-enum -f=$GOFILE --marshal

/*
ENUM(Undefined, InvalidParam, NotFound, Unauthorized, Forbidden, NotMatching, MatchingNotExpired)
*/
type Code int

func (x Code) HTTPStatusCode() int {
	switch x {
	case CodeInvalidParam:
		return http.StatusBadRequest
	case CodeNotFound:
		return http.StatusNotFound
	case CodeUnauthorized:
		return http.StatusUnauthorized
	case CodeForbidden:
		return http.StatusForbidden
	case CodeNotMatching:
		return http.StatusBadRequest
	case CodeMatchingNotExpired:
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}
