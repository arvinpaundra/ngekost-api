package request

type (
	CreateRent struct {
		LesseeId string `json:"lessee_id" form:"lessee_id" validate:"required"`
		RoomId   string `json:"room_id" form:"room_id" validate:"required"`
	}

	UpdateRent struct {
	}
)
