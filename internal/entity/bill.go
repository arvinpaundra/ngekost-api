package entity

import "time"

type Bill struct {
	ID          string
	LesseeId    string
	Invoice     string
	KostName    string
	RoomName    string
	RentType    string
	Status      string
	Price       float64
	Total       float64
	ServiceFee  *float64
	Discount    *float64
	DueDate     time.Time
	PaymentDate time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Lessee      Lessee
	Payment     Payment
}
