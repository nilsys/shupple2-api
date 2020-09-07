package facade

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	CfProjectFacade interface {
		SendAchievementEmail() error
		SendNewPostEmail() error
		ImportWithGifts(id int) error
	}

	CfProjectFacadeImpl struct {
		service.CfProjectCommandService
		service.CfReturnGiftCommandService
		repository.CfProjectQueryRepository
		repository.CfProjectCommandRepository
		repository.PostQueryRepository
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

func (f *CfProjectFacadeImpl) SendAchievementEmail() error {
	lastID := 0
	for {
		cfProjects, err := f.CfProjectQueryRepository.FindNotSentAchievementNoticeEmailAndAchievedListByLastID(lastID, cfProjectPartLimit)
		if err != nil {
			return errors.Wrap(err, "failed find cf_project")
		}
		if len(cfProjects.List) == 0 {
			break
		}

		for _, cfProject := range cfProjects.List {
			if err := f.CfProjectCommandService.SendAchievementMailToSupporter(cfProject); err != nil {
				return errors.Wrap(err, "failed send achievement email")
			}
		}

		lastID = cfProjects.List[len(cfProjects.List)-1].ID
	}
	return nil
}

func (f *CfProjectFacadeImpl) SendNewPostEmail() error {
	lastID := 0
	for {
		cfProjects, err := f.CfProjectQueryRepository.FindNotSentNewPostNoticeEmailByLastID(lastID, cfProjectPartLimit)
		if err != nil {
			return errors.Wrap(err, "failed find cf_project")
		}
		if len(cfProjects.List) == 0 {
			break
		}

		posts, err := f.PostQueryRepository.FindByIDs(cfProjects.LatestPostIDs())
		if err != nil {
			return errors.Wrap(err, "failed find post")
		}

		cfProjectIDPostDMap := posts.ToCfProjectIDMap()

		for _, cfProject := range cfProjects.List {
			if err := f.CfProjectCommandService.SendNewReportMailToSupporter(cfProject, cfProjectIDPostDMap[cfProject.ID]); err != nil {
				return errors.Wrap(err, "failed send new post email")
			}
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
