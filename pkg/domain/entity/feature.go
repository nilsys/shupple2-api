package entity

import (
	"github.com/stayway-corp/stayway-media-api/pkg/util"
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
		Times
	}

	Feature struct {
		FeatureTiny
		PostIDs []*FeaturePost
	}

	FeatureList struct {
		TotalNumber int
		Features    []*Feature
	}

	FeaturePost struct {
		FeatureID int `gorm:"primary_key"`
		PostID    int `gorm:"primary_key"`
	}

	FeatureDetail struct {
		FeatureTiny
		User    *User          `gorm:"foreignkey:UserID"`
		PostIDs []*FeaturePost `gorm:"foreignkey:FeatureID"`
	}

	FeatureDetailWithPosts struct {
		FeatureTiny
		Posts []*PostListTiny `gorm:"many2many:feature_post;jointable_foreignkey:feature_id;association_jointable_foreignkey:post_id;"`
		User  *User           `gorm:"foreignkey:UserID"`
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

func (fd FeatureDetail) TableName() string {
	return "feature"
}

func (f *FeatureDetailWithPosts) SetPosts(posts []*PostListTiny) {
	f.Posts = posts
}

func (f *FeatureDetailWithPosts) SetFeature(feature FeatureDetail) {
	f.FeatureTiny = feature.FeatureTiny
	f.User = feature.User
}

func (f *FeatureDetailWithPosts) AreaCategoryIDs() []int {
	ids := make([]int, 0)

	for _, post := range f.Posts {
		for _, area := range post.AreaCategories {
			ids = append(ids, area.AreaID)

			if area.SubAreaID.Valid {
				ids = append(ids, int(area.SubAreaID.Int64))
			}

			if area.SubSubAreaID.Valid {
				ids = append(ids, int(area.SubSubAreaID.Int64))
			}
		}
	}

	return util.RemoveDuplicatesAndZeroFromIntSlice(ids)
}

func (f *FeatureDetailWithPosts) ThemeCategoryIDs() []int {
	ids := make([]int, 0)

	for _, post := range f.Posts {
		for _, theme := range post.ThemeCategories {
			ids = append(ids, theme.ThemeID)

			if theme.SubThemeID.Valid {
				ids = append(ids, int(theme.SubThemeID.Int64))
			}
		}
	}

	return util.RemoveDuplicatesAndZeroFromIntSlice(ids)
}
