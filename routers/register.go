package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rasjsus/rastwit/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	h := r.Header.Get("Authorization")
	fmt.Println(h)
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error on user data: "+err.Error(), http.StatusBadRequest)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Error, email can't be empty", http.StatusBadRequest)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Error, password has to be longer than 6 characters", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
