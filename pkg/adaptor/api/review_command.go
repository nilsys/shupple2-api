package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/application/scenario"
)

type ReviewCommandController struct {
	scenario.ReviewCommandScenario
}

var ReviewCommandControllerSet = wire.NewSet(
	wire.Struct(new(ReviewCommandController), "*"),
)

// TODO: 認証処理挟む
func (c *ReviewCommandController) Store(ctx echo.Context) error {
	p := &param.StoreReviewParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation store review param")
	}

	// TODO: 認証処理挟む
	if err := c.ReviewCommandScenario.Create(&entity.User{ID: 1}, converter.ConvertCreateReviewParamToCommand(p)); err != nil {
		return errors.Wrap(err, "failed to store review")
	}

	return ctx.JSON(http.StatusOK, "hoge")
}
