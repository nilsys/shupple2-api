package entity

import "fmt"

type (
	ShippingAddress struct {
		ID            int `gorm:"primary_key"`
		UserID        int
		FirstName     string
		LastName      string
		FirstNameKana string
		LastNameKana  string
		PhoneNumber   string
		PostalNumber  string
		Prefecture    string
		City          string
		Address       string
		Building      string
		Email         string
		Times
	}
)

func NewShippingAddress(userID int, firstName, lastName, firstNameKana, lastNameKana, phoneNumber, postalNumber, prefecture, city, address, building, email string) *ShippingAddress {
	return &ShippingAddress{
		UserID:        userID,
		FirstName:     firstName,
		LastName:      lastName,
		FirstNameKana: firstNameKana,
		LastNameKana:  lastNameKana,
		PhoneNumber:   phoneNumber,
		PostalNumber:  postalNumber,
		Prefecture:    prefecture,
		City:          city,
		Address:       address,
		Building:      building,
		Email:         email,
	}
}

func (s *ShippingAddress) FullAddress() string {
	return fmt.Sprintf("〒%s %s%s%s%s", s.PostalNumber, s.Prefecture, s.City, s.Address, s.Building)
}
