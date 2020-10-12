package entity

import (
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
		Type         model.AreaCategoryType
		AreaGroup    model.AreaGroup
		AreaID       int
		SubAreaID    null.Int
		SubSubAreaID null.Int
		TimesWithoutDeletedAt
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

	AreaCategories []*AreaCategory
)

func (ac AreaCategory) CategoryType() string {
	return ac.Type.String()
}

func (acd *AreaCategoryDetail) Set(area *AreaCategory) {
	acd.AreaCategory = *area
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

func (a AreaCategories) AreaIDs() []int {
	resolve := make([]int, 0, len(a))
	for _, area := range a {
		if area.Type == model.AreaCategoryTypeArea {
			resolve = append(resolve, area.ID)
		}
	}
	return resolve
}

func (a AreaCategories) SubAreaIDs() []int {
	resolve := make([]int, 0, len(a))
	for _, area := range a {
		if area.Type == model.AreaCategoryTypeSubArea {
			resolve = append(resolve, area.ID)
		}
	}
	return resolve
}

func (a AreaCategories) SubSubAreaIDs() []int {
	resolve := make([]int, 0, len(a))
	for _, area := range a {
		if area.Type == model.AreaCategoryTypeSubSubArea {
			resolve = append(resolve, area.ID)
		}
	}
	return resolve
}
