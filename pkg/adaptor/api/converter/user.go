package converter

import (
	"time"

	"github.com/uma-co82/shupple2-api/pkg/adaptor/api/input"
	"github.com/uma-co82/shupple2-api/pkg/domain/model/command"
)

func (c Converters) ConvertRegisterUserInput2Cmd(in *input.RegisterUser) command.StoreUser {
	images := make([]command.StoreUserImage, len(in.Images))
	for i, image := range images {
		images[i] = command.StoreUserImage{
			Priority:    image.Priority,
			MimeType:    image.MimeType,
			ImageBase64: image.ImageBase64,
		}
	}

	return command.StoreUser{
		Name:           in.Name,
		Email:          in.Email,
		Birthdate:      time.Time(in.Birthdate),
		Profile:        in.Profile,
		Gender:         in.Gender,
		Prefecture:     in.Prefecture,
		MatchingReason: in.MatchingReason,
		Images:         images,
	}
}
