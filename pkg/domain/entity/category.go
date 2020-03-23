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
		ID                     int `gorm:"primary_key"`
		Name                   string
		Slug                   string
		Type                   model.CategoryType
		MetasearchAreaID       int `gorm:"-"` // 基本的にアプリケーション内で更新することはないのでgorm:"-"にしておく。取得はできる。
		MetasearchSubAreaID    int `gorm:"-"` // 基本的にアプリケーション内で更新することはないのでgorm:"-"にしておく。取得はできる。
		MetasearchSubSubAreaID int `gorm:"-"` // 基本的にアプリケーション内で更新することはないのでgorm:"-"にしておく。取得はできる。
		ParentID               *int
		CreatedAt              time.Time `gorm:"-"`
		UpdatedAt              time.Time `gorm:"-"`
	}
)
