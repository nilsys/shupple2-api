package repository

import (
	"context"

	"github.com/google/wire"
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

func (r *CfProjectQueryRepositoryImpl) Lock(c context.Context, id int) (*entity.CfProject, error) {
	var rows entity.CfProject
	if err := r.LockDB(c).Find(&rows, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "cf_project(id=%d)", id)
	}
	return &rows, nil
}
