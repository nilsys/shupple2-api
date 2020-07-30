package payjp

import (
	"github.com/payjp/payjp-go/v1"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
)

const (
	codeInvalidID       = "invalid_id"
	codeAlreadyHaveCard = "already_have_card"
	codeAlreadyExistID  = "already_exist_id"
)

func handleError(err error, message string, vs ...interface{}) error {
	if payjpError, ok := err.(*payjp.Error); ok {
		switch payjpError.Code {
		case codeInvalidID:
			return serror.New(err, serror.CodeNotFound, message, vs...)
		case codeAlreadyHaveCard:
			return serror.New(err, serror.CodeDuplicateCard, message, vs...)
		case codeAlreadyExistID:
			return serror.New(err, serror.CodeInvalidParam, message, vs...)
		}
	}

	return errors.Wrapf(err, message, vs...)
}
