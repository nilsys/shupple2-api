package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity/wordpress"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
)

// WordpressServiceに生やすとserviceの相互依存が発生し分かりにくそうなのでsreviceを分けた
type (
	WordpressCallbackService interface {
		Import(entityType wordpress.EntityType, id int) error
	}

	WordpressCallbackServiceImpl struct {
		UserCommandService
		CategoryCommandService
		ComicCommandService
		FeatureCommandService
		SpotCategoryCommandService
		PostCommandService
		TouristSpotCommandService
		VlogCommandService
	}
)

var WordpressCallbackServiceSet = wire.NewSet(
	wire.Struct(new(WordpressCallbackServiceImpl), "*"),
	wire.Bind(new(WordpressCallbackService), new(*WordpressCallbackServiceImpl)),
)

func (s *WordpressCallbackServiceImpl) Import(entityType wordpress.EntityType, id int) error {
	var err error
	switch entityType {
	case wordpress.EntityTypeUser:
		err = s.UserCommandService.ImportFromWordpressByID(id)
	case wordpress.EntityTypePost:
		_, err = s.PostCommandService.ImportFromWordpressByID(id)
	case wordpress.EntityTypeLocation:
		_, err = s.TouristSpotCommandService.ImportFromWordpressByID(id)
	case wordpress.EntityTypeMovie:
		_, err = s.VlogCommandService.ImportFromWordpressByID(id)
	case wordpress.EntityTypeComic:
		_, err = s.ComicCommandService.ImportFromWordpressByID(id)
	case wordpress.EntityTypeFeature:
		_, err = s.FeatureCommandService.ImportFromWordpressByID(id)
	case wordpress.EntityTypeCategory:
		err = s.CategoryCommandService.ImportFromWordpressByID(id)
	case wordpress.EntityTypeLocationCat:
		err = s.SpotCategoryCommandService.ImportFromWordpressByID(id)
	case wordpress.EntityTypeRevision:
		// nop
	default:
		err = serror.New(nil, serror.CodeInvalidParam, "unknown wordpress entity type; %s", entityType)
	}

	if serror.IsErrorCode(err, serror.CodeImportDeleted) {
		return nil
	}

	return errors.Wrap(err, "failed to import wordpress entity")
}
