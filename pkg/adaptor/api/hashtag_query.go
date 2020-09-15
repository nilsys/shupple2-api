package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"

	"github.com/stayway-corp/stayway-media-api/pkg/application/scenario"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
)

type HashtagQueryController struct {
	converter.Converters
	scenario.HashtagQueryScenario
}

var HashtagQueryControllerSet = wire.NewSet(
	wire.Struct(new(HashtagQueryController), "*"),
)

func (c *HashtagQueryController) ListRecommendHashtag(ctx echo.Context, ouser entity.OptionalUser) error {
	p := &input.ListRecommendHashtagParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrapf(err, "validation show recommend hashtag list input")
	}

	recommendHashTags, isFollowMap, err := c.HashtagQueryScenario.ListRecommend(p.AreaID, p.SubAreaID, p.SubSubAreaID, p.PerPage, &ouser)
	if err != nil {
		return errors.Wrap(err, "failed show recommend hashtags")
	}

	return ctx.JSON(http.StatusOK, c.ConvertHashtagListToOutput(recommendHashTags, isFollowMap))
}

func (c *HashtagQueryController) Show(ctx echo.Context, ouser entity.OptionalUser) error {
	i := &input.ShowHashtag{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "failed bind input")
	}

	hashtag, isFollowMap, err := c.HashtagQueryScenario.Show(string(i.Name), &ouser)
	if err != nil {
		return errors.Wrap(err, "failed show hashtag")
	}

	return ctx.JSON(http.StatusOK, output.NewHashtag(hashtag.ID, hashtag.Name, isFollowMap[hashtag.ID]))
}
