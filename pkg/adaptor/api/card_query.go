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

type CardQueryController struct {
	service.CardQueryService
	converter.Converters
}

var CardQueryControllerSet = wire.NewSet(
	wire.Struct(new(CardQueryController), "*"),
)

func (c *CardQueryController) ShowCard(ctx echo.Context, user entity.User) error {
	card, err := c.CardQueryService.ShowCard(&user)
	if err != nil {
		return errors.Wrap(err, "failed show card")
	}
	return ctx.JSON(http.StatusOK, c.ConvertCardToOutput(card))
}
