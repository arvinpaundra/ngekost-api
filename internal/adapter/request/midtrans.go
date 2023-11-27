package request

type (
	CallbackPaymentMidtrans struct {
		TransactionTime   string `json:"transaction_time"`
		TransactionStatus string `json:"transaction_status"`
		PaymentType       string `json:"payment_type"`
		OrderID           string `json:"order_id"`
		GrossAmount       string `json:"gross_amount"`
		FraudStatus       string `json:"fraud_status"`
		Currency          string `json:"currency"`
		Bank              string `json:"bank"`
	}

	CreateBeneficiary struct {
		Name      string `json:"name" form:"name" validate:"required"`
		Account   string `json:"account" form:"account" validate:"required"`
		Bank      string `json:"bank" form:"bank" validate:"required"`
		AliasName string `json:"alias_name" form:"alias_name" validate:"required"`
		Email     string `json:"email" form:"email" validate:"required,email"`
	}

	UpdateBeneficiary struct {
		Name      string `json:"name" form:"name" validate:"required"`
		Account   string `json:"account" form:"account" validate:"required"`
		Bank      string `json:"bank" form:"bank" validate:"required"`
		AliasName string `json:"alias_name" form:"alias_name" validate:"required"`
		Email     string `json:"email" form:"email" validate:"required,email"`
	}

	CreatePayout struct {
		BeneficiaryName    string `json:"beneficiary_name" form:"beneficiary_name" validate:"required"`
		BeneficiaryAccount string `json:"beneficiary_account" form:"beneficiary_account" validate:"required"`
		BeneficiaryBank    string `json:"beneficiary_bank" form:"beneficiary_bank" validate:"required"`
		Amount             string `json:"amount" form:"amount" validate:"required"`
		Notes              string `json:"notes" form:"notes" validate:"required"`
	}

	ApprovePayout struct {
		ReferenceNos []string `json:"reference_nos" form:"reference_nos" validate:"required"`
	}

	RejectPayout struct {
		ReferenceNos []string `json:"reference_nos" form:"reference_nos" validate:"required"`
		RejectReason string   `json:"reject_reason" form:"reject_reason" validate:"required"`
	}
)
