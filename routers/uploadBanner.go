package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/rasjsus/rastwit/db"
	"github.com/rasjsus/rastwit/models"
)

func UploadBanner(w http.ResponseWriter, r *http.Request) {
	var perfil models.Usuario

	file, handler, err := r.FormFile("banner")
	if err != nil {
		http.Error(w, "banner file not found!"+err.Error(), http.StatusBadRequest)
		return
	}

	perfil, err = db.LookPerfil(IDUser)
	if err != nil {
		http.Error(w, "user not found!"+err.Error(), http.StatusBadRequest)
		return
	}

	if perfil.Banner != "" {
		err = os.Remove("uploads/banners/" + perfil.Banner)
		if err != nil {
		}
	}

	stringSplit := strings.Split(handler.Filename, ".")
	var extension = strings.Split(handler.Filename, ".")[len(stringSplit)-1]
	var archivo string = "uploads/banners/" + IDUser + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error while uploading image !"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error while copying image !"+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Banner = IDUser + "." + extension
	status, err = db.UpdateRegister(usuario, IDUser)
	if err != nil || status == false {
		http.Error(w, "Error while inserting on db !"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
