package serror

import "net/http"

//go:generate go-enum -f=$GOFILE --marshal

/*
ENUM(Undefined, InvalidParam, NotFound, ImportDeleted, Unauthorized, Forbidden, InvalidCategoryType, PayAgentError, DuplicateCard, UnsupportedMedia, DuplicateReport, Expired)
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
	case CodeUnauthorized:
		return http.StatusUnauthorized
	case CodeForbidden:
		return http.StatusForbidden
	case CodePayAgentError:
		return http.StatusPaymentRequired
	case CodeDuplicateReport:
		return http.StatusBadRequest
	case CodeUnsupportedMedia:
		return http.StatusBadRequest
	case CodeExpired:
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}
