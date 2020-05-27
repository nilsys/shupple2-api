package service

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	InterestQueryService interface {
		ListAll(group model.InterestGroup) ([]*entity.Interest, error)
	}

	InterestQueryServiceImpl struct {
		repository.InterestQueryRepository
	}
)

var InterestQueryServiceSet = wire.NewSet(
	wire.Struct(new(InterestQueryServiceImpl), "*"),
	wire.Bind(new(InterestQueryService), new(*InterestQueryServiceImpl)),
)

func (s *InterestQueryServiceImpl) ListAll(group model.InterestGroup) ([]*entity.Interest, error) {
	return s.InterestQueryRepository.FindAllByGroup(group)
}
