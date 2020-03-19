package api

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/application/scenario"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"net/http"
)

type ReviewCommandController struct {
	scenario.ReviewCommandScenario
	service.ReviewCommandService
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

func (c *ReviewCommandController) StoreReviewComment(ctx echo.Context, user entity.User) error {
	reviewCommentParam := &param.CreateReviewCommand{}
	if err := BindAndValidate(ctx, reviewCommentParam); err != nil {
		return errors.Wrap(err, "Failed to bind parameters")
	}

	if err := c.ReviewCommandService.CreateReviewComment(&user, reviewCommentParam.ID, reviewCommentParam.Body); err != nil {
		return errors.Wrap(err, "Failed to create review comment")
	}

	return ctx.JSON(http.StatusOK, nil)
}