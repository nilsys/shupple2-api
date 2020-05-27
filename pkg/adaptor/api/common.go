package api

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	validatorv9 "gopkg.in/go-playground/validator.v9"
)

var validator = validatorv9.New()

type Validator interface {
	Validate() error
}

func BindAndValidate(ctx echo.Context, v interface{}) error {
	if err := ctx.Bind(v); err != nil {
		return errors.Wrapf(err, "failed to Bind Param")
	}

	if err := validator.Struct(v); err != nil {
		return errors.Wrapf(err, "field validation failed")
	}

	if v, ok := v.(Validator); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrapf(err, "struct validation failed")
		}
	}

	return nil
}
