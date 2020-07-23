package facade

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	CfProjectFacade interface {
		SendAchievementMail() error
		ImportWithGifts(id int) error
	}

	CfProjectFacadeImpl struct {
		service.CfProjectCommandService
		service.CfReturnGiftCommandService
		repository.CfProjectQueryRepository
		repository.WordpressQueryRepository
	}
)

var CfProjectFacadeSet = wire.NewSet(
	wire.Struct(new(CfProjectFacadeImpl), "*"),
	wire.Bind(new(CfProjectFacade), new(*CfProjectFacadeImpl)),
)

const (
	cfProjectPartLimit = 100
)

func (f *CfProjectFacadeImpl) SendAchievementMail() error {
	lastID := 0
	for {
		cfProjects, err := f.CfProjectQueryRepository.FindNotSentAchievementNoticeMailAndAchievedListByLastID(lastID, cfProjectPartLimit)
		if err != nil {
			return errors.Wrap(err, "failed find cf_project")
		}
		if len(cfProjects.List) == 0 {
			break
		}
		for _, cfProject := range cfProjects.List {
			return f.CfProjectCommandService.SendAchievementMailToSupporter(cfProject)
		}

		lastID = cfProjects.List[len(cfProjects.List)-1].ID
	}
	return nil
}

func (f *CfProjectFacadeImpl) ImportWithGifts(id int) error {
	if err := f.CfProjectCommandService.ImportFromWordpressByID(id); err != nil {
		return errors.Wrap(err, "failed to import cf_project")
	}

	gifts, err := f.WordpressQueryRepository.FindCfReturnGiftsByCfProjectID(id)
	if err != nil {
		return errors.Wrap(err, "failed to list target cf_return_gifts")
	}

	for _, gift := range gifts {
		if err := f.CfReturnGiftCommandService.ImportFromWordpressByID(gift.ID); err != nil {
			return errors.Wrap(err, "failed to import cf_return_gift")
		}
	}

	return nil
}
