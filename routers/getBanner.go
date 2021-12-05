package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/rasjsus/rastwit/db"
)

func GetBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	perfil, err := db.LookPerfil(ID)
	if err != nil {
		http.Error(w, "user not found", http.StatusBadRequest)
		return
	}

	openFile, err := os.Open("uploads/banners/" + perfil.Banner)
	if err != nil {
		http.Error(w, "image not found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, openFile)
	if err != nil {
		http.Error(w, "error while copying image", http.StatusBadRequest)
	}
}
