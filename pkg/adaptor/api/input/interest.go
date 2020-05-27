package input

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type ListInterest struct {
	InterestGroup model.InterestGroup `query:"interestGroup"`
}
