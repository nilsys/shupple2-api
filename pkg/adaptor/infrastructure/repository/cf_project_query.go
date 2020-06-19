package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	CfProjectQueryRepositoryImpl struct {
		DAO
	}
)

var CfProjectQueryRepositorySet = wire.NewSet(
	wire.Struct(new(CfProjectQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.CfProjectQueryRepository), new(*CfProjectQueryRepositoryImpl)),
)

func (r *CfProjectQueryRepositoryImpl) LockCfProjectListByIDs(c context.Context, ids []int) (*entity.CfProjectList, error) {
	var rows entity.CfProjectList
	if err := r.LockDB(c).Where("id IN (?)", ids).Find(&rows.List).Error; err != nil {
		return nil, errors.Wrap(err, "failed find cf_project")
	}
	return &rows, nil
}
