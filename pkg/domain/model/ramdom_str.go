package model

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/pkg/errors"
)

// 指定した桁数でletters内の候補からランダムな文字列生成
func RandomStr(n int) (string, error) {
	randomBytes := make([]byte, n)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", errors.Wrap(err, "failed reads cryptographically secure pseudorandom")
	}
	return base64.StdEncoding.EncodeToString(randomBytes)[:n], nil
}
