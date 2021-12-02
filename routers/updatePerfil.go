package routers

import (
	"encoding/json"
	"net/http"

	"github.com/rasjsus/rastwit/db"
	"github.com/rasjsus/rastwit/models"
)

func UpdatePerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Incorrect Data: "+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool
	status, err = db.UpdateRegister(t, IDUser)
	if err != nil {
		http.Error(w, "An error while trying updating the register, try again: "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "user register not updated", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
