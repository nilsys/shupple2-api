package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/google/wire"

	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type S3CommandController struct {
	service.S3CommandService
}

var S3CommandControllerSet = wire.NewSet(
	wire.Struct(new(S3CommandController), "*"),
)

func (c *S3CommandController) Post(ctx echo.Context, user entity.User) error {
	s3param := input.S3{}
	if err := BindAndValidate(ctx, &s3param); err != nil {
		return errors.Wrap(err, "validation post s3")
	}

	resp, err := c.S3CommandService.GenerateS3Signature(s3param.ContentType)
	if err != nil {
		return errors.Wrap(err, "failed generate s3 signature")
	}
	return ctx.JSON(http.StatusOK, resp)
}
