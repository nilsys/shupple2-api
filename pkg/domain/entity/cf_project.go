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
		SupportCommentCount int // == SupporterCount
		FavoriteCount       int
		AchievedPrice       int
		Times
	}

	CfProjectSnapshotTable struct {
		ID          int `gorm:"primary_key"`
		CfProjectID int
		UserID      int
		Title       string
		Summary     string
		Body        string
		GoalPrice   int
		Deadline    time.Time
		IsAttention bool
		Times
	}

	CfProjectSnapshotThumbnail struct {
		CfProjectSnapshotID int
		Priority            int
		Thumbnail           string
		TimesWithoutDeletedAt
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
		Snapshot   *CfProjectSnapshot            `gorm:"foreignkey:ID;association_foreignkey:LatestSnapshotID"`
		User       *User                         `gorm:"foreignkey:ID;association_foreignkey:UserID"`
		Thumbnails []*CfProjectSnapshotThumbnail `gorm:"foreignkey:CfProjectSnapshotID;association_foreignkey:LatestSnapshotID"`
		Times
	}

	CfProjectSnapshot struct {
		CfProjectSnapshotTable
		AreaCategories  []*AreaCategory  `gorm:"many2many:cf_project_snapshot_area_category;jointable_foreignkey:cf_project_snapshot_id;"`
		ThemeCategories []*ThemeCategory `gorm:"many2many:cf_project_snapshot_theme_category;jointable_foreignkey:cf_project_snapshot_id;"`
	}

	CfProjectSupportComment struct {
		CfProjectSupportCommentTable
		User *User `gorm:"foreignkey:ID;association_foreignkey:UserID"`
	}

	CfProjectSnapshotAreaCategory struct {
		CfProjectSnapshotID int `gorm:"primary_key"`
		AreaCategoryID      int `gorm:"primary_key"`
		TimesWithoutDeletedAt
	}

	CfProjectSnapshotThemeCategory struct {
		CfProjectSnapshotID int `gorm:"primary_key"`
		ThemeCategoryID     int `gorm:"primary_key"`
		TimesWithoutDeletedAt
	}

	UserFavoriteCfProject struct {
		UserID      int `gorm:"primary_key"`
		CfProjectID int `gorm:"primary_key"`
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

func NewUserFavoriteCfProject(userID, projectID int) *UserFavoriteCfProject {
	return &UserFavoriteCfProject{
		UserID:      userID,
		CfProjectID: projectID,
	}
}

func (c *CfProjectTable) TableName() string {
	return "cf_project"
}

func (c *CfProjectSnapshotTable) TableName() string {
	return "cf_project_snapshot"
}

func (c *CfProjectSupportCommentTable) TableName() string {
	return "cf_project_support_comment"
}
