package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type NoticeCommandController struct {
	NoticeCommand service.NoticeCommandService
}

var NoticeCommandControllerSet = wire.NewSet(
	wire.Struct(new(NoticeCommandController), "*"),
)

func (c *NoticeCommandController) MarkAsRead(ctx echo.Context, user entity.User) error {
	var i input.IDParam
	if err := BindAndValidate(ctx, &i); err != nil {
		return errors.Wrap(err, "failed bind input")
	}

	unreadCount, err := c.NoticeCommand.MarkAsRead(&user, i.ID)
	if err != nil {
		return errors.Wrap(err, "failed mark as read push_notice")
	}

	return ctx.JSON(http.StatusOK, struct {
		UnreadCount int `json:"unreadCount"`
	}{UnreadCount: unreadCount})
}
