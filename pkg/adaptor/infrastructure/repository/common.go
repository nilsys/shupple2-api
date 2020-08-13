package repository

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"

	payjp2 "github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository/payjp"

	"github.com/payjp/payjp-go/v1"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file" // register driver
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // register driver
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
)

const (
	defaultAcquisitionNumber        = 1000
	defaultRangeSearchKm            = 5
	dummyCredential                 = "staywaydummy"
	defaultSearchSuggestionsNumber  = 10
	defaultFollowRecommendUserLimit = 20
)

type (
	DAO struct {
		UnderlyingDB *gorm.DB
	}
)

func (d DAO) DB(c context.Context) *gorm.DB {
	if c == nil {
		return d.UnderlyingDB
	}

	if db, ok := c.Value(model.ContextKeyTransaction).(*gorm.DB); ok {
		return db
	}

	return d.UnderlyingDB
}

func (d DAO) LockDB(c context.Context) *gorm.DB {
	if db, ok := c.Value(model.ContextKeyTransaction).(*gorm.DB); ok {
		return db.Set("gorm:query_option", "FOR UPDATE")
	}

	db := d.UnderlyingDB.New()
	_ = db.AddError(serror.New(nil, serror.CodeUndefined, "try to lock outside transaction"))
	return db
}

var RepositoriesSet = wire.NewSet(
	ProvideDB,
	wire.Struct(new(DAO), "*"),
	AreaCategoryCommandRepositorySet,
	AreaCategoryQueryRepositorySet,
	ThemeCategoryCommandRepositorySet,
	ThemeCategoryQueryRepositorySet,
	ComicCommandRepositorySet,
	ComicQueryRepositorySet,
	ComicFavoriteCommandRepositorySet,
	ComicFavoriteQueryRepositorySet,
	FeatureCommandRepositorySet,
	FeatureQueryRepositorySet,
	SpotCategoryCommandRepositorySet,
	SpotCategoryQueryRepositorySet,
	TouristSpotCommandRepositorySet,
	TouristSpotQueryRepositorySet,
	ShippingQueryRepositorySet,
	ShippingCommandRepositorySet,
	PaymentQueryRepositorySet,
	PaymentCommandRepositorySet,
	payjp2.CustomerQueryRepositorySet,
	payjp2.CustomerCommandRepositorySet,
	CardQueryRepositorySet,
	CfProjectQueryRepositorySet,
	CfProjectCommandRepositorySet,
	CardCommandRepositorySet,
	payjp2.CardCommandRepositorySet,
	payjp2.ChargeCommandRepositorySet,
	PostCommandRepositorySet,
	PostQueryRepositorySet,
	PostFavoriteCommandRepositorySet,
	PostFavoriteQueryRepositorySet,
	UserQueryRepositorySet,
	UserCommandRepositorySet,
	UserSalesHistoryCommandRepositorySet,
	VlogCommandRepositorySet,
	VlogQueryRepositorySet,
	VlogFavoriteCommandRepositorySet,
	VlogFavoriteQueryRepositorySet,
	WordpressQueryRepositorySet,
	CfReturnGiftQueryRepositorySet,
	CfReturnGiftCommandRepositorySet,
	ReviewQueryRepositorySet,
	ReviewCommandRepositorySet,
	ReviewFavoriteQueryRepositorySet,
	ReviewFavoriteCommandRepositorySet,
	InnQueryRepositorySet,
	HashtagCommandRepositorySet,
	HashtagQueryRepositorySet,
	HealthCheckRepositorySet,
	TransactionServiceSet,
	InterestQueryRepositorySet,
	NoticeQueryRepositorySet,
	NoticeCommandRepositorySet,
	ReportCommandRepositorySet,
	ReportQueryRepositorySet,
	SlackRepositorySet,
	ProvideAWSSession,
	ProvidePayjp,
	MetasearchAreaQueryRepositorySet,
	ProvideMailer,
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

func ProvideAWSSession(config *config.Config) (*session.Session, error) {
	cfgs := aws.NewConfig().WithRegion(config.AWS.Region)

	if config.AWS.Endpoint != "" {
		cfgs = cfgs.
			WithEndpoint(config.AWS.Endpoint).
			WithS3ForcePathStyle(true).
			WithCredentials(credentials.NewStaticCredentials(dummyCredential, dummyCredential, ""))
	}

	if config.IsDev() {
		cfgs = cfgs.WithLogLevel(aws.LogDebug)
	}

	return session.NewSession(cfgs)
}

func ProvideS3Uploader(sess *session.Session) *s3manager.Uploader {
	return s3manager.NewUploader(sess)
}

func ProvidePayjp(config *config.Config) *payjp.Service {
	return payjp.New(config.Payjp.SecretKey, nil)
}

func ProvideMailer(config *config.Config, sess *session.Session) repository.MailCommandRepository {
	if config.IsDev() {
		return &MailCommandRepositoryForLocalImpl{}
	}

	return &MailCommandRepositoryImpl{
		AWSSession: sess,
		AWSConfig:  config.AWS,
	}
}

func Transaction(db *gorm.DB, f func(db *gorm.DB) error) (err error) {
	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
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

func ErrorToIsExist(err error, resource string, args ...interface{}) (bool, error) {
	resource = fmt.Sprintf(resource, args...)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return true, errors.Wrapf(err, "failed to get %s", resource)
	}
	return true, nil
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

func MigrateUp(database, migrationsDir string) error {
	source := "file://" + migrationsDir

	// passwordのエスケープ周りで不整合があるので、migrate.Newは使えない
	db, err := sql.Open("mysql", database+"&multiStatements=true")
	if err != nil {
		return errors.Wrap(err, "failed to connect db for migration")
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return errors.Wrap(err, "failed to create migrate driver")
	}

	m, err := migrate.NewWithDatabaseInstance(source, "mysql", driver)
	if err != nil {
		return errors.Wrap(err, "failed to create migration instance")
	}
	defer m.Close()

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			logger.Info("database is up-to-date")
			return nil
		}
		return errors.Wrap(err, "failed to migrate up")
	}

	return nil
}
