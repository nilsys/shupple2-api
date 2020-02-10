package service

import (
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity/wordpress"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

const (
	wpPageDelimiter = "<p><!--nextpage--></p>"
)

type (
	WordpressService interface {
		ConvertPost(*wordpress.Post) (*entity.Post, error)
		ConvertLocation(*wordpress.Location) (*entity.TouristSpot, error)
		ConvertCategory(*wordpress.Category) *entity.Category
		ConvertLcategory(*wordpress.LocationCategory) *entity.Lcategory
		ConvertVlog(*wordpress.Vlog) (*entity.Vlog, error)
		ConvertFeature(*wordpress.Feature) (*entity.Feature, error)
		ConvertComic(*wordpress.Comic) (*entity.Comic, error)
	}

	WordpressServiceImpl struct {
		WordpressQueryRepository repository.WordpressQueryRepository
		UserQueryRepository      repository.UserQueryRepository
	}
)

var WordpressServiceSet = wire.NewSet(
	wire.Struct(new(WordpressServiceImpl), "*"),
	wire.Bind(new(WordpressService), new(*WordpressServiceImpl)),
)

func (s *WordpressServiceImpl) ConvertPost(wpPost *wordpress.Post) (*entity.Post, error) {
	user, err := s.UserQueryRepository.FindByWordpressID(wpPost.Author)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user corresponding to wordpress user")
	}

	toc, err := s.extractTOC(wpPost)
	if err != nil {
		return nil, errors.Wrap(err, "failed to extract toc")
	}
	bodies := strings.Split(wpPost.Content.Rendered, wpPageDelimiter)

	post := entity.NewPost(entity.PostTiny{
		ID:        wpPost.ID,
		UserID:    user.ID,
		Title:     wpPost.Title.Rendered,
		TOC:       toc,
		Slug:      wpPost.Slug,
		CreatedAt: time.Time(wpPost.Date),
		UpdatedAt: time.Time(wpPost.Modified),
	}, bodies, wpPost.Categories)

	return &post, nil
}

func (s *WordpressServiceImpl) ConvertLocation(wpLocation *wordpress.Location) (*entity.TouristSpot, error) {
	lat, err := strconv.ParseFloat(wpLocation.Attributes.Map.Lat, 64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse lat")
	}

	lng, err := strconv.ParseFloat(wpLocation.Attributes.Map.Lng, 64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse lng")
	}

	touristSpot := entity.NewTouristSpot(entity.TouristSpotTiny{
		ID:           wpLocation.ID,
		Name:         wpLocation.Title.Rendered,
		WebsiteURL:   wpLocation.Attributes.OfficialURL,
		City:         wpLocation.Attributes.City,
		Address:      wpLocation.Attributes.Address,
		Lat:          lat,
		Lng:          lng,
		AccessCar:    wpLocation.Attributes.AccessCar,
		AccessTrain:  wpLocation.Attributes.AccessTrain,
		AccessBus:    wpLocation.Attributes.AccessBus,
		OpeningHours: wpLocation.Attributes.OpeningHours,
		TEL:          wpLocation.Attributes.TEL,
		Price:        wpLocation.Attributes.Price,
		InstagramURL: wpLocation.Attributes.Instagram,
		SearchInnURL: wpLocation.Attributes.Inn,
	}, wpLocation.LocationCat, wpLocation.Categories)

	return &touristSpot, nil
}

func (s *WordpressServiceImpl) ConvertCategory(wpCategory *wordpress.Category) *entity.Category {
	category := &entity.Category{
		ID:   wpCategory.ID,
		Name: wpCategory.Name,
	}

	return category
}

func (s *WordpressServiceImpl) ConvertLcategory(wpLocationCategory *wordpress.LocationCategory) *entity.Lcategory {
	touristSpotCategory := &entity.Lcategory{
		ID:   wpLocationCategory.ID,
		Name: wpLocationCategory.Name,
	}

	return touristSpotCategory
}

func (s *WordpressServiceImpl) ConvertVlog(wpVlog *wordpress.Vlog) (*entity.Vlog, error) {
	user, err := s.UserQueryRepository.FindByWordpressID(wpVlog.Attributes.Creator.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user corresponding to wordpress user")
	}

	vlog := entity.NewVlog(entity.VlogTiny{
		ID:         wpVlog.ID,
		UserID:     user.ID,
		Slug:       wpVlog.Slug,
		Title:      wpVlog.Title.Rendered,
		Body:       wpVlog.Content.Rendered,
		YoutubeURL: wpVlog.Attributes.Youtube,
		Series:     wpVlog.Attributes.Series,
		UserSNS:    wpVlog.Attributes.CreatorSns,
		EditorName: wpVlog.Attributes.FilmEdit,
		YearMonth:  wpVlog.Attributes.YearMonth,
		PlayTime:   wpVlog.Attributes.RunTime,
		Timeline:   wpVlog.Attributes.MovieTimeline,
		CreatedAt:  time.Time(wpVlog.Date),
		UpdatedAt:  time.Time(wpVlog.Modified),
	}, wpVlog.Categories, wpVlog.Attributes.MovieLocation)

	return &vlog, nil
}

func (s *WordpressServiceImpl) ConvertFeature(wpFeature *wordpress.Feature) (*entity.Feature, error) {
	user, err := s.UserQueryRepository.FindByWordpressID(wpFeature.Author)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user corresponding to wordpress user")
	}

	postIDs := make([]int, len(wpFeature.Attributes.FeatureArticle))
	for i, feature := range wpFeature.Attributes.FeatureArticle {
		postIDs[i] = feature.ID
	}

	feature := entity.NewFeature(entity.FeatureTiny{
		ID:        wpFeature.ID,
		UserID:    user.ID,
		Slug:      wpFeature.Slug,
		Title:     wpFeature.Title.Rendered,
		Body:      wpFeature.Content.Rendered,
		CreatedAt: time.Time(wpFeature.Date),
		UpdatedAt: time.Time(wpFeature.Modified),
	}, postIDs)

	return &feature, nil
}

func (s *WordpressServiceImpl) ConvertComic(wpComic *wordpress.Comic) (*entity.Comic, error) {
	user, err := s.UserQueryRepository.FindByWordpressID(wpComic.Author)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user corresponding to wordpress user")
	}

	comic := entity.Comic{
		ID:        wpComic.ID,
		UserID:    user.ID,
		Slug:      wpComic.Slug,
		Title:     wpComic.Title.Rendered,
		Body:      wpComic.Content.Rendered,
		CreatedAt: time.Time(wpComic.Date),
		UpdatedAt: time.Time(wpComic.Modified),
	}

	return &comic, nil
}

func (s *WordpressServiceImpl) extractTOC(wpPost *wordpress.Post) (string, error) {
	d, err := goquery.NewDocumentFromReader(strings.NewReader(wpPost.Content.Rendered))
	if err != nil {
		return "", errors.Wrap(err, "failed to parse content html")
	}

	toc, err := goquery.OuterHtml(d.Find("#toc_container"))
	return toc, errors.Wrap(err, "failed to find toc")
}
