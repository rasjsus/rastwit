package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rasjsus/rastwit/db"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")
	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "page is required and shouel be > 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	result, status := db.GetUsers(IDUser, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error while reading users: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
