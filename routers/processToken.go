package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/rasjsus/rastwit/db"
	"github.com/rasjsus/rastwit/models"
)

var Email string
var IDUser string

func ProcessToken(token string) (*models.Claim, bool, string, error) {
	myKey := []byte("some_secret_key")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid format token")
	}

	tk := strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, found, _ := db.CheckUsersExists(claims.Email)
		if found {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}

		return claims, found, IDUser, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}

	return claims, false, string(""), err
}
