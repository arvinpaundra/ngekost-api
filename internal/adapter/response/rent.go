package response

import (
	"time"

	"github.com/arvinpaundra/ngekost-api/internal/entity"
)

type (
	LesseeRent struct {
		ID              string     `json:"id"`
		LesseeId        string     `json:"lessee_id"`
		RoomId          string     `json:"room_id"`
		KostId          string     `json:"kost_id"`
		KostName        string     `json:"kost_name"`
		KostType        string     `json:"kost_type"`
		RoomName        string     `json:"room_name"`
		RoomCategory    *string    `json:"room_category"`
		RoomPrice       float64    `json:"room_price"`
		PaymentInterval string     `json:"payment_interval"`
		City            string     `json:"city"`
		StartDate       time.Time  `json:"start_date"`
		EndDate         *time.Time `json:"end_date"`
		KostImage       *string    `json:"kost_image"`
	}

	KostRenter struct {
		ID        string     `json:"id"`
		LesseeId  string     `json:"lessee_id"`
		RoomId    string     `json:"room_id"`
		KostId    string     `json:"kost_id"`
		Fullname  string     `json:"fullname"`
		Gender    string     `json:"gender"`
		Phone     string     `json:"phone"`
		City      string     `json:"city"`
		StartDate time.Time  `json:"start_date"`
		EndDate   *time.Time `json:"end_date"`
		Photo     *string    `json:"photo"`
	}
)

func ToResponseLesseeRents(rents []*entity.Rent) []*LesseeRent {
	var res []*LesseeRent

	for _, rent := range rents {
		res = append(res, &LesseeRent{
			ID:              rent.ID,
			LesseeId:        rent.LesseeId,
			RoomId:          rent.RoomId,
			KostId:          rent.Room.KostId,
			KostName:        rent.Room.Kost.Name,
			KostType:        rent.Room.Kost.Type,
			RoomName:        rent.Room.Name,
			RoomCategory:    rent.Room.Category,
			RoomPrice:       rent.Room.Price,
			PaymentInterval: rent.Room.Kost.PaymentInterval,
			City:            rent.Lessee.City,
			StartDate:       rent.StartDate,
			EndDate:         rent.EndDate,
			KostImage:       rent.Room.Kost.Image,
		})
	}

	return res
}

func ToResponseKostRenters(rents []*entity.Rent) []*KostRenter {
	var res []*KostRenter

	for _, rent := range rents {
		res = append(res, &KostRenter{
			ID:        rent.ID,
			LesseeId:  rent.LesseeId,
			RoomId:    rent.RoomId,
			KostId:    rent.Room.KostId,
			Fullname:  rent.Lessee.Fullname,
			Gender:    rent.Lessee.Gender,
			Phone:     rent.Lessee.Phone,
			City:      rent.Lessee.City,
			StartDate: rent.StartDate,
			EndDate:   rent.EndDate,
			Photo:     rent.Lessee.Photo,
		})
	}

	return res
}
