package constant

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrOwnerNotFound     = errors.New("owner not found")
	ErrSessionNotFound   = errors.New("session not found")
	ErrKostNotFound      = errors.New("kost not found")
	ErrRoomNotFound      = errors.New("room not found")
	ErrRoomAssetNotFound = errors.New("room asset not found")
	ErrKostRuleNotFound  = errors.New("kost rule not found")
	ErrBillNotFound      = errors.New("bill not found")
	ErrPaymentNotFound   = errors.New("payment not found")
	ErrRentNotFound      = errors.New("rent not found")
	ErrLesseeNotFound    = errors.New("lessee not found")

	ErrUsernameAlreadyUsed = errors.New("username already used")
	ErrPasswordIncorrect   = errors.New("password is incorrect")
	ErrInvalidRole         = errors.New("invalid role")
	ErrActiveRent          = errors.New("you have an active rent")
)
