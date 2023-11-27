package request

type (
	QueryParamBill struct {
		Common
		Status    string
		StartDate string
		EndDate   string
	}
)
