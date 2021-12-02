package routers

import (
	"encoding/json"
	"net/http"

	"github.com/rasjsus/rastwit/db"
)

func Perfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	perfil, err := db.LookPerfil(ID)
	if err != nil {
		http.Error(w, "Error while looking for a register"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
