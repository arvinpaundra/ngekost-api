package request

import "mime/multipart"

type (
	RoomPathParam struct {
		KostId string
		RoomId string
	}

	CreateRoom struct {
		Name        string                `json:"name" form:"name" validate:"required,min=3,max=250"`
		Quantity    int                   `json:"quantity" form:"quantity" validate:"required,number,min=1"`
		Price       float64               `json:"price" form:"price" validate:"required,number"`
		Description *string               `json:"description" form:"description"`
		Category    *string               `json:"category" form:"category"`
		Image       *multipart.FileHeader `json:"image" form:"image"`
	}

	UpdateRoom struct {
		Name        string                `json:"name" form:"name" validate:"required,min=3,max=250"`
		Quantity    int                   `json:"quantity" form:"quantity" validate:"required,number,min=1"`
		Price       float64               `json:"price" form:"price" validate:"required,number"`
		Description *string               `json:"description" form:"description"`
		Category    *string               `json:"category" form:"category"`
		Image       *multipart.FileHeader `json:"image" form:"image"`
	}
)
