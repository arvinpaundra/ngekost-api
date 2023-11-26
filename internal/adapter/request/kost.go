package request

type (
	CreateKost struct {
		OwnerId         string  `json:"owner_id" form:"owner_id" validate:"required"`
		Name            string  `json:"name" form:"name" validate:"required,min=3,max=250"`
		Description     string  `json:"description" form:"description" validate:"required"`
		Type            string  `json:"type" form:"type" validate:"required,oneof=male female mixed"`
		PaymentInterval string  `json:"payment_interval" form:"payment_interval" validate:"required,oneof=monthly daily weekly annualy"`
		Province        string  `json:"province" form:"province" validate:"required"`
		City            string  `json:"city" form:"city" validate:"required"`
		District        string  `json:"district" form:"district" validate:"required"`
		Subdistrict     string  `json:"subdistrict" form:"subdistrict" validate:"required"`
		Latitude        float64 `json:"latitude" form:"latitude" validate:"required"`
		Longitude       float64 `json:"longitude" form:"longitude" validate:"required"`
	}

	UpdateKost struct {
		Name            string  `json:"name" form:"name" validate:"required,min=3,max=250"`
		Description     string  `json:"description" form:"description" validate:"required"`
		Type            string  `json:"type" form:"type" validate:"required,oneof=male female mixed"`
		PaymentInterval string  `json:"payment_interval" form:"payment_interval" validate:"required,oneof=monthly daily weekly annualy"`
		Province        string  `json:"province" form:"province" validate:"required"`
		City            string  `json:"city" form:"city" validate:"required"`
		District        string  `json:"district" form:"district" validate:"required"`
		Subdistrict     string  `json:"subdistrict" form:"subdistrict" validate:"required"`
		Latitude        float64 `json:"latitude" form:"latitude" validate:"required"`
		Longitude       float64 `json:"longitude" form:"longitude" validate:"required"`
	}
)
