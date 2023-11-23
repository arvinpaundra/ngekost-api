package entity

import "time"

type Owner struct {
	ID        string
	UserId    string
	Fullname  string
	Gender    string
	Phone     string
	Address   string
	City      string
	Birthdate *time.Time
	Status    *string
	Photo     *string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
}
