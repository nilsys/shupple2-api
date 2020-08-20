package input

type (
	ListPayment struct {
		CfProjectID int `query:"cfProjectId"`
		PerPage     int `query:"perPage"`
		Page        int `query:"page"`
	}
)

const defaultListPaymentPerPage = 10

func (i ListPayment) GetLimit() int {
	if i.PerPage == 0 {
		return defaultListPaymentPerPage
	}
	return i.PerPage
}

func (i ListPayment) GetOffSet() int {
	if i.Page == 1 || i.Page == 0 {
		return 0
	}
	return i.GetLimit() * (i.Page - 1)
}
