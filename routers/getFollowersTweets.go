package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rasjsus/rastwit/db"
)

func GetFollowersTweets(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "page is required", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "page not allowed", http.StatusBadRequest)
		return
	}

	response, correct := db.GetFollowersTweets(IDUser, page)
	if correct == false {
		http.Error(w, "error while reading followers tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
