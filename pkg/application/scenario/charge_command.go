package scenario

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository/payjp"
	domainService "github.com/stayway-corp/stayway-media-api/pkg/domain/service"
)

type (
	ChargeCommandScenario interface {
		Create(user *entity.User, cmd *command.CreateCharge) (*entity.CaptureResult, error)
		InstantCreate(cmd *command.CreateCharge, cardToken string, address *entity.ShippingAddress) (*entity.CaptureResult, error)
	}

	ChargeCommandScenarioImpl struct {
		service.ChargeCommandService
		repository.CardCommandRepository
		repository.UserCommandRepository
		repository.ShippingCommandRepository
		service.UserCommandService
		PayjpCardCommandRepository payjp.CardCommandRepository
		payjp.CustomerQueryRepository
		payjp.CustomerCommandRepository
		domainService.UserValidatorDomainService
		service.TransactionService
	}
)

var ChargeCommandScenarioSet = wire.NewSet(
	wire.Struct(new(ChargeCommandScenarioImpl), "*"),
	wire.Bind(new(ChargeCommandScenario), new(*ChargeCommandScenarioImpl)),
)

func (s *ChargeCommandScenarioImpl) Create(user *entity.User, cmd *command.CreateCharge) (*entity.CaptureResult, error) {
	return s.ChargeCommandService.Create(user, cmd)
}

func (s *ChargeCommandScenarioImpl) InstantCreate(cmd *command.CreateCharge, cardToken string, address *entity.ShippingAddress) (*entity.CaptureResult, error) {
	// 簡易的なユーザーを作成
	userTiny, err := entity.NewIsNonLoginUserTiny(address.FullName())
	if err != nil {
		return nil, errors.Wrap(err, "failed gen uuid")
	}

	user := &entity.User{UserTiny: *userTiny}

	if err := s.UserValidatorDomainService.Do(user); err != nil {
		return nil, errors.Wrap(err, "failed validate user")
	}

	return s.ChargeCommandService.InstantCreate(user, cmd, cardToken, address)
}
