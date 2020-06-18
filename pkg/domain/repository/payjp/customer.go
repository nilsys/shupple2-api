package payjp

type (
	CustomerCommandRepository interface {
		StoreCustomer(customerID string, email string) error
	}
)
