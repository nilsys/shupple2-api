package util

import (
	"fmt"
	"reflect"
)

func FillDymmyString(v interface{}, idx int) {
	rv := reflect.ValueOf(v).Elem()
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		if field.Type.Kind() == reflect.String {
			rv.Field(i).SetString(fmt.Sprint(field.Name, idx))
		}
	}
}
