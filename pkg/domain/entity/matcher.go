package entity

import (
	"fmt"
	"reflect"
	"time"

	"github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	"github.com/pkg/errors"
)

type equalEntityMatcher struct {
	expected interface{}
	failed   *equalEntityMatcherChildMatcher
}

type equalEntityMatcherChildMatcher struct {
	fieldPath string
	matcher   types.GomegaMatcher
	actual    interface{}
}

func EqualEntity(expected interface{}) types.GomegaMatcher {
	return &equalEntityMatcher{
		expected: expected,
	}
}

func (matcher *equalEntityMatcher) Match(actual interface{}) (success bool, err error) {
	expectedType := reflect.TypeOf(matcher.expected)
	actualType := reflect.TypeOf(actual)

	if expectedType != actualType {
		return false, errors.Errorf("wrong comparison between %s and %s", expectedType, actualType)
	}

	return matcher.comparison("$", reflect.ValueOf(matcher.expected), reflect.ValueOf(actual))
}

func (matcher *equalEntityMatcher) comparison(fieldPath string, expected, actual reflect.Value) (bool, error) {
	switch expected.Kind() {
	case reflect.Struct:
		if !expected.Type().ConvertibleTo(reflect.TypeOf(time.Time{})) {
			return matcher.comparisonStruct(fieldPath, expected, actual)
		}
	case reflect.Array, reflect.Slice:
		return matcher.comparisonSlice(fieldPath, expected, actual)
	case reflect.Map:
		return false, errors.New("comparison of map is not implemented")
	case reflect.Ptr:
		return matcher.comparison(fieldPath, expected.Elem(), actual.Elem())
	}

	equal := gomega.Equal(expected.Interface())
	success, err := equal.Match(actual.Interface())
	if !success {
		matcher.failed = &equalEntityMatcherChildMatcher{
			fieldPath: fieldPath,
			matcher:   equal,
			actual:    actual.Interface(),
		}
	}
	return success, err
}

func (matcher *equalEntityMatcher) comparisonStruct(fieldPath string, expected, actual reflect.Value) (bool, error) {
	expectedType := expected.Type()

	if expectedType == reflect.TypeOf(Times{}) {
		actualTimes := actual.Interface().(Times)
		return matcher.combine([]*equalEntityMatcherChildMatcher{
			{"CreatedAt", gomega.Not(gomega.BeZero()), actualTimes.CreatedAt},
			{"UpdatedAt", gomega.Not(gomega.BeZero()), actualTimes.UpdatedAt},
			{"DeletedAt", gomega.BeNil(), actualTimes.DeletedAt},
		})
	}

	if expectedType == reflect.TypeOf(TimesWithoutDeletedAt{}) {
		actualTimes := actual.Interface().(TimesWithoutDeletedAt)
		return matcher.combine([]*equalEntityMatcherChildMatcher{
			{"CreatedAt", gomega.Not(gomega.BeZero()), actualTimes.CreatedAt},
			{"UpdatedAt", gomega.Not(gomega.BeZero()), actualTimes.UpdatedAt},
		})
	}

	for i := 0; i < expectedType.NumField(); i++ {
		field := expectedType.Field(i)
		childFieldPath := fmt.Sprint(fieldPath, ".", field.Name) // TODO: isAnonymous
		if success, err := matcher.comparison(childFieldPath, expected.Field(i), actual.Field(i)); !success {
			return success, err
		}
	}

	return true, nil
}

func (matcher *equalEntityMatcher) comparisonSlice(fieldPath string, expected, actual reflect.Value) (bool, error) {
	lenMatcher := gomega.HaveLen(expected.Len())
	if success, err := lenMatcher.Match(actual.Interface()); !success {
		matcher.failed = &equalEntityMatcherChildMatcher{
			fieldPath: fieldPath + ".length",
			matcher:   lenMatcher,
			actual:    actual.Interface(),
		}
		return success, err
	}

	for i := 0; i < expected.Len(); i++ {
		childFieldPath := fmt.Sprint(fieldPath, "[", i, "]")
		if success, err := matcher.comparison(childFieldPath, expected.Index(i), actual.Index(i)); !success {
			return success, err
		}
	}

	return true, nil
}

func (matcher *equalEntityMatcher) combine(childMatchers []*equalEntityMatcherChildMatcher) (bool, error) {
	for _, childMatcher := range childMatchers {
		if success, err := childMatcher.matcher.Match(childMatcher.actual); !success {
			matcher.failed = childMatcher
			return success, err
		}
	}

	return true, nil
}

func (matcher *equalEntityMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("<%s> %s", matcher.failed.fieldPath, matcher.failed.matcher.FailureMessage(matcher.failed.actual))
}

func (matcher *equalEntityMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("<%s> %s", matcher.failed.fieldPath, matcher.failed.matcher.NegatedFailureMessage(matcher.failed.actual))
}
