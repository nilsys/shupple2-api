package model

import (
	"encoding/json"
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

// MEMO: メール上の時間表記やチェックイン、チェックアウトで使われる事を想定している
// 名前は微妙
type TimeFront time.Time

const (
	timeFrontFmt = "2006/01/02 15:04"
)

func (tf *TimeFront) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	tm, err := time.Parse(timeFrontFmt, str)
	if err != nil {
		return err
	}
	*tf = TimeFront(tm.In(time.UTC))
	return nil
}

func (tf TimeFront) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(tf).In(util.JSTLoc).Format(timeFrontFmt))
}

func (tf TimeFront) ToString() string {
	return time.Time(tf).In(util.JSTLoc).Format(timeFrontFmt)
}
