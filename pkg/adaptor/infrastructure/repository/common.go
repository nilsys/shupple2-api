package repository

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // register driver
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
)

func ProvideDB(config *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", config.Database)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect db")
	}

	db.LogMode(config.IsDev())
	db.SingularTable(true)
	db.Callback().Create().Remove("gorm:update_time_stamp")
	db.Callback().Update().Remove("gorm:update_time_stamp")

	origUpdateAssociationsCallback := db.Callback().Update().Get("gorm:save_after_associations")
	db.Callback().Update().
		Replace("gorm:save_after_associations", wrapUpdateAssociationsCallback(origUpdateAssociationsCallback))

	db = db.Set("gorm:auto_preload", true)

	return db, nil
}

func Transaction(db *gorm.DB, f func(db *gorm.DB) error) (err error) {
	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			err = serror.New(err, serror.CodeUndefined, "%s", r)
		}
		if err != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := f(tx); err != nil {
		return err
	}

	return tx.Commit().Error
}

func ErrorToFindSingleRecord(err error, resource string, args ...interface{}) error {
	resource = fmt.Sprintf(resource, args...)
	if gorm.IsRecordNotFoundError(err) {
		return serror.New(err, serror.CodeNotFound, "%s not found", resource)
	}
	return errors.Wrapf(err, "failed to get %s", resource)
}

func wrapUpdateAssociationsCallback(base func(scope *gorm.Scope)) func(scope *gorm.Scope) {
	return func(scope *gorm.Scope) {
		db := scope.DB()
		if db.Error == nil && db.RowsAffected > 0 {
			clearHasMany(scope)
			if !scope.HasError() {
				base(scope)
			}
		}
	}
}

// TODO: many_to_manyとhas_oneはどうするか
func clearHasMany(scope *gorm.Scope) {
	for _, field := range scope.Fields() {
		rel := field.Relationship
		if rel == nil || rel.Kind != "has_many" {
			continue
		}

		db := scope.DB()
		for i, foreignKey := range rel.ForeignDBNames {
			referencedField, ok := scope.FieldByName(rel.AssociationForeignFieldNames[i])
			if !ok {
				return
			}
			db = db.Where(fmt.Sprintf("%v = ?", scope.Quote(foreignKey)), referencedField.Field.Interface())
		}

		if db.Delete(reflect.New(field.Field.Type()).Interface()).Error != nil {
			return
		}
	}
}
