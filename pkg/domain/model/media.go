package model

func UploadedS3Path(uuid string) string {
	return "tmp/" + uuid
}
