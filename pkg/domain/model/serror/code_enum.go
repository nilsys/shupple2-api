// Code generated by go-enum
// DO NOT EDIT!

package serror

import (
	"fmt"
)

const (
	// CodeUndefined is a Code of type Undefined
	CodeUndefined Code = iota
	// CodeInvalidParam is a Code of type InvalidParam
	CodeInvalidParam
	// CodeNotFound is a Code of type NotFound
	CodeNotFound
	// CodeImportDeleted is a Code of type ImportDeleted
	CodeImportDeleted
	// CodeUnauthorized is a Code of type Unauthorized
	CodeUnauthorized
	// CodeForbidden is a Code of type Forbidden
	CodeForbidden
	// CodeInvalidCategoryType is a Code of type InvalidCategoryType
	CodeInvalidCategoryType
	// CodePayAgentError is a Code of type PayAgentError
	CodePayAgentError
	// CodeDuplicateCard is a Code of type DuplicateCard
	CodeDuplicateCard
	// CodeUnsupportedMedia is a Code of type UnsupportedMedia
	CodeUnsupportedMedia
	// CodeDuplicateReport is a Code of type DuplicateReport
	CodeDuplicateReport
	// CodeExpired is a Code of type Expired
	CodeExpired
)

const _CodeName = "UndefinedInvalidParamNotFoundImportDeletedUnauthorizedForbiddenInvalidCategoryTypePayAgentErrorDuplicateCardUnsupportedMediaDuplicateReportExpired"

var _CodeMap = map[Code]string{
	0:  _CodeName[0:9],
	1:  _CodeName[9:21],
	2:  _CodeName[21:29],
	3:  _CodeName[29:42],
	4:  _CodeName[42:54],
	5:  _CodeName[54:63],
	6:  _CodeName[63:82],
	7:  _CodeName[82:95],
	8:  _CodeName[95:108],
	9:  _CodeName[108:124],
	10: _CodeName[124:139],
	11: _CodeName[139:146],
}

// String implements the Stringer interface.
func (x Code) String() string {
	if str, ok := _CodeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Code(%d)", x)
}

var _CodeValue = map[string]Code{
	_CodeName[0:9]:     0,
	_CodeName[9:21]:    1,
	_CodeName[21:29]:   2,
	_CodeName[29:42]:   3,
	_CodeName[42:54]:   4,
	_CodeName[54:63]:   5,
	_CodeName[63:82]:   6,
	_CodeName[82:95]:   7,
	_CodeName[95:108]:  8,
	_CodeName[108:124]: 9,
	_CodeName[124:139]: 10,
	_CodeName[139:146]: 11,
}

// ParseCode attempts to convert a string to a Code
func ParseCode(name string) (Code, error) {
	if x, ok := _CodeValue[name]; ok {
		return x, nil
	}
	return Code(0), fmt.Errorf("%s is not a valid Code", name)
}

// MarshalText implements the text marshaller method
func (x Code) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *Code) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseCode(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
