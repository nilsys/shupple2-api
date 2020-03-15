package repository

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

type InnQueryRepository interface {
	FindIDsByAreaID(areaId, subAreaId, subSubAreaId int) ([]int, error)
	FindAreaIDsByID(id int) (*entity.InnAreaTypeIDs, error)
}
