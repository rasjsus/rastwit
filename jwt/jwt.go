package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/rasjsus/rastwit/models"
)

func GenerateJWT(user models.Usuario) (string, error) {
	myKey := []byte("some_secret_key")

	payload := jwt.MapClaims{
		"email":      user.Email,
		"name":       user.Name,
		"last_name":  user.LastName,
		"birth_date": user.BirthDate,
		"biography":  user.Biography,
		"ubication":  user.Ubication,
		"_id":        user.ID.Hex(),
		"website":    user.WebSite,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return tokenString, err
	}

	return tokenString, nil

}
