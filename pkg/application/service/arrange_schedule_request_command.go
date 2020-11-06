package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/uma-co82/shupple2-api/pkg/domain/entity"
	"github.com/uma-co82/shupple2-api/pkg/domain/model/command"
	"github.com/uma-co82/shupple2-api/pkg/domain/model/serror"
	"github.com/uma-co82/shupple2-api/pkg/domain/repository"
)

type (
	ArrangeScheduleRequestCommandService interface {
		Store(cmd *command.StoreArrangeScheduleRequest, user *entity.UserTiny) error
	}

	ArrangeScheduleRequestCommandServiceImpl struct {
		repository.UserQueryRepository
		repository.ArrangeScheduleRequestCommandRepository
	}
)

func (s *ArrangeScheduleRequestCommandServiceImpl) Store(cmd *command.StoreArrangeScheduleRequest, user *entity.UserTiny) error {
	isMainMatching, err := s.UserQueryRepository.IsExistMainMatchingUserMatchingHistory(user.ID, cmd.MatchingUserID)
	if err != nil {
		return errors.Wrap(err, "failed find user_matching_history")
	}

	// 本マッチングしていないユーザーに送ろうとした場合
	if !isMainMatching {
		return serror.New(nil, serror.CodeNotMatching, "not matching")
	}

	req := entity.NewArrangeScheduleRequest(user.ID, cmd.MatchingUserID, cmd.Date, cmd.Remark)

	return s.ArrangeScheduleRequestCommandRepository.Store(context.Background(), req)
}
