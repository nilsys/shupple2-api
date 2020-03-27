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
	feature := Feature{FeatureTiny: tiny}
	feature.SetPosts(postIDs)
	return feature
}

func (feature *Feature) SetPosts(postIDs []int) {
	feature.PostIDs = make([]*FeaturePost, len(postIDs))
	for i, p := range postIDs {
		feature.PostIDs[i] = &FeaturePost{
			FeatureID: feature.ID,
			PostID:    p,
		}
	}
}

func (queryFeature QueryFeature) TableName() string {
	return "feature"
}
