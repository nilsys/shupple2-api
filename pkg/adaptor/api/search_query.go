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

type SearchQueryController struct {
	converter.Converters
	service.SearchQueryService
}

var SearchQueryControllerSet = wire.NewSet(
	wire.Struct(new(SearchQueryController), "*"),
)

func (c *SearchQueryController) ListSearchSuggestion(ctx echo.Context) error {
	p := &input.Keyward{}

	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation show search suggestion list")
	}

	suggestions, err := p.SearchResult(c.SearchQueryService.ListSuggestionsByKeyward, c.SearchQueryService.ListSuggestionsByType)

	if err != nil {
		return errors.Wrap(err, "failed to show search suggestion list")
	}

	return ctx.JSON(http.StatusOK, suggestions)
}
