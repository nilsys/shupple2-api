package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/application/service"

	"github.com/stayway-corp/stayway-media-api/pkg/application/scenario"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
)

type (
	CfProjectQueryController struct {
		converter.Converters
		scenario.CfProjectQueryScenario
		service.CfProjectQueryService
	}
)

var CfProjectQueryControllerSet = wire.NewSet(
	wire.Struct(new(CfProjectQueryController), "*"),
)

func (c *CfProjectQueryController) List(ctx echo.Context, ouser entity.OptionalUser) error {
	i := &input.ListCfProject{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "failed bind input")
	}

	projects, idRelationFlgMap, idIsSupportMap, err := c.CfProjectQueryScenario.List(c.Converters.ConvertCfProjectListInputToQuery(i), &ouser)
	if err != nil {
		return errors.Wrap(err, "failed list cf_project")
	}

	return ctx.JSON(http.StatusOK, c.ConvertCfProjectDetailListToOutput(projects, idRelationFlgMap, idIsSupportMap))
}

func (c *CfProjectQueryController) Show(ctx echo.Context, ouser entity.OptionalUser) error {
	i := &input.IDParam{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "failed bind input")
	}
	project, idRelationFlgMap, idIsSupportMap, err := c.CfProjectQueryScenario.Show(i.ID, &ouser)
	if err != nil {
		return errors.Wrap(err, "failed show cf_project")
	}

	return ctx.JSON(http.StatusOK, c.ConvertCfProjectDetailToOutput(project, idRelationFlgMap, idIsSupportMap))

}

func (c *CfProjectQueryController) ListSupportComment(ctx echo.Context, ouser entity.OptionalUser) error {
	i := &input.ListCfProjectSupportComment{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "failed bind input")
	}

	comments, err := c.CfProjectQueryScenario.ListSupportComment(i.ID, i.GetLimit(), &ouser)
	if err != nil {
		return errors.Wrap(err, "failed list cf_project_support_comment")
	}

	return ctx.JSON(http.StatusOK, c.ConvertCfProjectSupportCommentListToOutput(comments))
}

func (c *CfProjectQueryController) ListSupported(ctx echo.Context, user entity.User) error {
	i := &input.PaginationQuery{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "failed bind input")
	}

	projects, idRelationFlgMap, idIsSupportMap, err := c.CfProjectQueryScenario.ListSupported(&user, c.Converters.ConvertSupportedCfProjectListInputToQuery(i))
	if err != nil {
		return errors.Wrap(err, "failed list cf_project")
	}

	return ctx.JSON(http.StatusOK, c.ConvertCfProjectDetailListToOutput(projects, idRelationFlgMap, idIsSupportMap))
}
