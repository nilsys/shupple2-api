package entity

import (
	"strconv"
	"time"
)

type (
	FeatureTiny struct {
		ID        int `gorm:"primary_key"`
		UserID    int
		Slug      string
		Title     string
		Body      string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time
	}

	Feature struct {
		FeatureTiny
		PostIDs []*FeaturePost
	}

	FeaturePost struct {
		FeatureID int `gorm:"primary_key"`
		PostID    int `gorm:"primary_key"`
	}

	QueryFeature struct {
		FeatureTiny
		Posts []*Post `gorm:"many2many:feature_post;jointable_foreignkey:feature_id;"`
		User  *User   `gorm:"foreignkey:UserID"`
	}
)

func NewFeature(tiny FeatureTiny, postIDs []int) Feature {
	featurePostIDs := make([]*FeaturePost, len(postIDs))
	for i, p := range postIDs {
		featurePostIDs[i] = &FeaturePost{
			FeatureID: tiny.ID,
			PostID:    p,
		}
	}

	return Feature{
		tiny,
		featurePostIDs,
	}
}

// TODO:
func (tiny FeatureTiny) GenerateThumbnailURL() string {
	return "https://file.staywayy.jp/feature/" + strconv.Itoa(tiny.ID)
}

func (queryFeature QueryFeature) TableName() string {
	return "feature"
}
