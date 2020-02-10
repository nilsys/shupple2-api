package repository

type InnQueryRepository interface {
	FindIDsByAreaID(areaId, subAreaId, subSubAreaId int) ([]int, error)
}
