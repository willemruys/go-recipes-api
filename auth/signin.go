package auth

import (
	"example.com/m/models"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(user models.User, attempt string) (string, error) {

	var err error
	err = user.VerifyPassword(user.Password, attempt)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return CreateToken(user.ID)
}