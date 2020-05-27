package entity

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type Interest struct {
	ID            int                 `gorm:"primary_key" json:"id"`
	Name          string              `json:"name"`
	InterestGroup model.InterestGroup `json:"interestGroup"`
}
