package service

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity/slack"

	"github.com/google/wire"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"

	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	ReportCommandService interface {
		Report(user *entity.User, cmd *command.Report) error
		MarkAsDone(cmd *command.MarkAsReport) error
	}

	ReportCommandServiceImpl struct {
		repository.ReviewQueryRepository
		repository.ReviewCommandRepository
		repository.SlackRepository
		repository.ReportCommandRepository
		repository.ReportQueryRepository
		TransactionService
	}
)

var ReportCommandServiceSet = wire.NewSet(
	wire.Struct(new(ReportCommandServiceImpl), "*"),
	wire.Bind(new(ReportCommandService), new(*ReportCommandServiceImpl)),
)

func (s *ReportCommandServiceImpl) Report(user *entity.User, cmd *command.Report) error {
	var targetURL string
	var body string
	var reportedUserID int

	// TODO:　ロジックを他に置く、ドメインサービス作った方が良さげ
	// 通報対象の情報を取得
	switch cmd.TargetType {
	case model.ReportTargetTypeREVIEW:
		review, err := s.ReviewQueryRepository.FindByID(cmd.TargetID)
		if err != nil {
			return errors.Wrap(err, "failed to find review")
		}
		targetURL = review.WebURL()
		body = review.Body
		reportedUserID = review.UserID
	case model.ReportTargetTypeCOMMENT:
		comment, err := s.ReviewQueryRepository.FindReviewCommentByID(cmd.TargetID)
		if err != nil {
			return errors.Wrap(err, "failed to find review_comment")
		}
		targetURL = comment.WebURL()
		body = comment.Body
		reportedUserID = comment.UserID
	case model.ReportTargetTypeREPLY:
		reply, err := s.ReviewQueryRepository.FindReviewCommentReplyByID(cmd.TargetID)
		if err != nil {
			return errors.Wrap(err, "failed to find review_comment_reply")
		}
		targetURL = reply.WebURL()
		body = reply.Body
		reportedUserID = reply.UserID
	default:
		return serror.New(nil, serror.CodeInvalidParam, "Invalid target Type")
	}

	// 通報が重複していないか確認
	isExist, err := s.ReportQueryRepository.IsExist(user.ID, cmd.TargetID, cmd.TargetType)
	if err != nil {
		return errors.Wrap(err, "failed to find report is exist")
	}
	if isExist {
		return serror.New(nil, serror.CodeInvalidParam, "Invalid duplicate report")
	}

	report := entity.NewReport(user.ID, cmd.TargetID, cmd.TargetType, cmd.Reason)
	slackReport := slack.NewSlackReport(cmd.TargetType, targetURL, cmd.TargetID, body, cmd.Reason, user.ID, reportedUserID)

	return s.TransactionService.Do(func(c context.Context) error {
		// 通報を永続化
		if err := s.ReportCommandRepository.Store(c, report); err != nil {
			return errors.Wrap(err, "failed to store report")
		}

		// Slackへ通知を送る
		if err := s.SlackRepository.SendReport(slackReport); err != nil {
			return errors.Wrap(err, "failed to send report to slack")
		}
		return nil
	})
}

func (s *ReportCommandServiceImpl) MarkAsDone(cmd *command.MarkAsReport) error {
	isExist, err := s.ReportQueryRepository.IsExist(cmd.UserID, cmd.TargetID, cmd.TargetType)
	if err != nil {
		return errors.Wrap(err, "failed to find report")
	}
	// 存在していない場合
	if !isExist {
		return serror.New(nil, serror.CodeNotFound, "Not Found")
	}

	return s.TransactionService.Do(func(c context.Context) error {
		// 通報を対応済に
		if err := s.ReportCommandRepository.MarkAsDone(c, cmd); err != nil {
			return errors.Wrap(err, "failed to report mark as done")
		}

		// 通報が承認された場合,通報対象を論理削除
		if cmd.IsApproved {
			switch cmd.TargetType {
			case model.ReportTargetTypeREVIEW:
				if err := s.ReviewCommandRepository.DeleteReviewByID(c, cmd.TargetID); err != nil {
					return errors.Wrap(err, "failed to delete review")
				}
			case model.ReportTargetTypeCOMMENT:
				if err := s.ReviewCommandRepository.DeleteReviewCommentByID(c, cmd.TargetID); err != nil {
					return errors.Wrap(err, "failed to delete review_comment")
				}
			case model.ReportTargetTypeREPLY:
				if err := s.ReviewCommandRepository.DeleteReviewCommentReplyByID(c, cmd.TargetID); err != nil {
					return errors.Wrap(err, "failed to delete review_comment_reply")
				}
			}
		}

		return nil
	})
}
