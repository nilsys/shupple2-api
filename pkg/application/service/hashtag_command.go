package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	HashtagCommandService interface {
		FindOrCreateHashtag(hashtag *entity.Hashtag) (*entity.Hashtag, error)
		FindOrCreateHashtags(hashtags []string) ([]*entity.Hashtag, error)
	}

	HashtagCommandServiceImpl struct {
		repository.HashtagQueryRepository
		repository.HashtagCommandRepository
	}
)

var HashtagCommandServiceSet = wire.NewSet(
	wire.Struct(new(HashtagCommandServiceImpl), "*"),
	wire.Bind(new(HashtagCommandService), new(*HashtagCommandServiceImpl)),
)

func (s *HashtagCommandServiceImpl) FindOrCreateHashtag(hashtag *entity.Hashtag) (*entity.Hashtag, error) {
	return s.HashtagCommandRepository.FirstOrCreate(hashtag)
}

func (s *HashtagCommandServiceImpl) FindOrCreateHashtags(hashtags []string) ([]*entity.Hashtag, error) {
	existingHashtags, err := s.HashtagQueryRepository.FindByNames(hashtags)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find existingHashtags")
	}

	result := make([]*entity.Hashtag, len(hashtags))
	for i, hashtag := range hashtags {
		if row, ok := existingHashtags[hashtag]; ok {
			result[i] = row
			continue
		}

		newHashtag := &entity.Hashtag{Name: hashtag}
		if err := s.HashtagCommandRepository.Store(newHashtag); err != nil {
			return nil, errors.Wrap(err, "failed to create new hashtag")
		}
		result[i] = newHashtag
	}

	return result, nil
}
