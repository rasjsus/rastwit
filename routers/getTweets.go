package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rasjsus/rastwit/db"
)

func GetTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Missing id variable", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Missing page variable", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page <= 0 {
		http.Error(w, "page not allowed", http.StatusBadRequest)
		return
	}

	pageInt := int64(page)
	response, correct := db.GetTweets(ID, pageInt)
	if correct == false {
		http.Error(w, "Errro while reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
