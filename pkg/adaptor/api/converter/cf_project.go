package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

func (c *Converters) ConvertCfProjectSupportCommentListToOutput(comments []*entity.CfProjectSupportComment) []*output.CfProjectSupportComment {
	response := make([]*output.CfProjectSupportComment, len(comments))
	for i, comment := range comments {
		response[i] = c.convertCfProjectSupportCommentToOutput(comment)
	}
	return response
}

func (c *Converters) ConvertCfProjectListToOutput(list *entity.CfProjectList) []*output.CfProject {
	response := make([]*output.CfProject, len(list.List))
	for i, project := range list.List {
		response[i] = c.ConvertCfProjectToOutput(project)
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
		SortBy:       i.SortBy,
	}
}

func (c *Converters) ConvertCfProjectToOutput(project *entity.CfProject) *output.CfProject {
	// TODO
	if project.Snapshot == nil {
		project.Snapshot = &entity.CfProjectSnapshot{}
	}
	return &output.CfProject{
		ID:              project.ID,
		SnapshotID:      project.Snapshot.ID,
		Title:           project.Snapshot.Title,
		Summary:         project.Snapshot.Summary,
		Thumbnail:       project.Snapshot.Thumbnail,
		Body:            project.Snapshot.Body,
		GoalPrice:       util.WithComma(project.Snapshot.GoalPrice),
		AchievedPrice:   util.WithComma(project.AchievedPrice),
		SupporterCount:  project.SupportCommentCount,
		FavoriteCount:   project.FavoriteCount,
		Creator:         c.NewCreatorFromUser(project.User),
		AreaCategories:  c.ConvertAreaCategoriesToOutput(project.Snapshot.AreaCategories),
		ThemeCategories: c.ConvertThemeCategoriesToOutput(project.Snapshot.ThemeCategories),
		DeadLine:        model.TimeResponse(project.Snapshot.Deadline),
	}
}
