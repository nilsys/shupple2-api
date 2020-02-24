// Code generated by go-enum
// DO NOT EDIT!

package wordpress

import (
	"fmt"
)

const (
	// CategoryTypeUndefined is a CategoryType of type Undefined
	CategoryTypeUndefined CategoryType = iota
	// CategoryTypeJapan is a CategoryType of type Japan
	CategoryTypeJapan
	// CategoryTypeWorld is a CategoryType of type World
	CategoryTypeWorld
	// CategoryTypeTheme is a CategoryType of type Theme
	CategoryTypeTheme
)

const _CategoryTypeName = "undefinedjapanworldtheme"

var _CategoryTypeMap = map[CategoryType]string{
	0: _CategoryTypeName[0:9],
	1: _CategoryTypeName[9:14],
	2: _CategoryTypeName[14:19],
	3: _CategoryTypeName[19:24],
}

// String implements the Stringer interface.
func (x CategoryType) String() string {
	if str, ok := _CategoryTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("CategoryType(%d)", x)
}

var _CategoryTypeValue = map[string]CategoryType{
	_CategoryTypeName[0:9]:   0,
	_CategoryTypeName[9:14]:  1,
	_CategoryTypeName[14:19]: 2,
	_CategoryTypeName[19:24]: 3,
}

// ParseCategoryType attempts to convert a string to a CategoryType
func ParseCategoryType(name string) (CategoryType, error) {
	if x, ok := _CategoryTypeValue[name]; ok {
		return x, nil
	}
	return CategoryType(0), fmt.Errorf("%s is not a valid CategoryType", name)
}

// MarshalText implements the text marshaller method
func (x CategoryType) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *CategoryType) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseCategoryType(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
