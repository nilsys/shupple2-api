package api

import (
	"fmt"
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	validatorv9 "gopkg.in/go-playground/validator.v9"
)

const stackTraceDepth = 5

type (
	ErrorResponse struct {
		Status string `json:"status"`
		Error  string `json:"error"`
		Body   string `json:"body"`
	}

	stackTracer interface {
		StackTrace() errors.StackTrace
	}
)

func ErrorHandler(err error, ctx echo.Context) {
	req := ctx.Request()
	code := GetStatusCode(err)

	if code/100 == 5 {
		logger.Error(err.Error(),
			zap.String("url", req.URL.String()),
			zap.Int("status", code),
			zap.String("stacktrace", getStackTrace(err)),
		)
	}

	// TODO: メッセージどうするか
	resp := ErrorResponse{
		Status: http.StatusText(code),
		Error:  GetErrorString(err),
		Body:   err.Error(),
	}

	if err := ctx.JSON(code, &resp); err != nil {
		logger.Error(err.Error(), zap.Error(err))
	}
}

func GetErrorString(err error) string {
	if serr, ok := err.(*serror.SError); ok {
		return serr.Code.String()
	}

	return serror.CodeUndefined.String()
}

func GetStatusCode(err error) int {
	switch err := errors.Cause(err).(type) {
	case *serror.SError:
		return err.Code.HTTPStatusCode()
	case *echo.HTTPError:
		return err.Code
	case validatorv9.ValidationErrors:
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}

func getStackTrace(err error) string {
	err = errors.Cause(err)
	if e, ok := err.(stackTracer); ok {
		st := e.StackTrace()
		if len(st) > stackTraceDepth {
			st = st[:stackTraceDepth]
		}
		return fmt.Sprintf("%+v", st)
	}
	return ""
}
