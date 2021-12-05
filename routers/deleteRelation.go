package routers

import (
	"net/http"

	"github.com/rasjsus/rastwit/db"
	"github.com/rasjsus/rastwit/models"
)

func DeleteRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "user id required", http.StatusBadRequest)
		return
	}

	var t models.Relation

	t.UserID = IDUser
	t.UserRelationID = ID

	status, err := db.DeleteRelation(&t)
	if err != nil || status == false {
		http.Error(w, "error while deleting relation: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
