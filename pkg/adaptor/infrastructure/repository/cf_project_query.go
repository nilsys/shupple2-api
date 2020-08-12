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

func (r *CfProjectQueryRepositoryImpl) FindByID(id int) (*entity.CfProjectDetail, error) {
	var row entity.CfProjectDetail

	if err := r.DB(context.Background()).Find(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "cf_project(id=%d)", id)
	}

	return &row, nil
}

func (r *CfProjectQueryRepositoryImpl) FindListByQuery(query *query.FindCfProjectQuery) (*entity.CfProjectDetailList, error) {
	var rows entity.CfProjectDetailList

	q := r.buildFindList(query)
	if err := q.Joins("JOIN cf_project_snapshot ON cf_project.latest_snapshot_id = cf_project_snapshot.id").
		Order(query.SortBy.GetCfProjectOrderQuery()).
		Limit(query.Limit).
		Offset(query.Offset).
		Find(&rows.List).Offset(0).Count(&rows.TotalNumber).Error; err != nil {
		return nil, errors.Wrap(err, "failed find cf_project")
	}
	return &rows, nil
}

func (r *CfProjectQueryRepositoryImpl) FindSupportCommentListByCfProjectID(projectID, limit int) (*entity.CfProjectSupportCommentList, error) {
	var rows entity.CfProjectSupportCommentList

	if err := r.DB(context.Background()).Where("cf_project_id = ?", projectID).Order("created_at DESC").Limit(limit).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed find cf_project_support_comment.cf_project_id")
	}

	return &rows, nil
}

// 達成メールが未送信かつ目標金額に達していているCfProject
func (r *CfProjectQueryRepositoryImpl) FindNotSentAchievementNoticeEmailAndAchievedListByLastID(lastID, limit int) (*entity.CfProjectDetailList, error) {
	var rows entity.CfProjectDetailList
	if err := r.DB(context.Background()).
		Joins("INNER JOIN cf_project_snapshot ON cf_project.latest_snapshot_id = cf_project_snapshot.id").
		Where("cf_project.id > ? AND cf_project_snapshot.goal_price <= cf_project.achieved_price AND cf_project.is_sent_achievement_email = false", lastID).
		Limit(limit).
		Find(&rows.List).Error; err != nil {
		return nil, errors.Wrap(err, "failed find cf_project")
	}
	return &rows, nil
}

// 通知メールが送信されていない報告(post)を持つCfProject
func (r *CfProjectQueryRepositoryImpl) FindNotSentNewPostNoticeEmailByLastID(lastID, limit int) (*entity.CfProjectDetailList, error) {
	var rows entity.CfProjectDetailList

	if err := r.DB(context.Background()).
		Where("is_sent_new_post_email = false AND latest_post_id IS NOT NULL AND id > ?", lastID).
		Limit(limit).
		Find(&rows.List).Error; err != nil {
		return nil, errors.Wrap(err, "failed find cf_project")
	}

	return &rows, nil
}

// Userが支援(購入)したCfProject一覧
func (r *CfProjectQueryRepositoryImpl) FindSupportedListByUserID(userID int, query *query.FindListPaginationQuery) (*entity.CfProjectDetailList, error) {
	var rows entity.CfProjectDetailList

	if err := r.DB(context.Background()).
		Joins("INNER JOIN (SELECT cf_project_id, MAX(created_at) AS created_at FROM payment_cf_return_gift WHERE payment_id IN (SELECT id FROM payment WHERE user_id = ?) GROUP BY cf_project_id) pc ON cf_project.id = pc.cf_project_id", userID).
		Order("pc.created_at DESC").
		Limit(query.Limit).
		Offset(query.Offset).
		Find(&rows.List).
		Offset(0).
		Count(&rows.TotalNumber).
		Error; err != nil {
		return nil, errors.Wrap(err, "failed find supported cf_project")
	}

	return &rows, nil
}

// UserがCfProjectを支援したかどうか
// cfProjectID:isSupportedのMap
func (r *CfProjectQueryRepositoryImpl) IsSupported(userID int, projectIDs []int) (map[int]bool, error) {
	var rows entity.UserSupportCfProjectList

	if err := r.DB(context.Background()).
		Select("payment.user_id AS user_id, payment_cf_return_gift.cf_project_id AS cf_project_id").
		Table("payment").
		Joins("INNER JOIN payment_cf_return_gift ON payment.id = payment_cf_return_gift.payment_id").
		Where("payment_cf_return_gift.cf_project_id IN (?) AND payment.user_id", projectIDs).
		Group("payment.user_id, payment_cf_return_gift.cf_project_id").
		Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed find user supported cf_project")
	}

	return rows.ToIDIsSupportMap(projectIDs), nil
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

	if query.UserID != 0 {
		q = q.Where("cf_project.user_id = ?", query.UserID)
	}

	if query.SortBy == model.CfProjectSortByPush {
		q = q.Where("cf_project.achieved_price / cf_project.goal_price >= 0.7")
	}

	if query.SortBy == model.CfProjectSortByAttention {
		q = q.Where("cf_project.latest_snapshot_id IN (SELECT id FROM cf_project_snapshot WHERE is_attention = ?)", true)
	}

	return q
}
