package api

import (
	"net/http"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
)

type SearchQueryController struct {
	service.SearchQueryService
}

var SearchQueryControllerSet = wire.NewSet(
	wire.Struct(new(SearchQueryController), "*"),
)

func (c *SearchQueryController) ShowSearchSuggestionList(ctx echo.Context) error {
	p := &param.Keyward{}

	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation show search suggestion list")
	}

	suggestions, err := c.SearchQueryService.ShowSearchSuggestionListByKeyward(p.Q)
	if err != nil {
		return errors.Wrap(err, "failed to show search suggestion list")
	}

	return ctx.JSON(http.StatusOK, suggestions)
}