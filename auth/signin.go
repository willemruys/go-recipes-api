package auth

import (
	"example.com/m/models"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(email, password string) (string, error) {

	var err error

	user := models.User{}

	err = user.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return CreateToken(user.ID)
}