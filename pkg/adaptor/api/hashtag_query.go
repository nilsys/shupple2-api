package api

import (
	"net/http"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
)

type HashtagQueryController struct {
	converter.Converters
	service.HashtagQueryService
}

var HashtagQueryControllerSet = wire.NewSet(
	wire.Struct(new(HashtagQueryController), "*"),
)

func (c *HashtagQueryController) ListRecommendHashtag(ctx echo.Context) error {
	p := &input.ListRecommendHashTagParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrapf(err, "validation show recommend hashtag list input")
	}

	recommendHashTags, err := c.HashtagQueryService.ShowRecommendList(p.AreaID, p.SubAreaID, p.SubSubAreaID)
	if err != nil {
		return errors.Wrap(err, "failed show recommend hashtags")
	}

	return ctx.JSON(http.StatusOK, c.ConvertHashtagListToOutput(recommendHashTags))
}
