package routers

import (
	"encoding/json"
	"net/http"

	"github.com/rasjsus/rastwit/db"
	"github.com/rasjsus/rastwit/models"
)

func GetRealation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	var response models.RelationResponse

	status, err := db.GetRelation(t)
	if err != nil || status == false {
		response.Status = false
	} else {
		response.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
