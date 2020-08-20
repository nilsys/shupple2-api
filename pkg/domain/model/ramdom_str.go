package model

import (
	"crypto/rand"
	"encoding/base32"

	"github.com/pkg/errors"
)

func RandomStr(n int) (string, error) {
	randomBytes := make([]byte, n)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", errors.Wrap(err, "failed reads cryptographically secure pseudorandom")
	}

	return base32.StdEncoding.EncodeToString(randomBytes)[:n], nil
}
