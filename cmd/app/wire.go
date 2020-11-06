//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/uma-co82/shupple2-api/pkg/adaptor/api"
	"github.com/uma-co82/shupple2-api/pkg/adaptor/api/converter"
	"github.com/uma-co82/shupple2-api/pkg/adaptor/api/middleware"
	"github.com/uma-co82/shupple2-api/pkg/adaptor/infrastructure/repository"
	"github.com/uma-co82/shupple2-api/pkg/application/service"
	"github.com/uma-co82/shupple2-api/pkg/config"
)

var controllerSet = wire.NewSet(
	api.HealthCheckControllerSet,
	api.UserCommandControllerSet,
	api.UserQueryControllerSet,
	api.ArrangeScheduleRequestCommandControllerSet,
)

var serviceSet = wire.NewSet(
	service.UserCommandServiceSet,
	service.UserQueryServiceSet,
	service.ArrangeScheduleRequestCommandServiceSet,
)

func InitializeApp(configFilePath config.FilePath) (*App, error) {
	wire.Build(
		service.ProvideAuthService,
		echo.New,
		controllerSet,
		serviceSet,
		middleware.AuthorizeSet,
		wire.Struct(new(App), "*"),
		config.GetConfig,
		wire.FieldsOf(new(*config.Config), "AWS"),
		converter.ConverterSet,
		repository.RepositoriesSet,
	)

	return new(App), nil
}
