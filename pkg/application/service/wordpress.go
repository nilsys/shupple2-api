package service

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity/wordpress"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
	"go.uber.org/zap"
	"gopkg.in/guregu/null.v3"
)

const (
	wpPageDelimiter   = "<!--nextpage-->"
	thumbnailS3Prefix = "https://s3-ap-northeast-1.amazonaws.com/"
)

type (
	WordpressService interface {
		PatchPost(*entity.Post, *wordpress.Post) error
		PatchTouristSpot(*entity.TouristSpot, *wordpress.Location) error
		PatchAreaCategory(*entity.AreaCategory, *wordpress.Category) error
		PatchThemeCategory(*entity.ThemeCategory, *wordpress.Category) error
		PatchSpotCategory(*entity.SpotCategory, *wordpress.LocationCategory) error
		PatchVlog(*entity.Vlog, *wordpress.Vlog) error
		PatchFeature(*entity.Feature, *wordpress.Feature) error
		PatchComic(*entity.Comic, *wordpress.Comic) error
		NewCfProject(*wordpress.CfProject) (*entity.CfProject, error)
		NewCfReturnGift(*wordpress.CfReturnGift) (*entity.CfReturnGift, error)
	}

	WordpressServiceImpl struct {
		repository.WordpressQueryRepository
		repository.UserQueryRepository
		repository.AreaCategoryQueryRepository
		repository.ThemeCategoryQueryRepository
		repository.SpotCategoryQueryRepository
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

	areaCategoryIDs, themeCategoryIDs, err := s.splitCategories(wpPost.Categories)
	if err != nil {
		return errors.Wrap(err, "failed to split post categories")
	}

	var cfProjectID null.Int
	if wpPost.Attributes.CfProject != nil {
		cfProjectID = null.IntFrom(int64(wpPost.Attributes.CfProject.ID))
	}

	post.ID = wpPost.ID
	post.UserID = user.ID
	post.Slug = string(wpPost.Slug)
	post.Thumbnail = thumbnail
	post.Title = wpPost.Title.Rendered
	post.TOC = toc
	post.CfProjectID = cfProjectID
	post.IsSticky = wpPost.Sticky
	post.HideAds = wpPost.Meta.IsAdsRemovedInPage
	post.SEOTitle = wpPost.Meta.SEOTitle
	post.SEODescription = wpPost.Meta.MetaDescription
	post.EditedAt = time.Time(wpPost.Modified)
	post.CreatedAt = time.Time(wpPost.Date)
	post.SetBodies(bodies)
	post.SetAreaCategories(areaCategoryIDs)
	post.SetThemeCategories(themeCategoryIDs)
	post.SetHashtags(hashtagIDs)

	return nil
}

func (s *WordpressServiceImpl) PatchTouristSpot(touristSpot *entity.TouristSpot, wpLocation *wordpress.Location) error {
	var (
		lat null.Float
		lng null.Float
	)
	latFloat, lngFloat, err := wpLocation.Attributes.LatLang()
	if err != nil {
		err = errors.Wrap(err, "invalid map location")
		logger.Warn(err.Error(), zap.Error(err))
	} else {
		lat = null.FloatFrom(latFloat)
		lng = null.FloatFrom(lngFloat)
	}

	var thumbnail string
	if wpLocation.FeaturedMedia != 0 {
		thumbnail, err = s.getThumbnail(wpLocation.FeaturedMedia)
		if err != nil {
			return errors.Wrap(err, "failed to get thumbnail")
		}
	}

	areaCategoryIDs, themeCategoryIDs, err := s.splitCategories(wpLocation.Categories)
	if err != nil {
		return errors.Wrap(err, "failed to split tourist spot categories")
	}

	touristSpot.ID = wpLocation.ID
	touristSpot.Name = wpLocation.Title.Rendered
	touristSpot.Slug = string(wpLocation.Slug)
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
	touristSpot.EditedAt = time.Time(wpLocation.Modified)
	touristSpot.CreatedAt = time.Time(wpLocation.Date)
	touristSpot.SetAreaCategories(areaCategoryIDs)
	touristSpot.SetThemeCategories(themeCategoryIDs)
	touristSpot.SetSpotCategories(wpLocation.LocationCat)

	return nil
}

// TODO: 親カテのカテゴリタイプが影響するので、カテゴリの更新があった場合は子カテの更新も行わないといけない
func (s *WordpressServiceImpl) PatchAreaCategory(category *entity.AreaCategory, wpCategory *wordpress.Category) error {
	category.ID = wpCategory.ID
	category.Name = wpCategory.Name
	category.Slug = string(wpCategory.Slug)

	if wpCategory.Parent != 0 {
		parent, err := s.AreaCategoryQueryRepository.FindByID(wpCategory.Parent)
		if err != nil {
			return errors.Wrap(err, "failed to find parent area category")
		}

		category.AreaGroup = parent.AreaGroup
		category.AreaID = parent.AreaID
		switch parent.Type {
		case model.AreaCategoryTypeArea:
			category.Type = model.AreaCategoryTypeSubArea
			category.SubAreaID = null.IntFrom(int64(category.ID))
		case model.AreaCategoryTypeSubArea:
			category.Type = model.AreaCategoryTypeSubSubArea
			category.SubAreaID = parent.SubAreaID
			category.SubSubAreaID = null.IntFrom(int64(category.ID))
		case model.AreaCategoryTypeSubSubArea, model.AreaCategoryTypeUndefined:
			logger.Warn("parent area category is sub_sub_area or undefined", zap.Int("id", category.ID), zap.Int("parent", wpCategory.Parent))
			category.Type = model.AreaCategoryTypeUndefined
			category.SubAreaID = parent.SubAreaID
			category.SubSubAreaID = parent.SubSubAreaID
		}

		return nil
	}

	category.Type = model.AreaCategoryTypeArea
	category.AreaID = category.ID
	switch wpCategory.Type {
	case wordpress.CategoryTypeJapan:
		category.AreaGroup = model.AreaGroupJapan
	case wordpress.CategoryTypeWorld:
		category.AreaGroup = model.AreaGroupWorld
	default:
		return serror.New(nil, serror.CodeUndefined, "invalid area group")
	}

	return nil
}

func (s *WordpressServiceImpl) PatchThemeCategory(category *entity.ThemeCategory, wpCategory *wordpress.Category) error {
	category.ID = wpCategory.ID
	category.Name = wpCategory.Name
	category.Slug = string(wpCategory.Slug)

	if wpCategory.Parent != 0 {
		parent, err := s.ThemeCategoryQueryRepository.FindByID(wpCategory.Parent)
		if err != nil {
			return errors.Wrap(err, "failed to find parent theme category")
		}

		switch parent.Type {
		case model.ThemeCategoryTypeTheme:
			category.Type = model.ThemeCategoryTypeSubTheme
			category.ThemeID = parent.ThemeID
			category.SubThemeID = null.IntFrom(int64(category.ID))
		case model.ThemeCategoryTypeSubTheme, model.ThemeCategoryTypeUndefined:
			logger.Warn("parent theme category is sub_theme or undefined", zap.Int("id", category.ID), zap.Int("parent", wpCategory.Parent))
			category.Type = model.ThemeCategoryTypeUndefined
			category.ThemeID = parent.ThemeID
			category.SubThemeID = parent.SubThemeID
		}

		return nil
	}

	category.Type = model.ThemeCategoryTypeTheme
	category.ThemeID = category.ID

	return nil
}

func (s *WordpressServiceImpl) PatchSpotCategory(spotCategory *entity.SpotCategory, wpLocationCategory *wordpress.LocationCategory) error {
	spotCategory.ID = wpLocationCategory.ID
	spotCategory.Name = wpLocationCategory.Name
	spotCategory.Slug = string(wpLocationCategory.Slug)

	if wpLocationCategory.Parent != 0 {
		parent, err := s.SpotCategoryQueryRepository.FindByID(wpLocationCategory.Parent)
		if err != nil {
			return errors.Wrap(err, "failed to find parent spotCategory spotCategory")
		}

		switch parent.Type {
		case model.SpotCategoryTypeSpotCategory:
			spotCategory.Type = model.SpotCategoryTypeSubSpotCategory
			spotCategory.SpotCategoryID = parent.SpotCategoryID
			spotCategory.SubSpotCategoryID = null.IntFrom(int64(spotCategory.ID))
		case model.SpotCategoryTypeSubSpotCategory, model.SpotCategoryTypeUndefined:
			logger.Warn("parent spotCategory spotCategory is sub_spotCategory or undefined", zap.Int("id", spotCategory.ID), zap.Int("parent", wpLocationCategory.Parent))
			spotCategory.Type = model.SpotCategoryTypeUndefined
			spotCategory.SpotCategoryID = parent.SpotCategoryID
			spotCategory.SubSpotCategoryID = parent.SubSpotCategoryID
		}

		return nil
	}

	spotCategory.Type = model.SpotCategoryTypeSpotCategory
	spotCategory.SpotCategoryID = spotCategory.ID

	return nil
}

func (s *WordpressServiceImpl) PatchVlog(vlog *entity.Vlog, wpVlog *wordpress.Vlog) error {
	user, err := s.UserQueryRepository.FindByWordpressID(wpVlog.Author)
	if err != nil {
		return errors.Wrap(err, "failed to get user corresponding to wordpress user")
	}

	editors := make([]int, len(wpVlog.Attributes.FilmEdit))
	for i, fe := range wpVlog.Attributes.FilmEdit {
		user, err := s.UserQueryRepository.FindByWordpressID(fe.ID)
		if err != nil {
			return errors.Wrap(err, "failed to get user corresponding to film_edit")
		}
		editors[i] = user.ID
	}

	thumbnail, err := s.getThumbnail(wpVlog.FeaturedMedia)
	if err != nil {
		return errors.Wrap(err, "failed to get thumbnail")
	}

	areaCategoryIDs, themeCategoryIDs, err := s.splitCategories(wpVlog.Categories)
	if err != nil {
		return errors.Wrap(err, "failed to split post categories")
	}

	vlog.ID = wpVlog.ID
	vlog.UserID = user.ID
	vlog.Slug = string(wpVlog.Slug)
	vlog.Thumbnail = thumbnail
	vlog.Title = wpVlog.Title.Rendered
	vlog.Body = wpVlog.Content.Rendered
	vlog.YoutubeURL = wpVlog.Attributes.Youtube
	vlog.Series = wpVlog.Attributes.Series
	vlog.YearMonth = wpVlog.Attributes.YearMonth
	vlog.PlayTime = wpVlog.Attributes.RunTime
	vlog.Timeline = wpVlog.Attributes.MovieTimeline
	vlog.EditedAt = time.Time(wpVlog.Modified)
	vlog.CreatedAt = time.Time(wpVlog.Date)
	vlog.SetTouristSpots(wpVlog.Attributes.MovieLocation)
	vlog.SetAreaCategories(areaCategoryIDs)
	vlog.SetThemeCategories(themeCategoryIDs)
	vlog.SetEditors(editors)

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
	feature.Slug = string(wpFeature.Slug)
	feature.Thumbnail = thumbnail
	feature.Title = wpFeature.Title.Rendered
	feature.Body = wpFeature.Content.Rendered
	feature.EditedAt = time.Time(wpFeature.Modified)
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
	comic.Slug = string(wpComic.Slug)
	comic.Thumbnail = thumbnail
	comic.Title = wpComic.Title.Rendered
	comic.Body = wpComic.Content.Rendered
	comic.EditedAt = time.Time(wpComic.Modified)
	comic.CreatedAt = time.Time(wpComic.Date)

	return nil
}

func (s *WordpressServiceImpl) NewCfProject(wpCfProject *wordpress.CfProject) (*entity.CfProject, error) {
	user, err := s.UserQueryRepository.FindByWordpressID(wpCfProject.Author)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user corresponding to wordpress user")
	}

	thumbnails := make([]string, 0)
	for _, gallaryItems := range wpCfProject.Attributes.PhotoGallery.Thumbnails {
		for _, gallaryItem := range gallaryItems {
			thumbnails = append(thumbnails, gallaryItem.FullImageURL)
		}
	}

	areaCategoryIDs, themeCategoryIDs, err := s.splitCategories(wpCfProject.Categories)
	if err != nil {
		return nil, errors.Wrap(err, "failed to split post categories")
	}

	var cfProject entity.CfProject

	cfProject.CfProjectTable = entity.CfProjectTable{
		ID:     wpCfProject.ID,
		UserID: user.ID,
	}

	cfProject.Snapshot.CfProjectSnapshotTable = entity.CfProjectSnapshotTable{
		CfProjectID: wpCfProject.ID,
		UserID:      user.ID,
		Title:       wpCfProject.Title.Rendered,
		Summary:     wpCfProject.Attributes.Summary,
		Body:        wpCfProject.Content.Rendered,
		GoalPrice:   int(wpCfProject.Attributes.GoalPrice),
		Deadline:    time.Time(wpCfProject.Attributes.Deadline),
		// IsAttention: wpCfProject.Attributes.IsAttention,
	}
	cfProject.Snapshot.SetThumbnails(thumbnails)
	cfProject.Snapshot.SetAreaCategories(areaCategoryIDs)
	cfProject.Snapshot.SetThemeCategories(themeCategoryIDs)

	return &cfProject, nil
}

func (s *WordpressServiceImpl) NewCfReturnGift(wpCfReturnGift *wordpress.CfReturnGift) (*entity.CfReturnGift, error) {
	thumbnail, err := s.getThumbnail(wpCfReturnGift.FeaturedMedia)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get thumbnail")
	}

	var cfReturnGift entity.CfReturnGift
	cfReturnGift.CfReturnGiftTiny = entity.CfReturnGiftTiny{
		ID:          wpCfReturnGift.ID,
		CfProjectID: wpCfReturnGift.Attributes.CfProject.ID,
		GiftType:    model.CfReturnGiftType(wpCfReturnGift.Attributes.GiftType),
	}

	cfReturnGift.Snapshot = &entity.CfReturnGiftSnapshotTiny{
		CfReturnGiftID: wpCfReturnGift.ID,
		Thumbnail:      thumbnail,
		Body:           wpCfReturnGift.Content.Rendered,
		Price:          int(wpCfReturnGift.Attributes.Price),
		FullAmount:     int(wpCfReturnGift.Attributes.FullAmount),
		DeliveryDate:   wpCfReturnGift.Attributes.DeliveryDate,
		Deadline:       null.Time(wpCfReturnGift.Attributes.Deadline),
		SortOrder:      int(wpCfReturnGift.Attributes.SortOrder),
	}

	return &cfReturnGift, nil
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
	media, err := s.WordpressQueryRepository.FindMediaByID(mediaID)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get thumbnail(id=%d)", mediaID)
	}

	thumbnail := media.SourceURL
	if strings.HasPrefix(thumbnail, thumbnailS3Prefix) {
		thumbnail = "https://" + thumbnail[len(thumbnailS3Prefix):]
	}
	return thumbnail, nil
}

func (s *WordpressServiceImpl) splitCategories(categoryIDs []int) ([]int, []int, error) {
	existingAreaCategories, err := s.AreaCategoryQueryRepository.FindByIDs(categoryIDs)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get area categories")
	}
	existingAreaCategoryIDsSet := make(map[int]struct{}, len(existingAreaCategories))
	for _, areaCategory := range existingAreaCategories {
		existingAreaCategoryIDsSet[areaCategory.ID] = struct{}{}
	}

	areaCategoryIDs := make([]int, 0, len(categoryIDs))
	themeCategoryIDs := make([]int, 0, len(categoryIDs))
	for _, categoryID := range categoryIDs {
		if _, existing := existingAreaCategoryIDsSet[categoryID]; existing {
			areaCategoryIDs = append(areaCategoryIDs, categoryID)
		} else {
			themeCategoryIDs = append(themeCategoryIDs, categoryID)
		}
	}

	return areaCategoryIDs, themeCategoryIDs, nil
}
