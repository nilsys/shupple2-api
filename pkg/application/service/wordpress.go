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
		PatchPost(*entity.Post, *wordpress.Post) error
		PatchTouristSpot(*entity.TouristSpot, *wordpress.Location) error
		PatchCategory(*entity.Category, *wordpress.Category) error
		PatchLcategory(*entity.Lcategory, *wordpress.LocationCategory) error
		PatchVlog(*entity.Vlog, *wordpress.Vlog) error
		PatchFeature(*entity.Feature, *wordpress.Feature) error
		PatchComic(*entity.Comic, *wordpress.Comic) error
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

func (s *WordpressServiceImpl) PatchPost(post *entity.Post, wpPost *wordpress.Post) error {
	user, err := s.UserQueryRepository.FindByWordpressID(wpPost.Author)
	if err != nil {
		return errors.Wrap(err, "failed to get user corresponding to wordpress user")
	}

	thumbnail, err := s.getThumbnail(wpPost.FeaturedMedia)
	if err != nil {
		return errors.Wrap(err, "failed to get thumbnail")
	}

	wpPostTags, err := s.WordpressQueryRepository.FindPostTagsByIDs(wpPost.Tags)
	if err != nil {
		return errors.Wrap(err, "failed to get post tags")
	}
	hashtags := make([]string, len(wpPostTags))
	for i, wpPostTag := range wpPostTags {
		hashtags[i] = wpPostTag.Name
	}

	toc, err := s.extractTOC(wpPost)
	if err != nil {
		return errors.Wrap(err, "failed to extract toc")
	}
	bodies := strings.Split(wpPost.Content.Rendered, wpPageDelimiter)

	hashtagEntieis, err := s.HashtagCommandService.FindOrCreateHashtags(hashtags)
	if err != nil {
		return errors.Wrap(err, "failed to find hashtags")
	}
	hashtagIDs := make([]int, len(hashtagEntieis))
	for i, hashtagEntity := range hashtagEntieis {
		hashtagIDs[i] = hashtagEntity.ID
	}

	post.ID = wpPost.ID
	post.UserID = user.ID
	post.Slug = wpPost.Slug
	post.Thumbnail = thumbnail
	post.Title = wpPost.Title.Rendered
	post.TOC = toc
	post.IsSticky = wpPost.Sticky
	post.HideAds = wpPost.Meta.IsAdsRemovedInPage
	post.SEOTitle = wpPost.Meta.SEOTitle
	post.SEODescription = wpPost.Meta.MetaDescription
	post.CreatedAt = time.Time(wpPost.Date)
	post.SetBodies(bodies)
	post.SetCategories(wpPost.Categories)
	post.SetHashtags(hashtagIDs)

	return nil
}

func (s *WordpressServiceImpl) PatchTouristSpot(touristSpot *entity.TouristSpot, wpLocation *wordpress.Location) error {
	lat, err := strconv.ParseFloat(wpLocation.Attributes.Map.Lat, 64)
	if err != nil {
		return errors.Wrap(err, "failed to parse lat")
	}

	lng, err := strconv.ParseFloat(wpLocation.Attributes.Map.Lng, 64)
	if err != nil {
		return errors.Wrap(err, "failed to parse lng")
	}

	var thumbnail string
	if wpLocation.FeaturedMedia != 0 {
		thumbnail, err = s.getThumbnail(wpLocation.FeaturedMedia)
		if err != nil {
			return errors.Wrap(err, "failed to get thumbnail")
		}
	}

	touristSpot.ID = wpLocation.ID
	touristSpot.Name = wpLocation.Title.Rendered
	touristSpot.Slug = wpLocation.Slug
	touristSpot.Thumbnail = thumbnail
	touristSpot.WebsiteURL = wpLocation.Attributes.OfficialURL
	touristSpot.City = wpLocation.Attributes.City
	touristSpot.Address = wpLocation.Attributes.Address
	touristSpot.Lat = lat
	touristSpot.Lng = lng
	touristSpot.AccessCar = wpLocation.Attributes.AccessCar
	touristSpot.AccessTrain = wpLocation.Attributes.AccessTrain
	touristSpot.AccessBus = wpLocation.Attributes.AccessBus
	touristSpot.OpeningHours = wpLocation.Attributes.OpeningHours
	touristSpot.TEL = wpLocation.Attributes.TEL
	touristSpot.Price = wpLocation.Attributes.Price
	touristSpot.InstagramURL = wpLocation.Attributes.Instagram
	touristSpot.SearchInnURL = wpLocation.Attributes.Inn
	touristSpot.CreatedAt = time.Time(wpLocation.Date)
	touristSpot.SetCategories(wpLocation.Categories)
	touristSpot.SetLcategories(wpLocation.LocationCat)

	return nil
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
func (s *WordpressServiceImpl) PatchCategory(category *entity.Category, wpCategory *wordpress.Category) error {
	categoryType, err := s.convertCategoryType(wpCategory)
	if err != nil {
		return errors.Wrap(err, "failed to convert category type")
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

	category.ID = wpCategory.ID
	category.Name = wpCategory.Name
	category.Slug = wpCategory.Slug
	category.Type = categoryType
	category.ParentID = pParentID

	return nil
}

func (s *WordpressServiceImpl) convertCategoryType(wpCategory *wordpress.Category) (model.CategoryType, error) {
	if wpCategory.Parent == 0 {
		if wpCategory.Type == wordpress.CategoryTypeJapan || wpCategory.Type == wordpress.CategoryTypeWorld {
			return model.CategoryTypeArea, nil
		}
		return model.CategoryTypeTheme, nil
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

func (s *WordpressServiceImpl) PatchLcategory(lcategory *entity.Lcategory, wpLocationCategory *wordpress.LocationCategory) error {
	lcategory.ID = wpLocationCategory.ID
	lcategory.Name = wpLocationCategory.Name

	return nil
}

func (s *WordpressServiceImpl) PatchVlog(vlog *entity.Vlog, wpVlog *wordpress.Vlog) error {
	user, err := s.UserQueryRepository.FindByWordpressID(wpVlog.Attributes.Creator.ID)
	if err != nil {
		return errors.Wrap(err, "failed to get user corresponding to wordpress user")
	}

	thumbnail, err := s.getThumbnail(wpVlog.FeaturedMedia)
	if err != nil {
		return errors.Wrap(err, "failed to get thumbnail")
	}

	vlog.ID = wpVlog.ID
	vlog.UserID = user.ID
	vlog.Slug = wpVlog.Slug
	vlog.Thumbnail = thumbnail
	vlog.Title = wpVlog.Title.Rendered
	vlog.Body = wpVlog.Content.Rendered
	vlog.YoutubeURL = wpVlog.Attributes.Youtube
	vlog.Series = wpVlog.Attributes.Series
	vlog.UserSNS = wpVlog.Attributes.CreatorSns
	vlog.EditorName = wpVlog.Attributes.FilmEdit
	vlog.YearMonth = wpVlog.Attributes.YearMonth
	vlog.PlayTime = wpVlog.Attributes.RunTime
	vlog.Timeline = wpVlog.Attributes.MovieTimeline
	vlog.CreatedAt = time.Time(wpVlog.Date)
	vlog.SetTouristSpots(wpVlog.Attributes.MovieLocation)
	vlog.SetCategories(wpVlog.Categories)

	return nil
}

func (s *WordpressServiceImpl) PatchFeature(feature *entity.Feature, wpFeature *wordpress.Feature) error {
	user, err := s.UserQueryRepository.FindByWordpressID(wpFeature.Author)
	if err != nil {
		return errors.Wrap(err, "failed to get user corresponding to wordpress user")
	}

	thumbnail, err := s.getThumbnail(wpFeature.FeaturedMedia)
	if err != nil {
		return errors.Wrap(err, "failed to get thumbnail")
	}

	postIDs := make([]int, len(wpFeature.Attributes.FeatureArticle))
	for i, feature := range wpFeature.Attributes.FeatureArticle {
		postIDs[i] = feature.ID
	}

	feature.ID = wpFeature.ID
	feature.UserID = user.ID
	feature.Slug = wpFeature.Slug
	feature.Thumbnail = thumbnail
	feature.Title = wpFeature.Title.Rendered
	feature.Body = wpFeature.Content.Rendered
	feature.CreatedAt = time.Time(wpFeature.Date)
	feature.SetPosts(postIDs)

	return nil
}

func (s *WordpressServiceImpl) PatchComic(comic *entity.Comic, wpComic *wordpress.Comic) error {
	user, err := s.UserQueryRepository.FindByWordpressID(wpComic.Author)
	if err != nil {
		return errors.Wrap(err, "failed to get user corresponding to wordpress user")
	}

	thumbnail, err := s.getThumbnail(wpComic.FeaturedMedia)
	if err != nil {
		return errors.Wrap(err, "failed to get thumbnail")
	}

	comic.ID = wpComic.ID
	comic.UserID = user.ID
	comic.Slug = wpComic.Slug
	comic.Thumbnail = thumbnail
	comic.Title = wpComic.Title.Rendered
	comic.Body = wpComic.Content.Rendered
	comic.CreatedAt = time.Time(wpComic.Date)

	return nil
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
