package entity

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

type (
	CfProjectTable struct {
		ID                  int `gorm:"primary_key"`
		UserID              int
		LatestSnapshotID    null.Int
		SupportCommentCount int
		Times
	}

	CfProjectSnapshotTable struct {
		ID          int `gorm:"primary_key"`
		CfProjectID int
		UserID      int
		Title       string
		Summary     string
		Thumbnail   string
		Body        string
		GoalPrice   int
		Deadline    time.Time
		IsAttention bool
		Times
	}

	CfProjectSupportCommentTable struct {
		ID          int `gorm:"primary_key"`
		UserID      int
		CfProjectID int
		Body        string
		TimesWithoutDeletedAt
	}

	CfProject struct {
		CfProjectTable
		Snapshot *CfProjectSnapshotTable `gorm:"foreignkey:LatestSnapshotID"`
		User     *User                   `gorm:"foreignkey:UserID"`
		Times
	}

	CfProjectSnapshotAreaCategory struct {
		CfProjectSnapshotID int
		AreaCategoryID      int
		TimesWithoutDeletedAt
	}

	CfProjectSnapshotThemeCategory struct {
		CfProjectSnapshotID int
		ThemeCategoryID     int
		TimesWithoutDeletedAt
	}

	CfProjectList struct {
		List []*CfProject
	}
)

func (c *CfProjectList) ToIDMap() map[int]*CfProject {
	result := make(map[int]*CfProject, len(c.List))
	for _, summary := range c.List {
		result[summary.ID] = summary
	}
	return result
}

// 名前おかしい(そもそもentityの名前がおかしい)
func NewCfProjectSupportTable(userID, projectID int, body string) *CfProjectSupportCommentTable {
	return &CfProjectSupportCommentTable{
		UserID:      userID,
		CfProjectID: projectID,
		Body:        body,
	}
}

func (c *CfProjectSnapshotTable) TableName() string {
	return "cf_project_snapshot"
}

func (c *CfProjectSupportCommentTable) TableName() string {
	return "cf_project_support_comment"
}

func (c *CfProjectTable) TableName() string {
	return "cf_project"
}
