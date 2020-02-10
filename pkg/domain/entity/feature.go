package entity

import "time"

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
