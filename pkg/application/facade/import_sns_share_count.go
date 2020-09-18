package facade

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
	facebookRepo "github.com/stayway-corp/stayway-media-api/pkg/domain/repository/facebook"
	widgetoonJsoonRepo "github.com/stayway-corp/stayway-media-api/pkg/domain/repository/widgetoonjsoon"
)

type (
	ImportSnsShareCountFacade interface {
		ImportPostSnsShareCount() error
		ImportVlogSnsShareCount() error
		ImportCfProjectSnsShareCount() error
	}

	ImportSnsShareCountFacadeImpl struct {
		FacebookQueryRepository       facebookRepo.QueryRepository
		WidgetoonJsoonQueryRepository widgetoonJsoonRepo.QueryRepository
		repository.PostQueryRepository
		repository.PostCommandRepository
		repository.VlogQueryRepository
		repository.VlogCommandRepository
		repository.CfProjectQueryRepository
		repository.CfProjectCommandRepository
		*config.Config
	}
)

var ImportSnsShareCountFacadeSet = wire.NewSet(
	wire.Struct(new(ImportSnsShareCountFacadeImpl), "*"),
	wire.Bind(new(ImportSnsShareCountFacade), new(*ImportSnsShareCountFacadeImpl)),
)

const (
	queryLimit = 100
)

func (s *ImportSnsShareCountFacadeImpl) ImportPostSnsShareCount() error {
	lastID := 0
	for {
		posts, err := s.PostQueryRepository.FindByLastID(lastID, queryLimit)
		if err != nil {
			return errors.Wrap(err, "failed to find by lastID")
		}
		if len(posts) == 0 {
			break
		}

		for _, post := range posts {
			mediaWebURLStr := post.MediaWebURL(s.Config.Stayway.Media.BaseURL).String()

			twitterShareCnt, err := s.WidgetoonJsoonQueryRepository.TwitterCountByURL(mediaWebURLStr)
			if err != nil {
				return errors.Wrap(err, "failed get twitter count")
			}

			facebookShareCnt, err := s.FacebookQueryRepository.GetShareCountByURL(mediaWebURLStr)
			if err != nil {
				return errors.Wrap(err, "failed get facebook share count")
			}

			if err := s.PostCommandRepository.UpdateTwitterCountByID(post.ID, twitterShareCnt.Count); err != nil {
				return errors.Wrap(err, "failed update twitter count")
			}
			if err := s.PostCommandRepository.UpdateFacebookCountByID(post.ID, facebookShareCnt); err != nil {
				return errors.Wrap(err, "failed update facebook count")
			}
		}

		lastID = posts[len(posts)-1].ID
	}

	return nil
}

func (s *ImportSnsShareCountFacadeImpl) ImportVlogSnsShareCount() error {
	lastID := 0
	for {
		vlogs, err := s.VlogQueryRepository.FindByLastID(lastID, queryLimit)
		if err != nil {
			return errors.Wrap(err, "failed to find by lastID")
		}
		if len(vlogs) == 0 {
			break
		}

		for _, vlog := range vlogs {
			mediaWebURLStr := vlog.MediaWebURL(s.Config.Stayway.Media.BaseURL).String()

			twitterShareCnt, err := s.WidgetoonJsoonQueryRepository.TwitterCountByURL(mediaWebURLStr)
			if err != nil {
				return errors.Wrap(err, "failed get twitter count")
			}

			facebookShareCnt, err := s.FacebookQueryRepository.GetShareCountByURL(mediaWebURLStr)
			if err != nil {
				return errors.Wrap(err, "failed get facebook share count")
			}

			if err := s.VlogCommandRepository.UpdateTwitterCountByID(vlog.ID, twitterShareCnt.Count); err != nil {
				return errors.Wrap(err, "failed update twitter count")
			}
			if err := s.VlogCommandRepository.UpdateFacebookCountByID(vlog.ID, facebookShareCnt); err != nil {
				return errors.Wrap(err, "failed update facebook count")
			}
		}

		lastID = vlogs[len(vlogs)-1].ID
	}

	return nil
}

func (s *ImportSnsShareCountFacadeImpl) ImportCfProjectSnsShareCount() error {
	lastID := 0
	for {
		projects, err := s.CfProjectQueryRepository.FindByLastID(lastID, 100)
		if err != nil {
			return errors.Wrap(err, "failed to find by lastID")
		}
		if len(projects) == 0 {
			break
		}

		for _, project := range projects {
			mediaWebURLStr := project.MediaWebURL(s.Config.Stayway.Media.BaseURL).String()

			twitterShareCnt, err := s.WidgetoonJsoonQueryRepository.TwitterCountByURL(mediaWebURLStr)
			if err != nil {
				return errors.Wrap(err, "failed get twitter count")
			}

			facebookShareCnt, err := s.FacebookQueryRepository.GetShareCountByURL(mediaWebURLStr)
			if err != nil {
				return errors.Wrap(err, "failed get facebook share count")
			}

			if err := s.CfProjectCommandRepository.UpdateTwitterCountByID(project.ID, twitterShareCnt.Count); err != nil {
				return errors.Wrap(err, "failed update twitter count")
			}
			if err := s.CfProjectCommandRepository.UpdateFacebookCountByID(project.ID, facebookShareCnt); err != nil {
				return errors.Wrap(err, "failed update facebook count")
			}
		}

		lastID = projects[len(projects)-1].ID
	}

	return nil
}
