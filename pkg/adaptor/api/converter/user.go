package converter

import (
	"time"

	"github.com/uma-co82/shupple2-api/pkg/domain/model"

	"github.com/uma-co82/shupple2-api/pkg/adaptor/api/output"
	"github.com/uma-co82/shupple2-api/pkg/domain/entity"

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

func (c Converters) ConvertUser2Output(user *entity.User) output.User {
	images := make([]output.UserImage, len(user.Images))
	for i, image := range user.Images {
		images[i] = output.UserImage{
			Priority: image.Priority,
			URL:      image.URL(c.filesURL()),
		}
	}

	return output.User{
		ID:             user.ID,
		Name:           user.Name,
		Birthdate:      model.Date(user.Birthdate),
		Profile:        user.Profile,
		Gender:         user.Gender,
		Prefecture:     user.Prefecture,
		MatchingReason: user.MatchingReason,
		IsMatching:     user.IsMatching,
		Images:         images,
	}
}
