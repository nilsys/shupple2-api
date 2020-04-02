package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type FeatureQueryRepositoryImpl struct {
	DB *gorm.DB
}

var FeatureQueryRepositorySet = wire.NewSet(
	wire.Struct(new(FeatureQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.FeatureQueryRepository), new(*FeatureQueryRepositoryImpl)),
)

// Feature参照
func (r *FeatureQueryRepositoryImpl) FindByID(id int) (*entity.Feature, error) {
	var row entity.Feature
	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "feature(id=%d)", id)
	}
	return &row, nil
}

// QueryFeature参照(Feature詳細)
func (r *FeatureQueryRepositoryImpl) FindQueryFeatureByID(id int) (*entity.QueryFeature, error) {
	var row entity.QueryFeature
	if err := r.DB.Table(row.TableName()).First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "feature(id=%d)", id)
	}
	return &row, nil
}

// 作成日時降順でFeature一覧参照
func (r *FeatureQueryRepositoryImpl) FindList(query *query.FindListPaginationQuery) (*entity.FeatureList, error) {
	var rows entity.FeatureList
	if err := r.DB.Order("created_at DESC").Limit(query.Limit).Offset(query.Offset).Find(&rows.Features).Count(&rows.TotalNumber).Error; err != nil {
		return nil, errors.Wrapf(err, "Failed find features")
	}
	return &rows, nil
}
