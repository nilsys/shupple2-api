package repository

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	InterestQueryRepository interface {
		FindAllByGroup(group model.InterestGroup) ([]*entity.Interest, error)
	}
)
