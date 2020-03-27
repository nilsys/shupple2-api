package serror

import (
	"fmt"

	"github.com/pkg/errors"
)

/*
List系のメソッドで1件も見つからなかった時にNotFoundを返すための関数
List系のメソッドでで正常に取得できて0件だった場合、errorが返却されないためそこをハンドリングする

```
rows, err := GetRows()
if err != nil {
	return errors.Wrap(err, "")
}
if len(rows) == 0 {
	return errors.New("")
}
```
を
```
rows, err := GetRows()
if err != nil || len(rows) == 0 {
	return serror.NewResourcesNotFoundError(err, "")
}
```
と書くためのメソッド
*/

func NewResourcesNotFoundError(err error, resource string, v ...interface{}) error {
	if len(v) > 0 {
		resource = fmt.Sprintf(resource, v...)
	}

	if err != nil {
		return errors.Wrapf(err, "%s not found", resource)
	}

	return errors.Errorf("%s not found", resource)
}
