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
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

const (
	wpPageDelimiter   = "<p><!--nextpage--></p>"
	thumbnailS3Prefix = "https://s3-ap-northeast-1.amazonaws.com/"
)

type (
	WordpressService interface {
		ConvertPost(*wordpress.Post) (*entity.Post, error)
		ConvertLocation(*wordpress.Location) (*entity.TouristSpot, error)
		ConvertCategory(*wordpress.Category) (*entity.Category, error)
		ConvertLcategory(*wordpress.LocationCategory) *entity.Lcategory
		ConvertVlog(*wordpress.Vlog) (*entity.Vlog, error)
		ConvertFeature(*wordpress.Feature) (*entity.Feature, error)
		ConvertComic(*wordpress.Comic) (*entity.Comic, error)
	}

	WordpressServiceImpl struct {
		repository.WordpressQueryRepository
		repository.UserQueryRepository
		repository.CategoryQueryRepository
		HashtagCommandService
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

	thumbnail, err := s.getThumbnail(wpPost.FeaturedMedia)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get thumbnail")
	}

	wpPostTags, err := s.WordpressQueryRepository.FindPostTagsByIDs(wpPost.Tags)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get post tags")
	}
	hashtags := make([]string, len(wpPostTags))
	for i, wpPostTag := range wpPostTags {
		hashtags[i] = wpPostTag.Name
	}

	toc, err := s.extractTOC(wpPost)
	if err != nil {
		return nil, errors.Wrap(err, "failed to extract toc")
	}
	bodies := strings.Split(wpPost.Content.Rendered, wpPageDelimiter)

	hashtagEntieis, err := s.HashtagCommandService.FindOrCreateHashtags(hashtags)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find hashtags")
	}
	hashtagIDs := make([]int, len(hashtagEntieis))
	for i, hashtagEntity := range hashtagEntieis {
		hashtagIDs[i] = hashtagEntity.ID
	}

	post := entity.NewPost(entity.PostTiny{
		ID:             wpPost.ID,
		UserID:         user.ID,
		Slug:           wpPost.Slug,
		Thumbnail:      thumbnail,
		Title:          wpPost.Title.Rendered,
		TOC:            toc,
		IsSticky:       wpPost.Sticky,
		HideAds:        wpPost.Meta.IsAdsRemovedInPage,
		SEOTitle:       wpPost.Meta.SEOTitle,
		SEODescription: wpPost.Meta.MetaDescription,
		CreatedAt:      time.Time(wpPost.Date),
		UpdatedAt:      time.Time(wpPost.Modified),
	}, bodies, wpPost.Categories, hashtagIDs)

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

	var thumbnail string
	if wpLocation.FeaturedMedia != 0 {
		thumbnail, err = s.getThumbnail(wpLocation.FeaturedMedia)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get thumbnail")
		}
	}

	touristSpot := entity.NewTouristSpot(entity.TouristSpotTiny{
		ID:           wpLocation.ID,
		Name:         wpLocation.Title.Rendered,
		Slug:         wpLocation.Slug,
		Thumbnail:    thumbnail,
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
		CreatedAt:    time.Time(wpLocation.Date),
		UpdatedAt:    time.Time(wpLocation.Modified),
	}, wpLocation.Categories, wpLocation.LocationCat)

	return &touristSpot, nil
}

/*
## 親カテについて

AreaGroupをwordpress側で作れないので、CategoryTypeがjapan,worldの場合は固定値を入れる

## カテゴリタイプについて

* ルートカテゴリかつCategoryTypeがjapan or worldの場合     -> Area
* ルートカテゴリかつCategoryTypeがjapan or worldでない場合 -> Theme
* 非ルートカテゴリかつ親カテがArea                         -> SubArea
* 非ルートカテゴリかつ親カテがSubArea                      -> SubSubArea
* 非ルートカテゴリかつ親カテがSubSubArea                   -> SubSubArea
* 非ルートカテゴリかつ親カテがTheme                        -> Theme

## TODO:

親カテのカテゴリタイプが影響するので、カテゴリの更新があった場合は子カテの更新も行わないといけない
*/
func (s *WordpressServiceImpl) ConvertCategory(wpCategory *wordpress.Category) (*entity.Category, error) {
	categoryType, err := s.convertCategoryType(wpCategory)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert category type")
	}

	parentID := wpCategory.Parent
	switch wpCategory.Type {
	case wordpress.CategoryTypeJapan:
		parentID = entity.AreaGroupIDJapan
	case wordpress.CategoryTypeWorld:
		parentID = entity.AreaGroupIDWorld
	}

	var pParentID *int
	if parentID > 0 {
		pParentID = &parentID
	}

	category := &entity.Category{
		ID:       wpCategory.ID,
		Name:     wpCategory.Name,
		Slug:     wpCategory.Slug,
		Type:     categoryType,
		ParentID: pParentID,
	}

	return category, nil
}

func (s *WordpressServiceImpl) convertCategoryType(wpCategory *wordpress.Category) (model.CategoryType, error) {
	if wpCategory.Parent == 0 {
		if wpCategory.Type == wordpress.CategoryTypeJapan || wpCategory.Type == wordpress.CategoryTypeWorld {
			return model.CategoryTypeArea, nil
		} else {
			return model.CategoryTypeTheme, nil
		}
	}

	parent, err := s.CategoryQueryRepository.FindByID(wpCategory.Parent)
	if err != nil {
		return model.CategoryType(0), errors.Wrap(err, "failed to find parent category")
	}

	switch parent.Type {
	case model.CategoryTypeArea:
		return model.CategoryTypeSubArea, nil
	case model.CategoryTypeSubArea:
		return model.CategoryTypeSubSubArea, nil
	case model.CategoryTypeSubSubArea:
		return model.CategoryTypeSubSubArea, nil
	case model.CategoryTypeTheme:
		return model.CategoryTypeTheme, nil
	}

	return model.CategoryType(0), serror.New(nil, serror.CodeUndefined, "invalid parent category type")
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

	thumbnail, err := s.getThumbnail(wpVlog.FeaturedMedia)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get thumbnail")
	}

	vlog := entity.NewVlog(entity.VlogTiny{
		ID:         wpVlog.ID,
		UserID:     user.ID,
		Slug:       wpVlog.Slug,
		Thumbnail:  thumbnail,
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

	thumbnail, err := s.getThumbnail(wpFeature.FeaturedMedia)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get thumbnail")
	}

	postIDs := make([]int, len(wpFeature.Attributes.FeatureArticle))
	for i, feature := range wpFeature.Attributes.FeatureArticle {
		postIDs[i] = feature.ID
	}

	feature := entity.NewFeature(entity.FeatureTiny{
		ID:        wpFeature.ID,
		UserID:    user.ID,
		Slug:      wpFeature.Slug,
		Thumbnail: thumbnail,
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

	thumbnail, err := s.getThumbnail(wpComic.FeaturedMedia)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get thumbnail")
	}

	comic := entity.Comic{
		ID:        wpComic.ID,
		UserID:    user.ID,
		Slug:      wpComic.Slug,
		Thumbnail: thumbnail,
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

func (s *WordpressServiceImpl) getThumbnail(mediaID int) (string, error) {
	media, err := s.WordpressQueryRepository.FindMediaByIDs([]int{mediaID})
	if err != nil || len(media) == 0 {
		return "", serror.NewResourcesNotFoundError(err, "thumbnail(id=%d)", mediaID)
	}

	thumbnail := media[0].SourceURL
	if strings.HasPrefix(thumbnail, thumbnailS3Prefix) {
		thumbnail = "https://" + thumbnail[len(thumbnailS3Prefix):]
	}
	return thumbnail, nil
}
