package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rasjsus/rastwit/db"
	"github.com/rasjsus/rastwit/models"
)

func TweetRouter(w http.ResponseWriter, r *http.Request) {
	var message models.TweetResponse

	err := json.NewDecoder(r.Body).Decode(&message)

	register := models.Tweet{
		UserId:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := db.InsertTweet(register)
	if err != nil {
		http.Error(w, "Error while inserting tweet: "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "Could not insert tweet ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
