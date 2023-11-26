package response

import "github.com/arvinpaundra/ngekost-api/internal/entity"

type (
	Kost struct {
		ID              string  `json:"id"`
		OwnerId         string  `json:"owner_id"`
		Name            string  `json:"name"`
		Owner           string  `json:"owner"`
		Phone           string  `json:"phone"`
		Description     string  `json:"description"`
		Type            string  `json:"type"`
		PaymentInterval string  `json:"payment_interval"`
		Province        string  `json:"province"`
		City            string  `json:"city"`
		District        string  `json:"district"`
		Subdistrict     string  `json:"subdistrict"`
		Latitude        float64 `json:"latitude"`
		Longitude       float64 `json:"longitude"`
		Image           *string `json:"image"`
	}
)

func ToResponseKosts(kosts []*entity.Kost) []*Kost {
	var res []*Kost

	for _, kost := range kosts {
		res = append(res, &Kost{
			ID:              kost.ID,
			OwnerId:         kost.OwnerId,
			Name:            kost.Name,
			Owner:           kost.Owner.Fullname,
			Phone:           kost.Owner.Phone,
			Description:     kost.Description,
			Type:            kost.Type,
			PaymentInterval: kost.PaymentInterval,
			Province:        kost.Province,
			City:            kost.City,
			District:        kost.District,
			Subdistrict:     kost.Subdistrict,
			Latitude:        kost.Latitude,
			Longitude:       kost.Longitude,
			Image:           kost.Image,
		})
	}

	return res
}

func ToResponseKost(kost *entity.Kost) *Kost {
	return &Kost{
		ID:              kost.ID,
		OwnerId:         kost.OwnerId,
		Name:            kost.Name,
		Owner:           kost.Owner.Fullname,
		Phone:           kost.Owner.Phone,
		Description:     kost.Description,
		Type:            kost.Type,
		PaymentInterval: kost.PaymentInterval,
		Province:        kost.Province,
		City:            kost.City,
		District:        kost.District,
		Subdistrict:     kost.Subdistrict,
		Latitude:        kost.Latitude,
		Longitude:       kost.Longitude,
		Image:           kost.Image,
	}
}
