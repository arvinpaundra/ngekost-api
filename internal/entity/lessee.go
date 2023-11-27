package entity

import "time"

type Lessee struct {
	ID        string
	UserId    string
	Fullname  string
	Gender    string
	Phone     string
	City      string
	Address   string
	Birthdate *time.Time
	Status    *string
	Photo     *string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
	Bills     []Bill
}
