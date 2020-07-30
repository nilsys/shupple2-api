package main

import (
	"fmt"
	"log"
	"path"
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
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
	"go.uber.org/zap"
	"gopkg.in/guregu/null.v3"
)

const (
	debug   = true
	perPage = 100
)

type IDContainer struct {
	ID int
}

type PostEntry struct {
	TargetTable   string
	WordpressName string
	Importer      interface{}
}

type TaxonomyEntry struct {
	Taxonomy string
	Getter   func(parentID int) ([]*IDContainer, error)
	Importer func(id int, termDeleted bool) error
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
	UserService           service.UserCommandService
	CategoryService       service.CategoryCommandService
	ComicService          service.ComicCommandService
	FeatureService        service.FeatureCommandService
	SpotCategoryService   service.SpotCategoryCommandService
	PostService           service.PostCommandService
	TouristSpotService    service.TouristSpotCommandService
	VlogService           service.VlogCommandService
	ReviewCommandScenario scenario.ReviewCommandScenario
}

type CustomizedWordpressRepo struct {
	*repositoryImpl.WordpressQueryRepositoryImpl
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

func (s Script) PostEntries() []PostEntry {
	return []PostEntry{
		PostEntry{"post", "post", s.PostService.ImportFromWordpressByID},
		PostEntry{"feature", "feature", s.FeatureService.ImportFromWordpressByID},
		PostEntry{"tourist_spot", "location", s.TouristSpotService.ImportFromWordpressByID},
		PostEntry{"vlog", "movie", s.VlogService.ImportFromWordpressByID},
		PostEntry{"comic", "comic", s.ComicService.ImportFromWordpressByID},
	}
}

func (s Script) TaxonomyEntries() []TaxonomyEntry {
	repo := &CustomizedWordpressRepo{s.WordpressRepo.(*repositoryImpl.WordpressQueryRepositoryImpl)}
	return []TaxonomyEntry{
		TaxonomyEntry{"category", repo.FindCategoryIDsByParentID, s.CategoryService.ImportFromWordpressByID},
		TaxonomyEntry{"spotCategory", repo.FindSpotCategoryIDsByParentID, s.SpotCategoryService.ImportFromWordpressByID},
	}
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

	for _, e := range s.TaxonomyEntries() {
		log.Printf("start to import %s\n", e.Taxonomy)
		if err := s.importTaxonomy(db, e); err != nil {
			return errors.Wrapf(err, "failed to import %s", e.Taxonomy)
		}
	}

	for _, e := range s.PostEntries() {
		log.Printf("start to import %s\n", e.TargetTable)
		if err := s.importPostData(db, e); err != nil {
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

func queryForGetMaxID(entry PostEntry) string {
	return fmt.Sprintf("select COALESCE(max(id), 0) from %s", entry.TargetTable)
}

func queryForWordpressIDs(entry PostEntry) string {
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
			err := s.UserService.ImportFromWordpressByID(id)
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
		UserTiny: entity.UserTiny{
			ID:          wpUser.ID,
			WordpressID: null.IntFrom(int64(wpUser.ID)),
			Name:        wpUser.Name,
			Profile:     wpUser.Description,
			Birthdate:   time.Date(2100, 10, 10, 0, 0, 0, 0, time.Local),
		},
	}
}

func (s Script) importPostData(wordpressDB *gorm.DB, entry PostEntry) error {
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
			importerResults := rImporter.Call([]reflect.Value{reflect.ValueOf(id), reflect.ValueOf(false)})
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

func (s Script) importTaxonomy(wordpressDB *gorm.DB, entry TaxonomyEntry) error {
	return errors.Wrapf(s.importCategorySub(wordpressDB, entry, 0), "failed to import taxonomies; %s", entry.Taxonomy)
}

func (s Script) importCategorySub(wordpressDB *gorm.DB, entry TaxonomyEntry, parentID int) error {
	// 一つのカテゴリに紐づく子カテの数はそれぞれ100件未満であることを事前に確認しているためページングは不要
	ids, err := entry.Getter(parentID)
	if err != nil {
		return errors.Wrapf(err, "failed to find %s", entry.Taxonomy)
	}
	if len(ids) == 100 {
		return errors.Errorf("script seems to be bugged for %s", entry.Taxonomy)
	}

	for _, id := range ids {
		if err := entry.Importer(id.ID, false); err != nil {
			return errors.Wrapf(err, "failed to import wordpress %s; id=%d", entry.Taxonomy, id.ID)
		}

		if debug {
			continue
		}

		if err := s.importCategorySub(wordpressDB, entry, id.ID); err != nil {
			return errors.Wrapf(err, "failed to import wordpress %s children; parent_id=%d", entry.Taxonomy, id.ID)
		}
	}

	return nil
}

func (r *CustomizedWordpressRepo) FindCategoryIDsByParentID(parentID int) ([]*IDContainer, error) {
	return r.FindIDsByParentID("/wp-json/wp/v2/categories/", parentID)
}

func (r *CustomizedWordpressRepo) FindSpotCategoryIDsByParentID(parentID int) ([]*IDContainer, error) {
	return r.FindIDsByParentID("/wp-json/wp/v2/location__cat/", parentID)
}

func (r *CustomizedWordpressRepo) FindIDsByParentID(listPath string, parentID int) ([]*IDContainer, error) {
	wURL := r.BaseURL
	wURL.Path = path.Join(wURL.Path, listPath)

	q := wURL.Query()
	q.Set("parent", fmt.Sprint(parentID))
	q.Set("per_page", fmt.Sprint(perPage))
	wURL.RawQuery = q.Encode()

	var resp []*IDContainer
	return resp, errors.Wrap(r.GetJSON(wURL.String(), &resp), "failed to get wordpress category")
}
