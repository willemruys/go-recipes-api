package auth

import (
	"golang.org/x/crypto/bcrypt"
	"recipes-api.com/m/models"
)

func AssignToken(user models.User, attempt string) (string, error) {

	var err error
	err = user.VerifyPassword(user.Password, attempt)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return CreateToken(uint32(user.ID))
}