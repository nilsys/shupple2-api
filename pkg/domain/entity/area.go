package entity

import (
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"gopkg.in/guregu/null.v3"
)

/**
NOTE:
AreaCategoryはかなり正規化を崩した構造になっているのでスキーマを変える際は気をつける。
崩れ方はAreaCategoryの中だけで完結するようにはなっているはず。
*/

type (
	AreaCategory struct {
		CategoryBase
		Type                   model.AreaCategoryType
		AreaGroup              model.AreaGroup
		AreaID                 int
		SubAreaID              null.Int
		SubSubAreaID           null.Int
		MetasearchAreaID       int
		MetasearchSubAreaID    int
		MetasearchSubSubAreaID int
		CreatedAt              time.Time `gorm:"-;default:current_timestamp"`
		UpdatedAt              time.Time `gorm:"-;default:current_timestamp"`
	}

	AreaCategoryDetail struct {
		AreaCategory
		Area       *AreaCategory
		SubArea    *AreaCategory
		SubSubArea *AreaCategory
	}

	AreaCategoryWithPostCount struct {
		AreaCategory
		PostCount int
	}
)

func (ac AreaCategory) CategoryType() string {
	return ac.Type.String()
}

func (acd *AreaCategoryDetail) SetArea(area *AreaCategory) {
	acd.Area = area
}

func (acd *AreaCategoryDetail) SetSubArea(subArea *AreaCategory) {
	acd.SubArea = subArea
}

func (acd *AreaCategoryDetail) SetSubSubArea(subSubArea *AreaCategory) {
	acd.SubSubArea = subSubArea
}

func (acwpc *AreaCategoryWithPostCount) TableName() string {
	return "area_category"
}
