package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type ReportCommandController struct {
	service.ReportCommandService
}

var ReportCommandControllerSet = wire.NewSet(
	wire.Struct(new(ReportCommandController), "*"),
)

func (c *ReportCommandController) Report(ctx echo.Context, user entity.User) error {
	p := param.Report{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "validation report param")
	}

	cmd := converter.ConvertReportToCmd(&p)

	if err := c.ReportCommandService.Report(&user, cmd); err != nil {
		return errors.Wrap(err, "failed to report")
	}

	return ctx.JSON(http.StatusOK, "ok")
}

func (c *ReportCommandController) MarkAsDone(ctx echo.Context) error {
	p := param.SlackCallbackPayload{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "validation report submit body")
	}

	src := param.SlackCallback{}
	if err := json.Unmarshal([]byte(p.Payload), &src); err != nil {
		return errors.Wrap(err, "invalid slack report callback response type")
	}

	cmd, err := converter.ConvertSlackReportCallbackPayloadToCmd(&p)
	if err != nil {
		return errors.Wrap(err, "failed to convert to cmd")
	}

	err = c.ReportCommandService.MarkAsDone(cmd)
	if err != nil {
		return errors.Wrap(err, "failed to mark as done")
	}

	return ctx.JSON(http.StatusOK, "ok")
}
