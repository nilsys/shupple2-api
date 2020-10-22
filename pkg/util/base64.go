package util

import (
	"bytes"
	"encoding/base64"

	"github.com/pkg/errors"
)

func Base64StrWriteBuffer(base string) (*bytes.Buffer, error) {
	data, err := base64.StdEncoding.DecodeString(base)
	if err != nil {
		return nil, errors.Wrap(err, "failed decode str")
	}
	wb := new(bytes.Buffer)
	wb.Write(data)
	return wb, nil
}
