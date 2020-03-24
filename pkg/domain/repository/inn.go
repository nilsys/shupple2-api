package repository

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type InnQueryRepository interface {
	FindIDsByAreaID(areaId, subAreaId, subSubAreaId int) ([]int, error)
	FindAreaIDsByID(id int) (*entity.InnAreaTypeIDs, error)
	FindByParams(query *query.FindInn) (*entity.Inns, error)
}
