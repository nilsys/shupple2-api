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
		Review(c context.Context, review *entity.Review) error
		ReviewComment(c context.Context, reviewComment *entity.ReviewComment, review *entity.Review) error
		ReviewCommentReply(c context.Context, reply *entity.ReviewCommentReply, comment *entity.ReviewComment) error
		FavoritePost(c context.Context, favoritePost *entity.UserFavoritePost, post *entity.Post) error
		FavoriteReview(c context.Context, favoriteReview *entity.UserFavoriteReview, review *entity.Review) error
		FavoriteReviewComment(c context.Context, favoriteReviewComment *entity.UserFavoriteReviewComment, reviewComment *entity.ReviewComment) error
		FavoriteReviewCommentReply(c context.Context, favoriteReviewCommentReply *entity.UserFavoriteReviewCommentReply, reviewCommentReply *entity.ReviewCommentReply) error
		FavoriteVlog(c context.Context, favoriteVlog *entity.UserFavoriteVlog, vlog *entity.Vlog) error
		FollowUser(c context.Context, following *entity.UserFollowing) error
		// TODO: enhancement
		//FavoriteVlog()
	}

	NoticeDomainServiceImpl struct {
		repository.NoticeCommandRepository
		TaggedUserDomainService
	}
)

var NoticeDomainServiceSet = wire.NewSet(
	wire.Struct(new(NoticeDomainServiceImpl), "*"),
	wire.Bind(new(NoticeDomainService), new(*NoticeDomainServiceImpl)),
)

func (s *NoticeDomainServiceImpl) Review(c context.Context, review *entity.Review) error {
	taggedUsers, err := s.TaggedUserDomainService.FindTaggedUsers(review.Body)
	if err != nil {
		return err
	}

	for _, taggedUser := range taggedUsers {
		if err = s.taggedAtReview(c, review, taggedUser.ID); err != nil {
			return err
		}
	}

	return nil
}

func (s *NoticeDomainServiceImpl) FavoritePost(c context.Context, favoritePost *entity.UserFavoritePost, post *entity.Post) error {
	notice := entity.NewNotice(
		post.UserID,
		favoritePost.UserID,
		model.NoticeActionTypeFAVORITE,
		model.NoticeActionTargetTypePOST,
		post.ID,
	)

	return s.send(c, notice)
}

func (s *NoticeDomainServiceImpl) FavoriteReview(c context.Context, favoriteReview *entity.UserFavoriteReview, review *entity.Review) error {
	notice := entity.NewNotice(
		review.UserID,
		favoriteReview.UserID,
		model.NoticeActionTypeFAVORITE,
		model.NoticeActionTargetTypeREVIEW,
		review.ID,
	)

	return s.send(c, notice)
}

func (s *NoticeDomainServiceImpl) FavoriteReviewComment(c context.Context, favoriteReviewComment *entity.UserFavoriteReviewComment, reviewComment *entity.ReviewComment) error {
	notice := entity.NewNotice(
		reviewComment.UserID,
		favoriteReviewComment.UserID,
		model.NoticeActionTypeFAVORITE,
		model.NoticeActionTargetTypeCOMMENT,
		reviewComment.ID,
	)

	return s.send(c, notice)
}
func (s *NoticeDomainServiceImpl) FavoriteReviewCommentReply(c context.Context, favoriteReviewCommentReply *entity.UserFavoriteReviewCommentReply, reviewCommentReply *entity.ReviewCommentReply) error {
	notice := entity.NewNotice(
		reviewCommentReply.UserID,
		favoriteReviewCommentReply.UserID,
		model.NoticeActionTypeFAVORITE,
		model.NoticeActionTargetTypeREPLY,
		reviewCommentReply.ID,
	)

	return s.send(c, notice)
}

func (s *NoticeDomainServiceImpl) FavoriteVlog(c context.Context, favoriteVlog *entity.UserFavoriteVlog, vlog *entity.Vlog) error {
	notice := entity.NewNotice(
		vlog.UserID,
		favoriteVlog.UserID,
		model.NoticeActionTypeFAVORITE,
		model.NoticeActionTargetTypeVLOG,
		vlog.ID,
	)

	return s.send(c, notice)
}

// レビューへコメントがあった場合
func (s *NoticeDomainServiceImpl) ReviewComment(c context.Context, reviewComment *entity.ReviewComment, review *entity.Review) error {
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

// レビューのコメントへリプライがあった場合
func (s *NoticeDomainServiceImpl) ReviewCommentReply(c context.Context, reply *entity.ReviewCommentReply, comment *entity.ReviewComment) error {
	notice := entity.NewNotice(
		comment.UserID,
		reply.UserID,
		model.NoticeActionTypeREPLY,
		model.NoticeActionTargetTypeCOMMENT,
		reply.ID,
	)
	// ユーザのタグ付がが含まれていれるユーザを取得
	taggedUsers, err := s.TaggedUserDomainService.FindTaggedUsers(reply.Body)
	if err != nil {
		return err
	}
	for _, taggedUser := range taggedUsers {
		// ユーザのタグ付がが含まれていれるユーザにNoticeを送る
		if err = s.taggedAtReply(c, reply, taggedUser.ID); err != nil {
			return err
		}
	}

	return s.send(c, notice)
}

func (s *NoticeDomainServiceImpl) FollowUser(c context.Context, following *entity.UserFollowing) error {
	notice := entity.NewNotice(
		following.TargetID,
		following.UserID,
		model.NoticeActionTypeFOLLOW,
		model.NoticeActionTargetTypeUSER,
		following.UserID,
	)

	return s.send(c, notice)
}

// ユーザがレビュー内でタグ付されていた場合
func (s *NoticeDomainServiceImpl) taggedAtReview(c context.Context, review *entity.Review, taggedUserID int) error {
	notice := entity.NewNotice(
		taggedUserID,
		review.UserID,
		model.NoticeActionTypeTAGGED,
		model.NoticeActionTargetTypeREVIEW,
		review.ID,
	)

	return s.send(c, notice)
}

// ユーザがコメント内でタグ付されていた場合
func (s *NoticeDomainServiceImpl) taggedAtComment(c context.Context, reviewComment *entity.ReviewComment, taggedUserID int) error {
	notice := entity.NewNotice(
		taggedUserID,
		reviewComment.UserID,
		model.NoticeActionTypeTAGGED,
		model.NoticeActionTargetTypeCOMMENT,
		reviewComment.ID,
	)

	return s.send(c, notice)
}

// ユーザがリプライ内でタグ付されていた場合
func (s *NoticeDomainServiceImpl) taggedAtReply(c context.Context, reviewCommentReply *entity.ReviewCommentReply, taggedUserID int) error {
	notice := entity.NewNotice(
		taggedUserID,
		reviewCommentReply.UserID,
		model.NoticeActionTypeTAGGED,
		model.NoticeActionTargetTypeREPLY,
		reviewCommentReply.ID,
	)

	return s.send(c, notice)
}

func (s *NoticeDomainServiceImpl) send(c context.Context, notice *entity.Notice) error {
	if notice.IsOwnNotice() {
		return nil
	}

	return s.NoticeCommandRepository.StoreNotice(c, notice)
}
