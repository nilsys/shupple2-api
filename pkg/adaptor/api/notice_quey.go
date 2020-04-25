package api

import (
	"net/http"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type NoticeQueryController struct {
	converter.Converters
	NoticeQueryService service.NoticeQueryService
}

var NoticeQueryControllerSet = wire.NewSet(
	wire.Struct(new(NoticeQueryController), "*"),
)

func (c *NoticeQueryController) ListNotices(ctx echo.Context, user entity.User) error {
	notices, err := c.NoticeQueryService.ListNotice(&user)
	if err != nil {
		return errors.Wrap(err, "Failed to show notice list")
	}
	return ctx.JSON(http.StatusOK, c.ConvertListNoticeToOutput(notices))
}
