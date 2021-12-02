package routers

import (
	"encoding/json"
	"net/http"

	"github.com/rasjsus/rastwit/db"
	"github.com/rasjsus/rastwit/jwt"
	"github.com/rasjsus/rastwit/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "User and / or password invalid "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email ir required ", http.StatusBadRequest)
		return
	}

	document, exists := db.DbLogin(t.Email, t.Password)
	if !exists {
		http.Error(w, "User and / or password invalid ", http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "error getting token "+err.Error(), http.StatusBadRequest)
	}

	resp := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}
