package entity

import (
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

/*
NOTE:
wordpressの方でareaGroupを作ることができない（階層構造が変わってしまいURLなどに影響が出るため）ので、
こちら側で一種のマスタデータとして管理する
*/
const (
	AreaGroupIDJapan = 1
	AreaGroupIDWorld = 2
)

type (
	Category struct {
		ID        int `gorm:"primary_key"`
		Name      string
		Type      model.CategoryType
		ParentID  *int
		CreatedAt time.Time `gorm:"-"`
		UpdatedAt time.Time `gorm:"-"`
	}
)
