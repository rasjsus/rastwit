package middlewares

import (
	"net/http"

	"github.com/rasjsus/rastwit/routers"
)

func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Invalid token "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
