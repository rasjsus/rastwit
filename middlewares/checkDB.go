package middlewares

import (
	"net/http"

	"github.com/rasjsus/rastwit/db"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Lost Connection", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	}
}
