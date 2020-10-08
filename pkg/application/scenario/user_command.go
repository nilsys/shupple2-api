package scenario

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	UserCommandScenario interface {
		DeleteImage(user *entity.User, imageType model.MediaType) error
	}

	UserCommandScenarioImpl struct {
		service.UserCommandService
	}
)

var UserCommandScenarioSet = wire.NewSet(
	wire.Struct(new(UserCommandScenarioImpl), "*"),
	wire.Bind(new(UserCommandScenario), new(*UserCommandScenarioImpl)),
)

func (s *UserCommandScenarioImpl) DeleteImage(user *entity.User, imageType model.MediaType) error {
	// アダプターでバリデーションは掛けているがチェック
	if imageType == model.MediaTypeUserIcon {
		return s.UserCommandService.DeleteUserIcon(user)
	} else if imageType == model.MediaTypeUserHeader {
		return s.UserCommandService.DeleteUserHeader(user)
	}

	return nil
}
