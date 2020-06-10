package model

import "fmt"

func UploadedS3Path(uuid string) string {
	return "tmp/" + uuid
}

func UserS3Path(uuid string) string {
	return fmt.Sprintf("user/%s", uuid)
}
