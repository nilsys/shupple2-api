package service

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/config"

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
		repository.UserQueryRepository
		config.StaywayMedia
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
	case model.ReportTargetTypeReview:
		review, err := s.ReviewQueryRepository.FindByID(cmd.TargetID)
		if err != nil {
			return errors.Wrap(err, "failed to find review")
		}
		targetURL = review.MediaWebURL(s.StaywayMedia.BaseURL).String()
		body = review.Body
		reportedUserID = review.UserID
	case model.ReportTargetTypeComment:
		comment, err := s.ReviewQueryRepository.FindReviewCommentDetailByID(cmd.TargetID)
		if err != nil {
			return errors.Wrap(err, "failed to find review_comment")
		}
		targetURL = comment.Review.MediaWebURL(s.StaywayMedia.BaseURL).String()
		body = comment.Body
		reportedUserID = comment.UserID
	case model.ReportTargetTypeReply:
		reply, err := s.ReviewQueryRepository.FindReviewCommentReplyDetailByID(cmd.TargetID)
		if err != nil {
			return errors.Wrap(err, "failed to find review_comment_reply")
		}
		targetURL = reply.ReviewCommentDetail.Review.MediaWebURL(s.StaywayMedia.BaseURL).String()
		body = reply.Body
		reportedUserID = reply.UserID
	case model.ReportTargetTypeUser:
		targetUser, err := s.UserQueryRepository.FindByID(cmd.TargetID)
		if err != nil {
			return errors.Wrap(err, "failed find user")
		}
		targetURL = targetUser.MediaWebURL(s.StaywayMedia.BaseURL).String()
		body = targetUser.Profile
		reportedUserID = targetUser.ID
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

	report := entity.NewReport(user.ID, cmd.TargetID, cmd.TargetType, cmd.Reason, cmd.Body)
	slackReport := slack.NewSlackReport(cmd.TargetType, targetURL, cmd.TargetID, body, cmd.Body, cmd.Reason, user.ID, reportedUserID)

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
			case model.ReportTargetTypeReview:
				if err := s.ReviewCommandRepository.DeleteReviewByID(c, cmd.TargetID); err != nil {
					return errors.Wrap(err, "failed to delete review")
				}
			case model.ReportTargetTypeComment:
				if err := s.ReviewCommandRepository.DeleteReviewCommentByID(c, cmd.TargetID); err != nil {
					return errors.Wrap(err, "failed to delete review_comment")
				}
			case model.ReportTargetTypeReply:
				if err := s.ReviewCommandRepository.DeleteReviewCommentReplyByID(c, cmd.TargetID); err != nil {
					return errors.Wrap(err, "failed to delete review_comment_reply")
				}
			}
		}

		return nil
	})
}
