package service

import (
	"context"

	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	NoticeDomainService interface {
		ReviewComment(c context.Context, reviewComment *entity.ReviewComment, review *entity.Review) error
	}

	NoticeDomainServiceImpl struct {
		repository.NoticeCommandRepository
		TaggedUserDomainService
	}
)

var NoticeDomainServiceSet = wire.NewSet(
	wire.Struct(new(NoticeDomainServiceImpl), "*"),
	wire.Bind(new(NoticeDomainService), new(NoticeDomainServiceImpl)),
)

// レビューへコメントがあった場合
func (s NoticeDomainServiceImpl) ReviewComment(c context.Context, reviewComment *entity.ReviewComment, review *entity.Review) error {
	notice := entity.NewNotice(
		review.UserID,
		reviewComment.UserID,
		model.NoticeActionTypeCOMMENT,
		model.NoticeActionTargetTypeREVIEW,
		reviewComment.ID,
	)
	// ユーザのタグ付がが含まれていれるユーザを取得
	taggedUsers, err := s.TaggedUserDomainService.FindTaggedUsers(reviewComment.Body)
	if err != nil {
		return err
	}
	for _, taggedUser := range taggedUsers {
		// ユーザのタグ付がが含まれていれるユーザにNoticeを送る
		if err = s.taggedAtComment(c, reviewComment, taggedUser.ID); err != nil {
			return err
		}
	}

	return s.send(c, notice)
}

// ユーザがレビュー内でタグ付されていた場合 (リンター回避のためコメントアウト)
//func (s NoticeDomainServiceImpl) taggedAtReview(c context.Context, review *entity.Review, taggedUserID int) error {
//	notice := entity.NewNotice(
//		taggedUserID,
//		review.UserID,
//		model.NoticeActionTypeTAGGED,
//		model.NoticeActionTargetTypeREVIEW,
//		review.ID,
//	)
//
//	return s.send(c, notice)
//}

// ユーザがコメント内でタグ付されていた場合
func (s NoticeDomainServiceImpl) taggedAtComment(c context.Context, reviewComment *entity.ReviewComment, taggedUserID int) error {
	notice := entity.NewNotice(
		taggedUserID,
		reviewComment.UserID,
		model.NoticeActionTypeTAGGED,
		model.NoticeActionTargetTypeCOMMENT,
		reviewComment.ID,
	)

	return s.send(c, notice)
}

// ユーザがリプライ内でタグ付されていた場合 (リンター回避のためコメントアウト)
//func (s NoticeDomainServiceImpl) taggedAtReply(c context.Context, reviewCommentReply *entity.ReviewCommentReply, taggedUserID int) error {
//	notice := entity.NewNotice(
//		taggedUserID,
//		reviewCommentReply.UserID,
//		model.NoticeActionTypeTAGGED,
//		model.NoticeActionTargetTypeREPLY,
//		reviewCommentReply.ID,
//	)
//
//	return s.send(c, notice)
//}

func (s NoticeDomainServiceImpl) send(c context.Context, notice *entity.Notice) error {
	if notice.IsOwnNotice() {
		return nil
	}

	return s.NoticeCommandRepository.StoreNotice(c, notice)
}
