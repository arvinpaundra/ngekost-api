package constant

import "errors"

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrOwnerNotFound   = errors.New("owner not found")
	ErrSessionNotFound = errors.New("session not found")

	ErrUsernameAlreadyUsed = errors.New("username already used")
	ErrPasswordIncorrect   = errors.New("password is incorrect")
	ErrInvalidRole         = errors.New("invalid role")
)
