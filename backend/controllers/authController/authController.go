package authController

import (
	"management-backend/shared"

	"golang.org/x/crypto/bcrypt"
)

type AuthController struct{}

func GetPasswordHash(p string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), shared.PW_COST)
	return string(bytes), err
}

func ComparePasswordHash(p string, h string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
	return err == nil
}
