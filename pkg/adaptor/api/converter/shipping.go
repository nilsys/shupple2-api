package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

func (c Converters) ConvertShippingAddressInputToEntity(userID int, i *input.ShippingAddress) *entity.ShippingAddress {
	return entity.NewShippingAddress(userID, i.FirstName, i.LastName, i.FirstNameKana, i.LastNameKana, i.PhoneNumber, i.PostalNumber, i.Prefecture, i.City, i.Address, i.Building, i.Email)
}
