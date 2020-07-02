package entity

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	MailTemplate interface {
		TemplateName() model.MailTemplateName
		DefaultData() (string, error)
		ToJSON() (string, error)
	}

	ThanksPurchaseTemplate struct {
		OwnerName             string `json:"ownername"`
		ReturnGiftDescription string `json:"returngiftdescription"`
		ChargeID              string `json:"chargeid"`
		Price                 string `json:"price"`
		UserEmail             string `json:"useremail"`
		UserShippingAddress   string `json:"usershippingaddress"`
		UserName              string `json:"username"`
	}
)

func NewThanksPurchaseTemplate(ownerName, returnGiftDesc, chargeID, price, userEmail, userShippingAddress, userName string) *ThanksPurchaseTemplate {
	return &ThanksPurchaseTemplate{
		OwnerName:             ownerName,
		ReturnGiftDescription: returnGiftDesc,
		ChargeID:              chargeID,
		Price:                 price,
		UserEmail:             userEmail,
		UserShippingAddress:   userShippingAddress,
		UserName:              userName,
	}
}

func (t *ThanksPurchaseTemplate) TemplateName() model.MailTemplateName {
	return model.MailTemplateNameThanksPurchase
}

func (t *ThanksPurchaseTemplate) DefaultData() (string, error) {
	s := ThanksPurchaseTemplate{}
	bytes, err := json.Marshal(s)
	if err != nil {
		return "", errors.Wrap(err, "failed marshal")
	}
	return string(bytes), nil
}

func (t *ThanksPurchaseTemplate) ToJSON() (string, error) {
	bytes, err := json.Marshal(t)
	if err != nil {
		return "", errors.Wrap(err, "failed marshal")
	}
	return string(bytes), nil
}
