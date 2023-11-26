package request

type (
	KostRulePathParam struct {
		KostId     string
		KostRuleId string
	}

	CreateKostRule struct {
		Title       string  `json:"title" form:"title" validate:"required,min=3,max=250"`
		Priority    string  `json:"priority" validate:"required,oneof=low medium high"`
		Description *string `json:"description"`
	}

	UpdateKostRule struct {
		Title       string  `json:"title" form:"title" validate:"required,min=3,max=250"`
		Priority    string  `json:"priority" validate:"required,oneof=low medium high"`
		Description *string `json:"description"`
	}
)
