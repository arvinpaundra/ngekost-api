package entity

import "time"

type Payment struct {
	ID        string
	BillId    string
	Url       string
	Nominal   string
	Method    string
	Status    string
	Via       *string
	Raw       *string
	CreatedAt time.Time
	UpdatedAt time.Time
}
