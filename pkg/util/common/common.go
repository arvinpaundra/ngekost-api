package common

import (
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GetID() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}

func HashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashed)
}

func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
