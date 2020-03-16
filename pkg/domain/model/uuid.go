package model

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewRandUUID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", errors.Wrap(err, "failed to generate uuid")
	}
	return u.String(), nil
}
