package service

import (
	"fmt"

	"github.com/lestrrat-go/jwx/jwk"

	"github.com/stayway-corp/stayway-media-api/pkg/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type (
	AuthService interface {
		Authorize(tokenString string) (string, error)
	}
)

func ProvideAuthService(config *config.Config) (AuthService, error) {
	if config.IsDev() {
		return NewAuthServiceForLocalImpl(config.Development), nil
	}

	return NewAuthServiceImpl(&config.AWS)
}

type (
	AuthServiceImpl struct {
		ClientID string
		Keys     *jwk.Set
	}
)

const (
	cognitoKeyURLFormat = "https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json"
)

func NewAuthServiceImpl(config *config.AWS) (*AuthServiceImpl, error) {
	keys, err := jwk.Fetch(fmt.Sprintf(cognitoKeyURLFormat, config.Region, config.UserPoolID))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get cognito keys")
	}

	return &AuthServiceImpl{ClientID: config.ClientID, Keys: keys}, nil
}

func (a AuthServiceImpl) Authorize(tokenString string) (string, error) {
	token, err := a.verifyToken(tokenString)
	if err != nil {
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)
	return claims.Subject, nil
}

func (a AuthServiceImpl) verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, a.getSignInKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	claims := token.Claims.(*jwt.StandardClaims)
	if !token.Valid || !claims.VerifyAudience(a.ClientID, true) {
		return nil, errors.New("token verification failed")
	}

	return token, nil
}

func (a AuthServiceImpl) getSignInKey(token *jwt.Token) (interface{}, error) {
	kid := fmt.Sprint(token.Header["kid"])
	candidateKeys := a.Keys.LookupKeyID(kid)
	if len(candidateKeys) == 0 {
		return nil, errors.Errorf("key not found; kid=%s", kid)
	}
	key := candidateKeys[0]

	if fmt.Sprint(token.Header["alg"]) != key.Algorithm() {
		return nil, errors.Errorf("unexpected jwt signing method=%v", token.Header["alg"])
	}
	return key.Materialize()
}

type (
	AuthServiceForLocalImpl struct {
		CognitoID string
	}
)

func NewAuthServiceForLocalImpl(config *config.Development) *AuthServiceForLocalImpl {
	return &AuthServiceForLocalImpl{CognitoID: config.CognitoID}
}

func (a *AuthServiceForLocalImpl) Authorize(tokenString string) (string, error) {
	return a.CognitoID, nil
}
