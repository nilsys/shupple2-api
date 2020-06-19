package entity

import "time"

type (
	CfProjectTable struct {
		ID               int `gorm:"primary_key"`
		UserID           int
		LatestSnapshotID int
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
