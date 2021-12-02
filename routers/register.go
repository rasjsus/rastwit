package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rasjsus/rastwit/db"
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

	_, found, _ := db.CheckUsersExists(t.Email)
	if found == true {
		http.Error(w, "There is an user with this email in BD", http.StatusBadRequest)
		return
	}

	_, status, err := db.InsertRegister(&t)
	if err != nil {
		http.Error(w, "An error happened on BD while trying to register an user"+err.Error(), http.StatusInternalServerError)
		return
	}
	if status == false {
		http.Error(w, "We couldn't insert the user in the DB", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
