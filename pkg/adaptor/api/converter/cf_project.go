package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

func (c *Converters) ConvertCfProjectSupportCommentListToOutput(comments *entity.CfProjectSupportCommentList) []*output.CfProjectSupportComment {
	response := make([]*output.CfProjectSupportComment, len([]*entity.CfProjectSupportComment(*comments)))
	for i, comment := range []*entity.CfProjectSupportComment(*comments) {
		response[i] = c.convertCfProjectSupportCommentToOutput(comment)
	}
	return response
}

func (c *Converters) ConvertCfProjectDetailListToOutput(list *entity.CfProjectDetailList, idIsFollowMap, idIsSupportMap map[int]bool) []*output.CfProject {
	response := make([]*output.CfProject, len(list.List))
	for i, project := range list.List {
		response[i] = c.ConvertCfProjectDetailToOutput(project, idIsFollowMap, idIsSupportMap)
	}
	return response
}

func (c *Converters) convertCfProjectSupportCommentToOutput(comment *entity.CfProjectSupportComment) *output.CfProjectSupportComment {
	return &output.CfProjectSupportComment{
		ID:        comment.ID,
		User:      c.NewUserSummaryFromUser(comment.User),
		Body:      comment.Body,
		CreatedAt: model.TimeResponse(comment.CreatedAt),
	}
}

func (c *Converters) ConvertCfProjectListInputToQuery(i *input.ListCfProject) *query.FindCfProjectQuery {
	return &query.FindCfProjectQuery{
		AreaID:       i.AreaID,
		SubAreaID:    i.SubAreaID,
		SubSubAreaID: i.SubSubAreaID,
		UserID:       i.UserID,
		SortBy:       i.SortBy,
	}
}

func (c *Converters) ConvertCfProjectDetailToOutput(project *entity.CfProjectDetail, idIsFollowMap, idIsSupportMap map[int]bool) *output.CfProject {
	// TODO
	if project.Snapshot == nil {
		project.Snapshot = &entity.CfProjectSnapshotDetail{}
	}

	thumbnails := make([]*output.CfProjectThumbnail, len(project.Snapshot.Thumbnails))
	for i, thumbnail := range project.Snapshot.Thumbnails {
		thumbnails[i] = &output.CfProjectThumbnail{
			Priority:  thumbnail.Priority,
			Thumbnail: thumbnail.Thumbnail,
		}
	}
	return &output.CfProject{
		ID:              project.ID,
		SnapshotID:      project.Snapshot.SnapshotID,
		Title:           project.Snapshot.Title,
		Summary:         project.Snapshot.Summary,
		Body:            project.Snapshot.Body,
		GoalPrice:       project.Snapshot.GoalPrice,
		AchievedPrice:   project.AchievedPrice,
		SupporterCount:  project.SupportCommentCount,
		FavoriteCount:   project.FavoriteCount,
		FacebookCount:   project.FacebookCount,
		TwitterCount:    project.TwitterCount,
		IsSupport:       idIsSupportMap[project.ID],
		Creator:         c.NewCreatorFromUser(project.User, idIsFollowMap[project.UserID]),
		Thumbnails:      thumbnails,
		AreaCategories:  c.ConvertAreaCategoriesToOutput(project.Snapshot.AreaCategories),
		ThemeCategories: c.ConvertThemeCategoriesToOutput(project.Snapshot.ThemeCategories),
		DeadLine:        model.TimeResponse(project.Snapshot.Deadline),
		CreatedAt:       model.TimeResponse(project.CreatedAt),
		EditedAt:        model.TimeResponse(project.Snapshot.CreatedAt),
	}
}

func (c Converters) ConvertSupportedCfProjectListInputToQuery(i *input.PaginationQuery) *query.FindListPaginationQuery {
	return &query.FindListPaginationQuery{
		Limit:  i.GetCfProjectLimit(),
		Offset: i.GetCfProjectOffset(),
	}
}
