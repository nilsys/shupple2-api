package entity

import (
	"database/sql"
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
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
		SubAreaID              sql.NullInt64
		SubSubAreaID           sql.NullInt64
		MetasearchAreaID       int
		MetasearchSubAreaID    int
		MetasearchSubSubAreaID int
		CreatedAt              time.Time `gorm:"-;default:current_timestamp"`
		UpdatedAt              time.Time `gorm:"-;default:current_timestamp"`
	}
)

func (ac AreaCategory) CategoryType() string {
	return ac.Type.String()
}
