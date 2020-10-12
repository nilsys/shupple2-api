package api

import (
	"fmt"
	"net/http"
	"syscall"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	validatorv9 "gopkg.in/go-playground/validator.v9"
)

const stackTraceDepth = 5

type (
	ErrorResponse struct {
		RequestID string `json:"request_id"`
		Status    string `json:"status"`
		Error     string `json:"error"`
		Body      string `json:"body"`
	}

	stackTracer interface {
		StackTrace() errors.StackTrace
	}

	errorHandler struct {
		env config.Env
	}
)

func CreateErrorHandler(env config.Env) echo.HTTPErrorHandler {
	return (errorHandler{env: env}).handle
}

func (h errorHandler) handle(err error, ctx echo.Context) {
	req := ctx.Request()
	code := GetStatusCode(err)

	requestID := ctx.Response().Header().Get(echo.HeaderXRequestID) // NOTE: middlewareの都合でresから取らないといけない

	if code/100 == 5 {
		logger.Error(err.Error(),
			zap.String("request_id", requestID),
			zap.String("url", req.URL.String()),
			zap.Int("status", code),
			zap.String("stacktrace", getStackTrace(err)),
		)
	}

	errResp := ErrorResponse{
		RequestID: requestID,
		Status:    http.StatusText(code),
		Error:     GetErrorString(err),
	}
	if !h.env.IsPrd() {
		errResp.Body = err.Error()
	}

	if err := ctx.JSON(code, &errResp); err != nil {
		if isBrokenPipe(err) {
			// broken pipeはどうしようもない上にそこそこ発生するのでINFOに
			logger.Info(err.Error(), zap.Error(err))
		} else {
			logger.Error(err.Error(), zap.Error(err))
		}
	}
}

func GetErrorString(err error) string {
	if serr, ok := errors.Cause(err).(*serror.SError); ok {
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

func isBrokenPipe(err error) bool {
	return errors.Is(err, syscall.EPIPE) || errors.Is(err, syscall.ECONNRESET)
}
