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
		SnapshotID  int `gorm:"column:id;primary_key"`
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
		CfProjectSnapshotID int `gorm:"primary_key"`
		Priority            int `gorm:"primary_key"`
		Thumbnail           string
		TimesWithoutDeletedAt
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

	CfProjectSnapshot struct {
		CfProjectSnapshotTable
		Thumbnails       []*CfProjectSnapshotThumbnail     `gorm:"foreignkey:CfProjectSnapshotID;association_foreignkey:SnapshotID"`
		AreaCategoryIDs  []*CfProjectSnapshotAreaCategory  `gorm:"foreignkey:CfProjectSnapshotID;association_foreignkey:SnapshotID"`
		ThemeCategoryIDs []*CfProjectSnapshotThemeCategory `gorm:"foreignkey:CfProjectSnapshotID;association_foreignkey:SnapshotID"`
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
		Snapshot CfProjectSnapshot
	}

	CfProjectDetail struct {
		CfProjectTable
		Snapshot *CfProjectSnapshotDetail `gorm:"foreignkey:ID;association_foreignkey:LatestSnapshotID"`
		User     *User                    `gorm:"foreignkey:ID;association_foreignkey:UserID"`
	}

	CfProjectSnapshotDetail struct {
		CfProjectSnapshotTable
		Thumbnails      []*CfProjectSnapshotThumbnail `gorm:"foreignkey:CfProjectSnapshotID;association_foreignkey:SnapshotID"`
		AreaCategories  []*AreaCategory               `gorm:"many2many:cf_project_snapshot_area_category;jointable_foreignkey:cf_project_snapshot_id;"`
		ThemeCategories []*ThemeCategory              `gorm:"many2many:cf_project_snapshot_theme_category;jointable_foreignkey:cf_project_snapshot_id;"`
	}

	CfProjectSupportComment struct {
		CfProjectSupportCommentTable
		User *User `gorm:"foreignkey:ID;association_foreignkey:UserID"`
	}

	UserFavoriteCfProject struct {
		UserID      int `gorm:"primary_key"`
		CfProjectID int `gorm:"primary_key"`
		TimesWithoutDeletedAt
	}

	CfProjectDetailList struct {
		List []*CfProjectDetail
	}
)

func (c *CfProjectDetailList) ToIDMap() map[int]*CfProjectDetail {
	result := make(map[int]*CfProjectDetail, len(c.List))
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

func (p *CfProjectSnapshot) SetThumbnails(thumbnails []string) {
	p.Thumbnails = make([]*CfProjectSnapshotThumbnail, len(thumbnails))
	for i, t := range thumbnails {
		p.Thumbnails[i] = &CfProjectSnapshotThumbnail{
			CfProjectSnapshotID: p.SnapshotID,
			Thumbnail:           t,
			Priority:            i + 1,
		}
	}
}

func (c *CfProjectSupportCommentTable) TableName() string {
	return "cf_project_support_comment"
}

func (p *CfProjectSnapshot) SetAreaCategories(areaCategoryIDs []int) {
	p.AreaCategoryIDs = make([]*CfProjectSnapshotAreaCategory, len(areaCategoryIDs))
	for i, c := range areaCategoryIDs {
		p.AreaCategoryIDs[i] = &CfProjectSnapshotAreaCategory{
			CfProjectSnapshotID: p.SnapshotID,
			AreaCategoryID:      c,
		}
	}
}

func (p *CfProjectSnapshot) SetThemeCategories(themeCategoryIDs []int) {
	p.ThemeCategoryIDs = make([]*CfProjectSnapshotThemeCategory, len(themeCategoryIDs))
	for i, c := range themeCategoryIDs {
		p.ThemeCategoryIDs[i] = &CfProjectSnapshotThemeCategory{
			CfProjectSnapshotID: p.SnapshotID,
			ThemeCategoryID:     c,
		}
	}
}
