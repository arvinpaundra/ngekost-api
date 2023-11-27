package response

type (
	MidtransBeneficiary struct {
		Name      string `json:"name"`
		Bank      string `json:"bank"`
		Account   string `json:"account"`
		AliasName string `json:"alias_name"`
		Email     string `json:"email"`
	}

	MidtransPayout struct {
		Status      string `json:"status"`
		ReferenceNo string `json:"reference_no"`
	}
)
