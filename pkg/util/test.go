package util

import (
	"fmt"
	"reflect"
)

func FillDummyString(v interface{}, idx int) {
	rv := reflect.ValueOf(v).Elem()
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		switch field.Type.Kind() {
		case reflect.String:
			rv.Field(i).SetString(fmt.Sprint(field.Name, idx))
		default:
			if field.Anonymous {
				FillDummyString(rv.Field(i).Addr().Interface(), idx)
			}
		}
	}
}
