package repository

import (
	"context"

	"github.com/pkg/errors"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"

	"github.com/jinzhu/gorm"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"

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

func (r *CfProjectQueryRepositoryImpl) FindByID(id int) (*entity.CfProject, error) {
	var row entity.CfProject

	if err := r.DB(context.Background()).Find(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "cf_project(id=%d)", id)
	}

	return &row, nil
}

func (r *CfProjectQueryRepositoryImpl) FindListByQuery(query *query.FindCfProjectQuery) (*entity.CfProjectList, error) {
	var rows entity.CfProjectList

	q := r.buildFindList(query)
	if err := q.Joins("JOIN cf_project_snapshot ON cf_project.latest_snapshot_id = cf_project_snapshot.id").Order(query.SortBy.GetCfProjectOrderQuery()).Find(&rows.List).Error; err != nil {
		return nil, errors.Wrap(err, "failed find cf_project")
	}
	return &rows, nil
}

func (r *CfProjectQueryRepositoryImpl) LockCfProjectListByIDs(c context.Context, ids []int) (*entity.CfProjectList, error) {
	var rows entity.CfProjectList
	if err := r.LockDB(c).Where("id IN (?)", ids).Find(&rows.List).Error; err != nil {
		return nil, errors.Wrap(err, "failed find cf_project")
	}
	return &rows, nil
}

func (r *CfProjectQueryRepositoryImpl) Lock(c context.Context, id int) (*entity.CfProject, error) {
	var rows entity.CfProject
	if err := r.LockDB(c).Find(&rows, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "cf_project(id=%d)", id)
	}
	return &rows, nil
}

func (r *CfProjectQueryRepositoryImpl) FindSupportCommentListByCfProjectID(projectID, limit int) ([]*entity.CfProjectSupportComment, error) {
	var rows []*entity.CfProjectSupportComment

	if err := r.DB(context.Background()).Where("cf_project_id = ?", projectID).Order("created_at DESC").Limit(limit).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed find cf_project_support_comment.cf_project_id")
	}

	return rows, nil
}

func (r *CfProjectQueryRepositoryImpl) buildFindList(query *query.FindCfProjectQuery) *gorm.DB {
	q := r.DB(context.Background())

	if query.AreaID != 0 {
		q = q.Where("cf_project.latest_snapshot_id IN (SELECT cf_project_snapshot_id FROM cf_project_snapshot_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE area_id = ?))", query.AreaID)
	}

	if query.SubAreaID != 0 {
		q = q.Where("cf_project.latest_snapshot_id IN (SELECT cf_project_snapshot_id FROM cf_project_snapshot_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE sub_area_id = ?))", query.SubAreaID)
	}

	if query.SubSubAreaID != 0 {
		q = q.Where("cf_project.latest_snapshot_id IN (SELECT cf_project_snapshot_id FROM cf_project_snapshot_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE sub_sub_area_id = ?))", query.SubSubAreaID)
	}

	if query.SortBy == model.CfProjectSortByPush {
		q = q.Where("cf_project_snapshot.achieved_price / cf_project_snapshot.goal_price >= 0.7")
	}

	if query.SortBy == model.CfProjectSortByAttention {
		q = q.Where("cf_project.latest_snapshot_id IN (SELECT id FROM cf_project_snapshot WHERE is_attention = ?)", true)
	}

	return q
}
