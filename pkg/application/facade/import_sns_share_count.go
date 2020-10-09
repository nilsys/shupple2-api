package facade

import (
	"fmt"
	"strconv"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
	facebookRepo "github.com/stayway-corp/stayway-media-api/pkg/domain/repository/facebook"
	widgetoonJsoonRepo "github.com/stayway-corp/stayway-media-api/pkg/domain/repository/widgetoonjsoon"
	"golang.org/x/sync/errgroup"
)

type (
	ImportSnsShareCountFacade interface {
		ImportPostFacebookShareCount() error
		ImportVlogFacebookShareCount() error
		ImportCfProjectFacebookShareCount() error
		ImportTwitterShareCount() error
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
		repository.BatchOptionCommandRepository
		repository.BatchOptionQueryRepository
		*config.Config
	}
)

var ImportSnsShareCountFacadeSet = wire.NewSet(
	wire.Struct(new(ImportSnsShareCountFacadeImpl), "*"),
	wire.Bind(new(ImportSnsShareCountFacade), new(*ImportSnsShareCountFacadeImpl)),
)

const (
	/*
		1h -> 200 graph api request limit

		facebook graph apiを用いてengagement(シェア数)を取得する際に
		https://graph.facebook.com/?id=https://stayway.jp/tourism/asia-heritage13/
		と
		https://graph.facebook.com/?id=https://stayway.jp/tourism/asia-heritage13
		をトレイリングスラッシュで区別する
		つまり、バッチリクエスト(50リクエストを1リクエストへ)を用いて1リクエストで25件集計する事が出来る
		よって25件ずつ処理する

		参考：
		バッチリクエスト：https://developers.facebook.com/docs/graph-api/making-multiple-requests
		リクエスト制限：https://developers.facebook.com/docs/graph-api/advanced/rate-limited
	*/
	queryLimitForFacebook = 25
	queryLimitForTwitter  = 100
	/*
		バッチリクエストでは最大50リクエストを1リクエストにまとめる事が出来るだけで、実際の挙動を見ると50リクエストとして換算されているので200/50=4
		実際の挙動では以下を確認
		(4リクエスト送って100%が98%へ)
		(7リクエスト送って98%が93%へ)
		これに初回のトークン取得で1回リクエストを消費するので、4-1=3とする
		上限に達する事が無い様に抑えめの値にしている
	*/
	facebookGraphAPIReqLimit = 3
)

func (s *ImportSnsShareCountFacadeImpl) ImportPostFacebookShareCount() error {
	lastIDStr, err := s.BatchOptionQueryRepository.FirstOrCreateByName(model.BatchOptionNameImportFacebookShareCountLastPostID)
	if err != nil {
		return errors.Wrap(err, "failed ref batch_option")
	}

	var lastID int

	// 初回は空文字が返ってくるのでチェック
	if lastIDStr == "" {
		lastID = 0
	} else {
		lastID, err = strconv.Atoi(lastIDStr)
		if err != nil {
			return errors.Wrap(err, "can't assert batch_option value")
		}
	}

	for graphAPIReqCount := 0; graphAPIReqCount < facebookGraphAPIReqLimit; graphAPIReqCount++ {
		posts, err := s.PostQueryRepository.FindByLastID(lastID, queryLimitForFacebook)
		if err != nil {
			return errors.Wrap(err, "failed to find by lastID")
		}
		// 最新のPostまで集計が完了した場合は次のバッチでid = 1から処理する為0を入れる
		if len(posts) == 0 {
			if err := s.BatchOptionCommandRepository.UpdateByName(model.BatchOptionNameImportFacebookShareCountLastPostID, strconv.Itoa(0)); err != nil {
				return errors.Wrap(err, "failed update batch_option")
			}
			break
		}

		result, err := s.FacebookQueryRepository.GetShareCountByURLBatchRequest(posts.ToGraphAPIBatchRequestQueryStr(s.Config.Stayway.Media.BaseURL))
		if err != nil {
			return errors.Wrap(err, "failed facebook error")
		}
		for _, post := range posts {
			fmt.Printf("--------------- slug: %s, cnt: %d ------------", post.Slug, result.GetShareCountBySuffixKey(post.Slug))
			if err := s.PostCommandRepository.UpdateFacebookCountByID(post.ID, result.GetShareCountBySuffixKey(post.Slug)); err != nil {
				return errors.Wrap(err, "failed update facebook facebook_count")
			}
		}

		lastID = posts[len(posts)-1].ID
		// 次のバッチで続きから処理する為、集計した最後のidを入れる
		if err := s.BatchOptionCommandRepository.UpdateByName(model.BatchOptionNameImportFacebookShareCountLastPostID, strconv.Itoa(lastID)); err != nil {
			return errors.Wrap(err, "failed update batch_option")
		}
	}

	return nil
}

func (s *ImportSnsShareCountFacadeImpl) ImportVlogFacebookShareCount() error {
	lastIDStr, err := s.BatchOptionQueryRepository.FirstOrCreateByName(model.BatchOptionNameImportFacebookShareCountLastVlogID)
	if err != nil {
		return errors.Wrap(err, "failed ref batch_option")
	}

	var lastID int

	// 初回は空文字が返ってくるのでチェック
	if lastIDStr == "" {
		lastID = 0
	} else {
		lastID, err = strconv.Atoi(lastIDStr)
		if err != nil {
			return errors.Wrap(err, "can't assert batch_option value")
		}
	}

	for graphAPIReqCount := 0; graphAPIReqCount < facebookGraphAPIReqLimit; graphAPIReqCount++ {
		vlogs, err := s.VlogQueryRepository.FindByLastID(lastID, queryLimitForFacebook)
		if err != nil {
			return errors.Wrap(err, "failed to find by lastID")
		}
		// 最新のVlogまで集計が完了した場合は次のバッチでid = 1から処理する為0を入れる
		if len(vlogs) == 0 {
			if err := s.BatchOptionCommandRepository.UpdateByName(model.BatchOptionNameImportFacebookShareCountLastVlogID, strconv.Itoa(0)); err != nil {
				return errors.Wrap(err, "failed update batch_option")
			}
			break
		}

		result, err := s.FacebookQueryRepository.GetShareCountByURLBatchRequest(vlogs.ToGraphAPIBatchRequestQueryStr(s.Config.Stayway.Media.BaseURL))
		if err != nil {
			return errors.Wrap(err, "failed facebook error")
		}
		for _, vlog := range vlogs {
			if err := s.VlogCommandRepository.UpdateFacebookCountByID(vlog.ID, result.GetShareCountBySuffixKey(strconv.Itoa(vlog.ID))); err != nil {
				return errors.Wrap(err, "failed update facebook count")
			}
		}

		lastID = vlogs[len(vlogs)-1].ID
		// 次のバッチで続きから処理する為、集計出来た最後のidを入れる
		if err := s.BatchOptionCommandRepository.UpdateByName(model.BatchOptionNameImportFacebookShareCountLastVlogID, strconv.Itoa(lastID)); err != nil {
			return errors.Wrap(err, "failed update batch_option")
		}
	}

	return nil
}

func (s *ImportSnsShareCountFacadeImpl) ImportCfProjectFacebookShareCount() error {
	lastIDStr, err := s.BatchOptionQueryRepository.FirstOrCreateByName(model.BatchOptionNameImportFacebookShareCountLastCfProjectID)
	if err != nil {
		return errors.Wrap(err, "failed ref batch_option")
	}

	var lastID int

	// 初回は空文字が返ってくるのでチェック
	if lastIDStr == "" {
		lastID = 0
	} else {
		lastID, err = strconv.Atoi(lastIDStr)
		if err != nil {
			return errors.Wrap(err, "can't assert batch_option value")
		}
	}

	for graphAPIReqCount := 0; graphAPIReqCount < facebookGraphAPIReqLimit; graphAPIReqCount++ {
		projects, err := s.CfProjectQueryRepository.FindByLastID(lastID, queryLimitForFacebook)
		if err != nil {
			return errors.Wrap(err, "failed to find by lastID")
		}
		// 最新のCfProjectまで集計が完了した場合は次のバッチでid = 1から処理する為0を入れる
		if len(projects) == 0 {
			if err := s.BatchOptionCommandRepository.UpdateByName(model.BatchOptionNameImportFacebookShareCountLastCfProjectID, strconv.Itoa(0)); err != nil {
				return errors.Wrap(err, "failed update batch_option")
			}
			break
		}

		result, err := s.FacebookQueryRepository.GetShareCountByURLBatchRequest(projects.ToGraphAPIBatchRequestQueryStr(s.Config.Stayway.Media.BaseURL))
		if err != nil {
			return errors.Wrap(err, "failed facebook error")
		}
		for _, project := range projects {
			if err := s.CfProjectCommandRepository.UpdateFacebookCountByID(project.ID, result.GetShareCountBySuffixKey(strconv.Itoa(project.ID))); err != nil {
				return errors.Wrap(err, "failed update facebook count")
			}
		}

		lastID = projects[len(projects)-1].ID
		// 次のバッチで続きから処理する為、集計出来た最後のidを入れる
		if err := s.BatchOptionCommandRepository.UpdateByName(model.BatchOptionNameImportFacebookShareCountLastCfProjectID, strconv.Itoa(lastID)); err != nil {
			return errors.Wrap(err, "failed update batch_option")
		}
	}

	return nil
}

func (s *ImportSnsShareCountFacadeImpl) ImportTwitterShareCount() error {
	eg := errgroup.Group{}
	eg.Go(s.importPostTwitterShareCount)
	eg.Go(s.importVlogTwitterShareCount)
	eg.Go(s.importCfProjectTwitterShareCount)
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

func (s *ImportSnsShareCountFacadeImpl) importPostTwitterShareCount() error {
	lastID := 0
	for {
		posts, err := s.PostQueryRepository.FindByLastID(lastID, queryLimitForTwitter)
		if err != nil {
			return errors.Wrap(err, "failed to find by lastID")
		}
		if len(posts) == 0 {
			break
		}

		for _, post := range posts {
			mediaWebURLStr := post.MediaWebURL(s.Config.Stayway.Media.BaseURL).String()

			twitterShareCnt, err := s.WidgetoonJsoonQueryRepository.GetTwitterCountByURL(mediaWebURLStr)
			if err != nil {
				return errors.Wrap(err, "failed get twitter count")
			}

			if err := s.PostCommandRepository.UpdateTwitterCountByID(post.ID, twitterShareCnt.Count); err != nil {
				return errors.Wrap(err, "failed update twitter count")
			}
		}

		lastID = posts[len(posts)-1].ID
	}

	return nil
}

func (s *ImportSnsShareCountFacadeImpl) importVlogTwitterShareCount() error {
	lastID := 0
	for {
		vlogs, err := s.VlogQueryRepository.FindByLastID(lastID, queryLimitForTwitter)
		if err != nil {
			return errors.Wrap(err, "failed to find by lastID")
		}
		if len(vlogs) == 0 {
			break
		}

		for _, vlog := range vlogs {
			mediaWebURLStr := vlog.MediaWebURL(s.Config.Stayway.Media.BaseURL).String()

			twitterShareCnt, err := s.WidgetoonJsoonQueryRepository.GetTwitterCountByURL(mediaWebURLStr)
			if err != nil {
				return errors.Wrap(err, "failed get twitter count")
			}

			if err := s.VlogCommandRepository.UpdateTwitterCountByID(vlog.ID, twitterShareCnt.Count); err != nil {
				return errors.Wrap(err, "failed update twitter count")
			}
		}

		lastID = vlogs[len(vlogs)-1].ID
	}

	return nil
}

func (s *ImportSnsShareCountFacadeImpl) importCfProjectTwitterShareCount() error {
	lastID := 0
	for {
		projects, err := s.CfProjectQueryRepository.FindByLastID(lastID, queryLimitForTwitter)
		if err != nil {
			return errors.Wrap(err, "failed to find by lastID")
		}
		if len(projects) == 0 {
			break
		}

		for _, project := range projects {
			mediaWebURLStr := project.MediaWebURL(s.Config.Stayway.Media.BaseURL).String()

			twitterShareCnt, err := s.WidgetoonJsoonQueryRepository.GetTwitterCountByURL(mediaWebURLStr)
			if err != nil {
				return errors.Wrap(err, "failed get twitter count")
			}

			if err := s.CfProjectCommandRepository.UpdateTwitterCountByID(project.ID, twitterShareCnt.Count); err != nil {
				return errors.Wrap(err, "failed update twitter count")
			}
		}

		lastID = projects[len(projects)-1].ID
	}

	return nil
}
