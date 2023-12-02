package entity

import "time"

type Rent struct {
	ID        string
	RoomId    string
	LesseeId  string
	StartDate time.Time
	EndDate   *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	Room      Room
	Lessee    Lessee
}
