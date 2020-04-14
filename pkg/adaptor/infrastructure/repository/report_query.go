package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type ReportQueryRepositoryImpl struct {
	DB *gorm.DB
}

var ReportQueryRepositorySet = wire.NewSet(
	wire.Struct(new(ReportQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.ReportQueryRepository), new(*ReportQueryRepositoryImpl)),
)

func (r *ReportQueryRepositoryImpl) IsExist(userID, targetID int, targetType model.ReportTargetType) (bool, error) {
	var row entity.Report

	err := r.DB.Where("user_id = ?", userID).Where("target_id = ?", targetID).Where("target_type = ?", targetType).First(&row).Error

	return ErrorToIsExist(err, "report(user_id=%d,target_id=%d,target_type=%s)", userID, targetID, targetType)
}
