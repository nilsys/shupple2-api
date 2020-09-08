package service

import (
	"context"
	"fmt"
	"strconv"

	firebaseEntity "github.com/stayway-corp/stayway-media-api/pkg/domain/entity/firebase"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"

	"github.com/pkg/errors"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository/firebase"

	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	NoticeDomainService interface {
		Review(c context.Context, review *entity.Review, triggeredUser *entity.User) error
		ReviewComment(c context.Context, reviewComment *entity.ReviewComment, review *entity.Review, triggeredUser *entity.User) error
		ReviewCommentReply(c context.Context, reply *entity.ReviewCommentReply, comment *entity.ReviewComment, review *entity.Review, triggeredUser *entity.User) error
		FavoritePost(c context.Context, favoritePost *entity.UserFavoritePost, post *entity.Post, triggeredUser *entity.User) error
		FavoriteReview(c context.Context, favoriteReview *entity.UserFavoriteReview, review *entity.Review, triggeredUser *entity.User) error
		FavoriteReviewComment(c context.Context, favoriteReviewComment *entity.UserFavoriteReviewComment, reviewComment *entity.ReviewComment, review *entity.Review, triggeredUser *entity.User) error
		FavoriteComic(c context.Context, favoriteComic *entity.UserFavoriteComic, comic *entity.Comic, triggeredUser *entity.User) error
		FavoriteReviewCommentReply(c context.Context, favoriteReviewCommentReply *entity.UserFavoriteReviewCommentReply, reviewCommentReply *entity.ReviewCommentReply, review *entity.Review, triggeredUser *entity.User) error
		FavoriteVlog(c context.Context, favoriteVlog *entity.UserFavoriteVlog, vlog *entity.Vlog, triggeredUser *entity.User) error
		FollowUser(c context.Context, following *entity.UserFollowing, triggeredUser *entity.User) error
		// TODO: enhancement
		//FavoriteVlog()
	}

	NoticeDomainServiceImpl struct {
		repository.NoticeCommandRepository
		repository.UserQueryRepository
		repository.NoticeQueryRepository
		firebase.CloudMessageCommandRepository
		TaggedUserDomainService
	}
)

var NoticeDomainServiceSet = wire.NewSet(
	wire.Struct(new(NoticeDomainServiceImpl), "*"),
	wire.Bind(new(NoticeDomainService), new(*NoticeDomainServiceImpl)),
)

func (s *NoticeDomainServiceImpl) Review(c context.Context, review *entity.Review, triggeredUser *entity.User) error {
	taggedUsers, err := s.TaggedUserDomainService.FindTaggedUsers(review.Body)
	if err != nil {
		return err
	}

	for _, taggedUser := range taggedUsers {
		if err = s.taggedAtReview(c, review, taggedUser.ID, triggeredUser); err != nil {
			return err
		}
	}

	return nil
}

func (s *NoticeDomainServiceImpl) FavoritePost(c context.Context, favoritePost *entity.UserFavoritePost, post *entity.Post, triggeredUser *entity.User) error {
	notice := entity.NewNotice(
		post.UserID,
		favoritePost.UserID,
		model.NoticeActionTypeFAVORITE,
		model.NoticeActionTargetTypePOST,
		post.ID,
		fmt.Sprintf("/%s", post.Slug),
	)

	return s.send(c, notice, post.UserID, triggeredUser)
}

func (s *NoticeDomainServiceImpl) FavoriteReview(c context.Context, favoriteReview *entity.UserFavoriteReview, review *entity.Review, triggeredUser *entity.User) error {
	notice := entity.NewNotice(
		review.UserID,
		favoriteReview.UserID,
		model.NoticeActionTypeFAVORITE,
		model.NoticeActionTargetTypeREVIEW,
		review.ID,
		s.reviewEndpoint(review),
	)

	return s.send(c, notice, review.UserID, triggeredUser)
}

func (s *NoticeDomainServiceImpl) FavoriteReviewComment(c context.Context, favoriteReviewComment *entity.UserFavoriteReviewComment, reviewComment *entity.ReviewComment, review *entity.Review, triggeredUser *entity.User) error {
	notice := entity.NewNotice(
		reviewComment.UserID,
		favoriteReviewComment.UserID,
		model.NoticeActionTypeFAVORITE,
		model.NoticeActionTargetTypeCOMMENT,
		reviewComment.ID,
		s.reviewCommentEndpoint(review, reviewComment),
	)

	return s.send(c, notice, reviewComment.UserID, triggeredUser)
}
func (s *NoticeDomainServiceImpl) FavoriteReviewCommentReply(c context.Context, favoriteReviewCommentReply *entity.UserFavoriteReviewCommentReply, reviewCommentReply *entity.ReviewCommentReply, review *entity.Review, triggeredUser *entity.User) error {
	notice := entity.NewNotice(
		reviewCommentReply.UserID,
		favoriteReviewCommentReply.UserID,
		model.NoticeActionTypeFAVORITE,
		model.NoticeActionTargetTypeREPLY,
		reviewCommentReply.ID,
		s.reviewCommentReplyEndpoint(review, reviewCommentReply),
	)

	return s.send(c, notice, reviewCommentReply.UserID, triggeredUser)
}

func (s *NoticeDomainServiceImpl) FavoriteComic(c context.Context, favoriteComic *entity.UserFavoriteComic, comic *entity.Comic, triggeredUser *entity.User) error {
	notice := entity.NewNotice(
		comic.UserID,
		favoriteComic.UserID,
		model.NoticeActionTypeFAVORITE,
		model.NoticeActionTargetTypeCOMIC,
		comic.ID,
		fmt.Sprintf("/tourism/comic/%d", comic.ID),
	)

	return s.send(c, notice, comic.UserID, triggeredUser)
}

func (s *NoticeDomainServiceImpl) FavoriteVlog(c context.Context, favoriteVlog *entity.UserFavoriteVlog, vlog *entity.Vlog, triggeredUser *entity.User) error {
	notice := entity.NewNotice(
		vlog.UserID,
		favoriteVlog.UserID,
		model.NoticeActionTypeFAVORITE,
		model.NoticeActionTargetTypeVLOG,
		vlog.ID,
		fmt.Sprintf("/tourism/movie/%d", vlog.ID),
	)

	return s.send(c, notice, vlog.UserID, triggeredUser)
}

// レビューへコメントがあった場合
func (s *NoticeDomainServiceImpl) ReviewComment(c context.Context, reviewComment *entity.ReviewComment, review *entity.Review, triggeredUser *entity.User) error {
	notice := entity.NewNotice(
		review.UserID,
		reviewComment.UserID,
		model.NoticeActionTypeCOMMENT,
		model.NoticeActionTargetTypeREVIEW,
		reviewComment.ID,
		s.reviewEndpoint(review),
	)
	// ユーザのタグ付がが含まれていれるユーザを取得
	taggedUsers, err := s.TaggedUserDomainService.FindTaggedUsers(reviewComment.Body)
	if err != nil {
		return err
	}
	for _, taggedUser := range taggedUsers {
		// ユーザのタグ付がが含まれていれるユーザにNoticeを送る
		if err = s.taggedAtComment(c, reviewComment, taggedUser.ID, review, triggeredUser); err != nil {
			return err
		}
	}

	return s.send(c, notice, review.UserID, triggeredUser)
}

// レビューのコメントへリプライがあった場合
func (s *NoticeDomainServiceImpl) ReviewCommentReply(c context.Context, reply *entity.ReviewCommentReply, comment *entity.ReviewComment, review *entity.Review, triggeredUser *entity.User) error {
	notice := entity.NewNotice(
		comment.UserID,
		reply.UserID,
		model.NoticeActionTypeREPLY,
		model.NoticeActionTargetTypeCOMMENT,
		reply.ID,
		s.reviewEndpoint(review),
	)
	// ユーザのタグ付がが含まれていれるユーザを取得
	taggedUsers, err := s.TaggedUserDomainService.FindTaggedUsers(reply.Body)
	if err != nil {
		return err
	}
	for _, taggedUser := range taggedUsers {
		// ユーザのタグ付がが含まれていれるユーザにNoticeを送る
		if err = s.taggedAtReply(c, reply, taggedUser.ID, review, triggeredUser); err != nil {
			return err
		}
	}

	return s.send(c, notice, comment.UserID, triggeredUser)
}

func (s *NoticeDomainServiceImpl) FollowUser(c context.Context, following *entity.UserFollowing, triggeredUser *entity.User) error {
	notice := entity.NewNotice(
		following.TargetID,
		following.UserID,
		model.NoticeActionTypeFOLLOW,
		model.NoticeActionTargetTypeUSER,
		following.UserID,
		fmt.Sprintf("/users/%s", triggeredUser.UID),
	)

	return s.send(c, notice, following.TargetID, triggeredUser)
}

// ユーザがレビュー内でタグ付されていた場合
func (s *NoticeDomainServiceImpl) taggedAtReview(c context.Context, review *entity.Review, taggedUserID int, triggeredUser *entity.User) error {
	notice := entity.NewNotice(
		taggedUserID,
		review.UserID,
		model.NoticeActionTypeTAGGED,
		model.NoticeActionTargetTypeREVIEW,
		review.ID,
		s.reviewEndpoint(review),
	)

	return s.send(c, notice, taggedUserID, triggeredUser)
}

// ユーザがコメント内でタグ付されていた場合
func (s *NoticeDomainServiceImpl) taggedAtComment(c context.Context, reviewComment *entity.ReviewComment, taggedUserID int, review *entity.Review, triggeredUser *entity.User) error {
	notice := entity.NewNotice(
		taggedUserID,
		reviewComment.UserID,
		model.NoticeActionTypeTAGGED,
		model.NoticeActionTargetTypeCOMMENT,
		reviewComment.ID,
		s.reviewCommentEndpoint(review, reviewComment),
	)

	return s.send(c, notice, taggedUserID, triggeredUser)
}

// ユーザがリプライ内でタグ付されていた場合
func (s *NoticeDomainServiceImpl) taggedAtReply(c context.Context, reviewCommentReply *entity.ReviewCommentReply, taggedUserID int, review *entity.Review, triggeredUser *entity.User) error {
	notice := entity.NewNotice(
		taggedUserID,
		reviewCommentReply.UserID,
		model.NoticeActionTypeTAGGED,
		model.NoticeActionTargetTypeREPLY,
		reviewCommentReply.ID,
		s.reviewCommentReplyEndpoint(review, reviewCommentReply),
	)

	return s.send(c, notice, taggedUserID, triggeredUser)
}

func (s *NoticeDomainServiceImpl) send(c context.Context, notice *entity.Notice, sendTargetUserID int, triggeredUser *entity.User) error {
	if notice.IsOwnNotice() {
		return nil
	}

	user, err := s.UserQueryRepository.FindByID(sendTargetUserID)
	if err != nil {
		return errors.Wrap(err, "failed find user")
	}

	// アプリユーザーでない場合は通知処理しない
	if !user.DeviceToken.Valid {
		return nil
	}

	if err := s.NoticeCommandRepository.StoreNotice(c, notice); err != nil {
		return errors.Wrap(err, "failed store notice")
	}

	// 未読のプッシュ通知数
	// アプリに表示するバッジの数
	unReadCount, err := s.NoticeQueryRepository.UnreadPushNoticeCount(c, sendTargetUserID)
	if err != nil {
		return errors.Wrap(err, "failed count unread push_notice")
	}

	cloudMsgData := firebaseEntity.NewCloudMessageData(user.DeviceToken.String, model.PushNoticeBody(triggeredUser.Name, notice.ActionTargetType, notice.ActionType), map[string]string{"endpoint": notice.Endpoint, "noticeId": strconv.Itoa(notice.ID)}, unReadCount)

	if err := s.CloudMessageCommandRepository.Send(cloudMsgData); err != nil {
		// MEMO: プッシュ通知のエラーは握り潰す
		logger.Error(err.Error())
	}

	return nil
}

func (s *NoticeDomainServiceImpl) reviewEndpoint(review *entity.Review) string {
	var endpoint string
	if review.TouristSpotID.Valid {
		endpoint = fmt.Sprintf("/tourism/location/%d/review/%d", int(review.TouristSpotID.Int64), review.ID)
	} else {
		endpoint = fmt.Sprintf("/hotels/h_%d/review/%d", int(review.InnID.Int64), review.ID)
	}

	return endpoint
}

func (s *NoticeDomainServiceImpl) reviewCommentEndpoint(review *entity.Review, comment *entity.ReviewComment) string {
	baseEndpoint := s.reviewEndpoint(review)
	return fmt.Sprintf("%s?type=Comment&id=%d", baseEndpoint, comment.ID)
}

func (s *NoticeDomainServiceImpl) reviewCommentReplyEndpoint(review *entity.Review, reply *entity.ReviewCommentReply) string {
	baseEndpoint := s.reviewEndpoint(review)
	return fmt.Sprintf("%s?type=Reply&id=%d", baseEndpoint, reply.ID)
}
