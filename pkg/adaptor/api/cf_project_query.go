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

type (
	CfProjectQueryController struct {
		converter.Converters
		service.CfProjectQueryService
	}
)

var CfProjectQueryControllerSet = wire.NewSet(
	wire.Struct(new(CfProjectQueryController), "*"),
)

func (c *CfProjectQueryController) List(ctx echo.Context) error {
	i := &input.ListCfProject{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "failed bind input")
	}

	projects, err := c.CfProjectQueryService.List(c.Converters.ConvertCfProjectListInputToQuery(i))
	if err != nil {
		return errors.Wrap(err, "failed list cf_project")
	}

	return ctx.JSON(http.StatusOK, c.ConvertCfProjectListToOutput(projects))
}

func (c *CfProjectQueryController) Show(ctx echo.Context) error {
	i := &input.IDParam{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "failed bind input")
	}
	project, err := c.CfProjectQueryService.Show(i.ID)
	if err != nil {
		return errors.Wrap(err, "failed show cf_project")
	}

	return ctx.JSON(http.StatusOK, c.ConvertCfProjectToOutput(project))

}

func (c *CfProjectQueryController) ListSupportComment(ctx echo.Context) error {
	i := &input.ListCfProjectSupportComment{}
	comments, err := c.CfProjectQueryService.ListSupportComment(i.ID, i.GetLimit())
	if err != nil {
		return errors.Wrap(err, "failed list cf_project_support_comment")
	}

	return ctx.JSON(http.StatusOK, c.ConvertCfProjectSupportCommentListToOutput(comments))
}
