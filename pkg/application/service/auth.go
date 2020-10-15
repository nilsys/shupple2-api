package service

import (
	"context"

	firebase "firebase.google.com/go"
	"github.com/pkg/errors"
	"github.com/uma-co82/shupple2-api/pkg/config"
)

type (
	AuthService interface {
		Authorize(token string) (string, error)
	}

	AuthServiceForLocalImpl struct {
		FirebaseID string
	}

	AuthServiceImpl struct {
		FirebaseApp *firebase.App
	}
)

func ProvideAuthService(config *config.Config) (AuthService, error) {
	return NewAuthServiceForLocalImpl(config.Development), nil
}

func NewAuthServiceForLocalImpl(config *config.Development) *AuthServiceForLocalImpl {
	return &AuthServiceForLocalImpl{FirebaseID: config.FirebaseID}
}

func (s *AuthServiceForLocalImpl) Authorize(token string) (string, error) {
	return s.FirebaseID, nil
}

func (s *AuthServiceImpl) Authorize(token string) (string, error) {
	auth, err := s.FirebaseApp.Auth(context.Background())
	if err != nil {
		return "", errors.Wrap(err, "")
	}

	id, err := auth.VerifyIDToken(context.Background(), token)
	if err != nil {
		return "", errors.Wrap(err, "")
	}

	return id.UID, nil
}
