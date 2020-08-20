package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	UserValidatorDomainService interface {
		Do(user *entity.User) error
	}

	UserValidatorDomainServiceImpl struct {
		repository.UserQueryRepository
	}
)

var UserValidatorDomainServiceSet = wire.NewSet(
	wire.Struct(new(UserValidatorDomainServiceImpl), "*"),
	wire.Bind(new(UserValidatorDomainService), new(*UserValidatorDomainServiceImpl)),
)

// 新規登録時のValidate
func (s *UserValidatorDomainServiceImpl) Do(user *entity.User) error {
	// uid重複チェック
	isExist, err := s.UserQueryRepository.IsExistByUID(user.UID)
	if err != nil {
		return errors.Wrap(err, "failed to get user")
	}
	if isExist {
		return serror.New(nil, serror.CodeInvalidParam, "uid: %s is duplicate", user.UID)
	}

	// 属性付与
	user.AddAttribute(model.UserAttributeCommon)

	return nil
}
