package model

import (
	"fmt"
	"io"
)

const QueueAttributeNameMediaType = "media_type"

var maxMediaSize = map[MediaType]Size{
	MediaTypeUserIcon:    {500, 500},
	MediaTypeUserHeader:  {1500, 1500},
	MediaTypeReviewImage: {1500, 1500},
	MediaTypeReviewVideo: {1500, 1500},
}

type (
	MediaBody struct {
		ContentType string
		Body        io.ReadCloser
	}

	PersistMediaRequest struct {
		UUID        string    `json:"uuid"`
		Destination string    `json:"destination"`
		MediaType   MediaType `json:"media_type"`
	}

	Size struct {
		Width  uint
		Height uint
	}
)

func UploadedS3Path(uuid string) string {
	return "tmp/" + uuid
}

func UserS3Path(uuid string) string {
	return fmt.Sprintf("user/%s", uuid)
}

func (s Size) JoinByX() string {
	return fmt.Sprint(s.Width, "x", s.Height)
}

func MaxMediaSize(mediaType MediaType) Size {
	return maxMediaSize[mediaType]
}
