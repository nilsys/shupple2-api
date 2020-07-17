package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

func (c Converters) ConvertShippingAddressInputToEntity(userID int, i *input.ShippingAddress) *entity.ShippingAddress {
	return entity.NewShippingAddress(userID, i.FirstName, i.LastName, i.FirstNameKana, i.LastNameKana, i.PhoneNumber, i.PostalNumber, i.Prefecture, i.City, i.Address, i.Building, i.Email)
}

func (c Converters) ConvertShippingAddressToOutput(address *entity.ShippingAddress) *output.ShippingAddress {
	return &output.ShippingAddress{
		ID:            address.ID,
		FirstName:     address.FirstName,
		LastName:      address.LastName,
		FirstNameKana: address.FirstNameKana,
		LastNameKana:  address.LastNameKana,
		PhoneNumber:   address.PhoneNumber,
		PostalNumber:  address.PostalNumber,
		Prefecture:    address.Prefecture,
		City:          address.City,
		Address:       address.Address,
		Building:      address.Building,
		Email:         address.Email,
	}
}
