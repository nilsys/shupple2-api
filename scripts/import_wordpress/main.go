package main

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	repositoryImpl "github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"
	"github.com/stayway-corp/stayway-media-api/pkg/application/scenario"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity/wordpress"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
	"go.uber.org/zap"
)

const (
	debug   = false
	perPage = 100
)

type Entry struct {
	IsTaxonomy    bool
	TargetTable   string
	WordpressName string
	Getter        interface{}
	Importer      interface{}
}

type Config struct {
	ImportWordpress struct {
		WordpressDB string `yaml:"wordpress_db"`
	} `yaml:"import_wordpress"`
}

type Script struct {
	DB                    *gorm.DB
	Config                *config.Config
	MediaUploader         *s3manager.Uploader
	AWSConfig             config.AWS
	WordpressRepo         repository.WordpressQueryRepository
	UserQueryRepository   repository.UserQueryRepository
	UserRepo              repository.UserCommandRepository
	CategoryCommandRepo   repository.CategoryCommandRepository
	UserService           service.UserCommandService
	CategoryService       service.CategoryCommandService
	ComicService          service.ComicCommandService
	FeatureService        service.FeatureCommandService
	LcategoryService      service.LcategoryCommandService
	PostService           service.PostCommandService
	TouristSpotService    service.TouristSpotCommandService
	VlogService           service.VlogCommandService
	ReviewCommandScenario scenario.ReviewCommandScenario
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	script, err := InitializeScript("./config.yaml")
	if err != nil {
		return errors.Wrap(err, "failed to initialize script")
	}

	return script.Run()
}

func (s Script) Entries() []Entry {
	return []Entry{
		Entry{true, "lcategory", "location__cat", s.WordpressRepo.FindLocationCategoriesByIDs, s.LcategoryService.ImportFromWordpressByID},
		Entry{false, "post", "post", s.WordpressRepo.FindPostsByIDs, s.PostService.ImportFromWordpressByID},
		Entry{false, "feature", "feature", s.WordpressRepo.FindFeaturesByIDs, s.FeatureService.ImportFromWordpressByID},
		Entry{false, "tourist_spot", "location", s.WordpressRepo.FindLocationsByIDs, s.TouristSpotService.ImportFromWordpressByID},
		Entry{false, "vlog", "movie", s.WordpressRepo.FindVlogsByIDs, s.VlogService.ImportFromWordpressByID},
		Entry{false, "comic", "comic", s.WordpressRepo.FindComicsByIDs, s.ComicService.ImportFromWordpressByID},
	}
}

func (s Script) getWordpressDBURL() string {
	return ""
}

func (s Script) Run() error {
	var config Config
	if err := s.Config.Scripts.Decode(&config); err != nil {
		return errors.Wrap(err, "failed to load script config")
	}

	db, err := connectWordpressDatabase(config.ImportWordpress.WordpressDB)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := s.importUser(db); err != nil {
		return errors.WithStack(err)
	}

	if err := s.importCategory(db); err != nil {
		return errors.WithStack(err)
	}

	for _, e := range s.Entries() {
		log.Printf("start to import %s\n", e.TargetTable)
		if err := s.importData(db, e); err != nil {
			return errors.Wrapf(err, "failed to import %s", e.TargetTable)
		}
	}

	return s.importReview(db)
}

func connectWordpressDatabase(db string) (*gorm.DB, error) {
	return repositoryImpl.ProvideDB(&config.Config{
		Database:    db,
		Development: &config.Development{CognitoID: "dummy"},
	})
}

func queryForGetMaxID(entry Entry) string {
	return fmt.Sprintf("select COALESCE(max(id), 0) from %s", entry.TargetTable)
}

func queryForWordpressIDs(entry Entry) string {
	if entry.IsTaxonomy {
		return "select term_id as id from wp_term_taxonomy where taxonomy = ? and term_id > ? order by term_id limit ?"
	}

	return "select id from wp_posts where post_status = 'publish' AND post_type = ? and id > ? order by id limit ?"
}

func (s Script) importUser(wordpressDB *gorm.DB) error {
	lastID := 0
	if err := s.DB.Raw("select COALESCE(max(wordpress_id), 0) from user").Row().Scan(&lastID); err != nil {
		return errors.WithStack(err)
	}

	for {
		ids := make([]int, 0, perPage)

		q := wordpressDB.Raw("select id from wp_users where id > ? order by id limit ?", lastID, perPage)
		if err := q.Pluck("id", &ids).Error; err != nil {
			return errors.WithStack(err)
		}

		if len(ids) == 0 {
			break
		}

		for _, id := range ids {
			err := s.UserService.RegisterWordpressUser(id)
			if err != nil {
				return errors.WithStack(err)
			}
		}
		lastID = ids[len(ids)-1]

		if debug {
			break
		}
	}

	return nil
}

func (s Script) convertUser(wpUser *wordpress.User) *entity.User {
	return &entity.User{
		ID:          wpUser.ID,
		WordpressID: wpUser.ID,
		Name:        wpUser.Name,
		Profile:     wpUser.Description,
		Birthdate:   time.Date(2100, 10, 10, 0, 0, 0, 0, time.Local),
	}
}

func (s Script) importData(wordpressDB *gorm.DB, entry Entry) error {
	lastID := 0
	if err := s.DB.Raw(queryForGetMaxID(entry)).Row().Scan(&lastID); err != nil {
		return errors.WithStack(err)
	}

	for {
		ids := make([]int, 0, perPage)

		q := wordpressDB.Raw(queryForWordpressIDs(entry), entry.WordpressName, lastID, perPage)
		if err := q.Pluck("id", &ids).Error; err != nil {
			return errors.WithStack(err)
		}

		if len(ids) == 0 {
			break
		}

		for _, id := range ids {
			rImporter := reflect.ValueOf(entry.Importer)
			importerResults := rImporter.Call([]reflect.Value{reflect.ValueOf(id)})
			err := importerResults[1]
			if !err.IsNil() {
				logger.Error(fmt.Sprintf("failed to import; id = %d ", id), zap.Error(err.Interface().(error)))
			}
		}
		lastID = ids[len(ids)-1]

		if debug {
			break
		}
	}

	return nil
}

func (s Script) importCategory(wordpressDB *gorm.DB) error {
	japan := entity.Category{
		ID:   entity.AreaGroupIDJapan,
		Name: "国内",
		Type: model.CategoryTypeAreaGroup,
	}
	if err := s.CategoryCommandRepo.Store(context.Background(), &japan); err != nil {
		return errors.Wrap(err, "failed to save japan category")
	}

	world := entity.Category{
		ID:   entity.AreaGroupIDWorld,
		Name: "海外",
		Type: model.CategoryTypeAreaGroup,
	}
	if err := s.CategoryCommandRepo.Store(context.Background(), &world); err != nil {
		return errors.Wrap(err, "failed to save world category")
	}

	return errors.Wrap(s.importCategorySub(wordpressDB, 0), "failed to import categories")
}

func (s Script) importCategorySub(wordpressDB *gorm.DB, parentID int) error {
	// 一つのカテゴリに紐づく子カテの数はそれぞれ100件未満であることを事前に確認しているためページングは不要
	categories, err := s.WordpressRepo.FindCategoriesByParentID(parentID, 0)
	if err != nil {
		return errors.Wrap(err, "failed to find categories")
	}
	if len(categories) == 100 {
		return errors.New("script seems to be bugged")
	}

	for _, c := range categories {
		if _, err := s.CategoryService.ImportFromWordpressByID(c.ID); err != nil {
			return errors.Wrapf(err, "failed to import wordpress category; id=%d", c.ID)
		}

		if debug {
			continue
		}

		if err := s.importCategorySub(wordpressDB, c.ID); err != nil {
			return errors.Wrapf(err, "failed to import wordpress category children; parent_id=%d", c.ID)
		}
	}

	return nil
}
