package entity

import "fmt"

type (
	ShippingAddress struct {
		ID            int    `gorm:"primary_key"`
		UserID        int    `json:"-"`
		FirstName     string `json:"firstName"`
		LastName      string `json:"lastName"`
		FirstNameKana string `json:"firstNameKana"`
		LastNameKana  string `json:"lastNameKana"`
		PhoneNumber   string `json:"phoneNumber"`
		PostalNumber  string `json:"postalNumber"`
		Prefecture    string `json:"prefecture"`
		City          string `json:"city"`
		Address       string `json:"address"`
		Building      string `json:"building"`
		Email         string `json:"email"`
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
