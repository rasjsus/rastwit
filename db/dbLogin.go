package db

import (
	"github.com/rasjsus/rastwit/models"
	"golang.org/x/crypto/bcrypt"
)

func DbLogin(email, password string) (models.Usuario, bool) {
	user, found, _ := CheckUsersExists(email)
	if !found {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}

	return user, true
}
