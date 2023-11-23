package request

type (
	Register struct {
		Username string `json:"username" form:"username" validate:"required,min=3,max=50"`
		Password string `json:"password" form:"password" validate:"required,min=8,max=255"`
		Role     string `json:"role" form:"role" validate:"required,oneof=owner lessee"`
		Fullname string `json:"fullname" form:"fullname" validate:"required,min=3,max=100"`
		Phone    string `json:"phone" form:"phone" validate:"required,numeric,min=11,max=15"`
		Gender   string `json:"gender" form:"gender" validate:"required,oneof=male female"`
		City     string `json:"city" form:"city" validate:"required,min=3,max=50"`
		Address  string `json:"address" form:"address" validate:"required,max=250"`
	}

	Login struct {
		Username   string  `json:"username" form:"username" validate:"required"`
		Password   string  `json:"password" form:"password" validate:"required"`
		DeviceId   string  `json:"device_id" form:"device_id" validate:"required"`
		Platform   string  `json:"platform" form:"platform" validate:"required,oneof=web android ios"`
		DeviceName string  `json:"device_name" form:"device_name" validate:"required,min=3,max=250"`
		FCMToken   *string `json:"fcm_token" form:"fcm_token"`
		IPAddress  string
	}

	Logout struct {
		UserId   string `json:"user_id" form:"user_id" validate:"required"`
		DeviceId string `json:"device_id" form:"device_id" validate:"required"`
	}
)
