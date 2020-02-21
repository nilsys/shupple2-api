package param

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"

// 検索キーワード(単数)
type Keyward struct {
	Q string `query:"q"`
}

func (keyward *Keyward) Validate() error {
	if keyward.Q == "" {
		return serror.New(nil, serror.CodeInvalidParam, "Invalid show search suggestions param")
	}

	return nil
}
