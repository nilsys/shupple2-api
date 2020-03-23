package dto

import (
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity/wordpress"
)

type (
	WordpressCategory struct {
		Attributes json.RawMessage `json:"acf"` // acfが空の時に配列を返してくるのでやむなし
		wordpress.Category
	}

	WordpressCategoryAttributes struct {
		CategoryType string `json:"category_type"`
	}

	WordpressCategories []*WordpressCategory
)

func (wc WordpressCategory) ToEntity() (*wordpress.Category, error) {
	if len(wc.Attributes) == 0 || wc.Attributes[0] == '[' { //存在しないか配列の場合
		wc.Category.Type = wordpress.CategoryTypeUndefined
		return &wc.Category, nil
	}

	var attrs WordpressCategoryAttributes
	if err := json.Unmarshal(wc.Attributes, &attrs); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal wordpress category attributes")
	}

	// CategoryTypeが未設定の場合
	if attrs.CategoryType == "" {
		wc.Category.Type = wordpress.CategoryTypeUndefined
		return &wc.Category, nil
	}

	typeKey := strings.TrimSpace(strings.Split(attrs.CategoryType, ":")[0])
	typeValue, err := wordpress.ParseCategoryType(typeKey)
	if err != nil {
		return nil, errors.Wrap(err, "invalid wordpress category type")
	}

	wc.Type = typeValue
	return &wc.Category, nil
}

func (wcs WordpressCategories) ToEntities() ([]*wordpress.Category, error) {
	res := make([]*wordpress.Category, len(wcs))
	var err error
	for i, wc := range wcs {
		res[i], err = wc.ToEntity()
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return res, nil
}
