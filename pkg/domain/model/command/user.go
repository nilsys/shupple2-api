package command

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type UpdateUser struct {
	Name         string
	Email        string
	BirthDate    model.Date
	Gender       model.Gender
	Profile      string
	IconUUID     string
	HeaderUUID   string
	URL          string
	FacebookURL  string
	InstagramURL string
	TwitterURL   string
	LivingArea   string
	Interests    []*entity.UserInterest
}
