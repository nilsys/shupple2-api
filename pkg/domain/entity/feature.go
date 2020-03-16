package entity

import (
	"time"
)

type (
	FeatureTiny struct {
		ID            int `gorm:"primary_key"`
		UserID        int
		Slug          string
		Thumbnail     string
		Title         string
		Body          string
		FacebookCount int
		TwitterCount  int
		Views         int
		CreatedAt     time.Time
		UpdatedAt     time.Time
		DeletedAt     *time.Time
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

func (queryFeature QueryFeature) TableName() string {
	return "feature"
}
