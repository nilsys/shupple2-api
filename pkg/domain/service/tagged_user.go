package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	TaggedUserDomainService interface {
		FindTaggedUsers(body string) ([]*entity.User, error)
	}

	TaggedUserDomainServiceImpl struct {
		repository.UserQueryRepository
	}
)

var TaggedUserDomainServiceSet = wire.NewSet(
	wire.Struct(new(TaggedUserDomainServiceImpl), "*"),
	wire.Bind(new(TaggedUserDomainService), new(TaggedUserDomainServiceImpl)),
)

func (s TaggedUserDomainServiceImpl) FindTaggedUsers(body string) ([]*entity.User, error) {
	taggedUIDs := model.FindTaggedUser(body)

	// メンションが含まれていないければからのSliceを返す
	if len(taggedUIDs) == 0 {
		return make([]*entity.User, 0), nil
	}

	// 重複を削除
	taggedDistinctUIDs := removeDuplicateUser(taggedUIDs)

	taggedUsers, err := s.UserQueryRepository.FindByUIDs(taggedDistinctUIDs)
	if err != nil {
		return nil, errors.Wrap(err, "fail to find users by tag")
	}
	return taggedUsers, nil
}

// sliceの重複を取り除く
// https://qiita.com/Sekky0905/items/ba2215981693b36e9982
func removeDuplicateUser(oldSlice []string) []string {
	newSlice := make([]string, 0)

	m := make(map[string]struct{})
	for _, element := range oldSlice {
		// mapでは、第二引数にその値が入っているかどうかの真偽値が入っている
		if _, ok := m[element]; !ok {
			m[element] = struct{}{}
			newSlice = append(newSlice, element)
		}
	}

	return newSlice
}
